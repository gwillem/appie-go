package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"text/tabwriter"

	appie "github.com/gwillem/appie-go"
)

type orderCommand struct {
	Closed bool `long:"closed" description:"List closed/delivered orders instead of open orders"`
	All    bool `long:"all" description:"List all orders, including open and closed"`

	Show   orderShowCommand   `command:"show" description:"Show contents of an order"`
	Add    orderAddCommand    `command:"add" description:"Add a product to an order"`
	Rm     orderRmCommand     `command:"rm" description:"Remove a product from an order"`
	Submit orderSubmitCommand `command:"submit" description:"Finalize reopened order changes"`
}

func (cmd *orderCommand) Execute(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("unknown argument %q, did you mean: appie order show %s", args[0], args[0])
	}
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	status, emptyLabel, err := cmd.listStatus()
	if err != nil {
		return err
	}

	fulfillments, err := client.GetFulfillmentsByStatus(ctx, status)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}

	if len(fulfillments) == 0 {
		fmt.Printf("No %s orders\n", emptyLabel)
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	fmt.Fprintf(w, "\t%s\t%s\t%s\t%s\t\n", "Order", "Status", "Delivery", "Total")
	for _, f := range fulfillments {
		delivery := f.Delivery.Slot.DateDisplay
		if f.Delivery.Slot.TimeDisplay != "" {
			delivery += "  " + f.Delivery.Slot.TimeDisplay
		}
		fmt.Fprintf(w, "\t%d\t%s\t%s\t%.2f\t\n", f.OrderID, f.Status, delivery, f.TotalPrice)
	}
	return w.Flush()
}

func (cmd *orderCommand) listStatus() (appie.FulfillmentStatus, string, error) {
	if cmd.Closed && cmd.All {
		return "", "", fmt.Errorf("--closed and --all cannot be used together")
	}
	if cmd.Closed {
		return appie.FulfillmentStatusClosed, "closed", nil
	}
	if cmd.All {
		return appie.FulfillmentStatusAll, "all", nil
	}
	return appie.FulfillmentStatusOpen, "open", nil
}

func findFulfillment(fulfillments []appie.Fulfillment, orderID string) *appie.Fulfillment {
	for i, f := range fulfillments {
		if strconv.Itoa(f.OrderID) == orderID {
			return &fulfillments[i]
		}
	}
	return nil
}

// ensureOrderOpen finds the fulfillment for orderID, validates it exists,
// reopens the order if SUBMITTED/CONFIRMED, and sets the client's active order ID.
func ensureOrderOpen(ctx context.Context, client *appie.Client, fulfillments []appie.Fulfillment, orderID int) error {
	var found *appie.Fulfillment
	for i, f := range fulfillments {
		if f.OrderID == orderID {
			found = &fulfillments[i]
			break
		}
	}

	if found == nil {
		return fmt.Errorf("order %d not found in open orders", orderID)
	}

	if found.Status == "SUBMITTED" || found.Status == "CONFIRMED" {
		if err := client.ReopenOrder(ctx, orderID); err != nil {
			return fmt.Errorf("failed to reopen order: %w", err)
		}
		fmt.Printf("Reopened order %d (was %s)\n", orderID, found.Status)
	}

	client.SetOrderID(orderID)
	return nil
}

func printOrder(order *appie.Order, f *appie.Fulfillment) error {
	fmt.Printf("Order %s  %s\n", order.ID, order.State)

	if f != nil {
		delivery := f.Delivery.Slot.DateDisplay
		if f.Delivery.Slot.TimeDisplay != "" {
			delivery += "  " + f.Delivery.Slot.TimeDisplay
		}
		fmt.Printf("Delivery: %s\n", delivery)
	}
	fmt.Println()

	if len(order.Items) == 0 {
		fmt.Println("No items")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, item := range order.Items {
		// Always show undiscounted price per line
		unitPrice := item.Product.Price.Now
		if item.Product.Price.Was > 0 {
			unitPrice = item.Product.Price.Was
		}
		linePrice := float64(item.Quantity) * unitPrice

		bonus := ""
		if item.Product.BonusMechanism != "" {
			bonus = "  " + item.Product.BonusMechanism
		}
		fmt.Fprintf(w, "  %d\t%s\t%s\t%d\t%6.2f%s\n", item.ProductID, item.Product.Title, item.Product.UnitSize, item.Quantity, linePrice, bonus)
	}

	// Use API-provided totals (from order summary or fulfillment)
	total := order.TotalPrice
	discount := order.TotalDiscount
	if total == 0 && f != nil && f.TotalPrice > 0 {
		total = f.TotalPrice
		subtotal := order.Subtotal()
		if subtotal > total {
			discount = subtotal - total
		}
	}

	fmt.Fprintf(w, "\t\t\t\t──────\n")
	if discount > 0 {
		fmt.Fprintf(w, "\t\t\t\t-%5.2f  bonus\n", discount)
	}
	fmt.Fprintf(w, "\t\t\t%d items\t%6.2f\n", len(order.Items), total)
	return w.Flush()
}

// show subcommand

type orderShowCommand struct {
	Args struct {
		OrderID int `positional-arg-name:"order-id" required:"true"`
	} `positional-args:"yes"`
}

func (cmd *orderShowCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	fulfillments, err := client.GetFulfillmentsByStatus(ctx, appie.FulfillmentStatusAll)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}

	orderID := cmd.Args.OrderID

	order, err := client.GetOrderDetails(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to get order details: %w", err)
	}

	f := findFulfillment(fulfillments, order.ID)
	if f == nil {
		closed, err := client.GetFulfillmentsByStatus(ctx, appie.FulfillmentStatusClosed)
		if err == nil {
			f = findFulfillment(closed, order.ID)
		}
	}

	// Try to get summary for totals on open orders. Delivered orders should use
	// their fulfillment total; the active summary can point at a different order.
	if f != nil && (f.Status == "DELIVERED" || f.Status == "CANCELLED") {
		order.TotalPrice = f.TotalPrice
		order.TotalDiscount = 0
	} else {
		client.SetOrderID(orderID)
		if summary, err := client.GetOrder(ctx); err == nil {
			order.TotalPrice = summary.TotalPrice
			order.TotalDiscount = summary.TotalDiscount
		}
	}

	return printOrder(order, f)
}

// add subcommand

type orderAddCommand struct {
	Args struct {
		OrderID int    `positional-arg-name:"order-id" required:"true"`
		Product string `positional-arg-name:"product" required:"true"`
	} `positional-args:"yes"`
	Quantity int `short:"n" long:"quantity" default:"1" description:"Quantity to add"`
}

func (cmd *orderAddCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	fulfillments, err := client.GetFulfillments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}

	orderID := cmd.Args.OrderID
	if err := ensureOrderOpen(ctx, client, fulfillments, orderID); err != nil {
		return err
	}

	product := cmd.Args.Product
	qty := cmd.Quantity

	// If numeric, use as product ID directly
	productID, err := strconv.Atoi(product)
	if err != nil {
		// Search for the product
		products, err := client.SearchProducts(ctx, product, 15)
		if err != nil {
			return fmt.Errorf("search failed: %w", err)
		}
		if len(products) == 0 {
			return fmt.Errorf("no products found for %q", product)
		}
		if len(products) > 1 {
			printProducts(products)
			return fmt.Errorf("multiple matches for %q, specify product ID", product)
		}
		productID = products[0].ID
		fmt.Printf("Found: %s\n", products[0].Title)
	}

	if err := client.AddToOrder(ctx, []appie.OrderItem{{ProductID: productID, Quantity: qty}}); err != nil {
		return err
	}

	fmt.Printf("Added %dx %d to order %d\n", qty, productID, orderID)
	return nil
}

// rm subcommand

type orderRmCommand struct {
	Args struct {
		OrderID   int `positional-arg-name:"order-id" required:"true"`
		ProductID int `positional-arg-name:"product-id" required:"true"`
	} `positional-args:"yes"`
}

func (cmd *orderRmCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	fulfillments, err := client.GetFulfillments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}

	orderID := cmd.Args.OrderID
	if err := ensureOrderOpen(ctx, client, fulfillments, orderID); err != nil {
		return err
	}

	productID := cmd.Args.ProductID
	if err := client.RemoveFromOrder(ctx, productID); err != nil {
		return err
	}

	fmt.Printf("Removed %d from order %d\n", productID, orderID)
	return nil
}

// submit subcommand

type orderSubmitCommand struct {
	Args struct {
		OrderID int `positional-arg-name:"order-id" required:"true"`
	} `positional-args:"yes"`
	Payment string `long:"payment" default:"auto" description:"Payment method: auto, dct, or pay-at-delivery"`
	Yes     bool   `long:"yes" description:"Actually submit the order; without this flag only validates readiness"`
}

func (cmd *orderSubmitCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	orderID := cmd.Args.OrderID
	info, err := client.GetOrderSubmissionInfo(ctx, orderID)
	if err != nil {
		return err
	}

	fulfillments, err := client.GetFulfillments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}
	fulfillment := findFulfillment(fulfillments, strconv.Itoa(orderID))

	paymentMethod, card, err := resolveSubmitPayment(ctx, client, cmd.Payment)
	if err != nil {
		return err
	}

	printSubmitReadiness(info, fulfillment, paymentMethod, card)

	if info.ValidationErrors > 0 || info.HasATPError {
		return fmt.Errorf("order %d has checkout validation errors", orderID)
	}
	if !info.ValueLimits.Submittable {
		return fmt.Errorf("order %d is not submittable", orderID)
	}
	if !cmd.Yes {
		fmt.Println()
		fmt.Println("Dry run only; rerun with --yes to submit.")
		return nil
	}

	opts := appie.OrderSubmitOptions{PaymentMethod: paymentMethod}
	if card != nil {
		opts.DCTCardID = card.CardID
	}
	result, err := client.SubmitOrder(ctx, orderID, opts)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Submitted order %d: %s\n", result.OrderID, result.OrderState)
	if len(result.PaymentStatuses) > 0 {
		fmt.Printf("Payment: %s\n", strings.Join(result.PaymentStatuses, ", "))
	}
	return nil
}

func resolveSubmitPayment(ctx context.Context, client *appie.Client, value string) (appie.PaymentMethod, *appie.DCTCard, error) {
	switch strings.ToLower(value) {
	case "", "auto", "dct":
		card, err := client.GetDefaultDCTCard(ctx)
		if err != nil {
			return "", nil, err
		}
		return appie.PaymentMethodDCT, card, nil
	case "pay-at-delivery":
		return appie.PaymentMethodPayAtDelivery, nil, nil
	default:
		return "", nil, fmt.Errorf("unsupported payment method %q", value)
	}
}

func printSubmitReadiness(info *appie.OrderSubmissionInfo, fulfillment *appie.Fulfillment, paymentMethod appie.PaymentMethod, card *appie.DCTCard) {
	fmt.Printf("Order %d  %s\n", info.OrderID, info.State)
	if fulfillment != nil {
		delivery := fulfillment.Delivery.Slot.DateDisplay
		if fulfillment.Delivery.Slot.TimeDisplay != "" {
			delivery += "  " + fulfillment.Delivery.Slot.TimeDisplay
		}
		fmt.Printf("Delivery: %s\n", delivery)
	}
	fmt.Printf("Total: %.2f\n", info.TotalPrice)
	fmt.Printf("Minimum: %.2f", info.ValueLimits.MinimumOrderValue.Amount)
	if info.ValueLimits.MinimumOrderValue.Deadline != "" {
		fmt.Printf(" by %s", info.ValueLimits.MinimumOrderValue.Deadline)
	}
	fmt.Printf(" (submittable: %t)\n", info.ValueLimits.Submittable)

	if info.ValidationErrors == 0 && !info.HasATPError {
		fmt.Println("Validation: ok")
	} else {
		fmt.Printf("Validation: %d errors, ATP error: %t\n", info.ValidationErrors, info.HasATPError)
		for _, validationError := range info.CheckoutErrors {
			label := validationError.Code
			if label == "" {
				label = validationError.TypeName
			}
			if validationError.Message != "" {
				fmt.Printf("  %s: %s\n", label, validationError.Message)
			} else if label != "" {
				fmt.Printf("  %s\n", label)
			}
		}
		if info.ATPError != nil {
			printCheckoutLimits("Stock limit", info.ATPError.StockLimits)
			printCheckoutLimits("Order limit", info.ATPError.OrderLimits)
		}
	}

	if paymentMethod == appie.PaymentMethodDCT && card != nil {
		label := card.CardAlias
		if label == "" {
			label = "default card"
		}
		if card.CardArtID != "" {
			label += " (" + card.CardArtID + ")"
		}
		fmt.Printf("Payment: DCT %s\n", label)
		return
	}
	fmt.Printf("Payment: %s\n", paymentMethod)
}

func printCheckoutLimits(label string, lines []appie.CheckoutOrderLine) {
	for _, line := range lines {
		productID := 0
		productTitle := "unknown product"
		productSize := ""
		if line.Product != nil {
			productID = line.Product.ID
			productTitle = line.Product.Title
			productSize = line.Product.UnitSize
		}
		if productSize != "" {
			productTitle += " " + productSize
		}
		fmt.Printf("  %s: %d %s (requested %d, available %d, type %s)\n",
			label, productID, productTitle, line.Count, line.Available, line.LimitType)
	}
}

// orderSetup creates an authenticated client and context.
func orderSetup() (context.Context, *appie.Client, error) {
	client, err := appie.NewWithConfig(globalOpts.Config, clientOpts()...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load config: %w", err)
	}

	if !client.IsAuthenticated() {
		return nil, nil, fmt.Errorf("not authenticated, run 'appie login' first")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	_ = cancel // cleaned up when process exits
	return ctx, client, nil
}
