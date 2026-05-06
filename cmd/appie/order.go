package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"text/tabwriter"
	"time"

	appie "github.com/gwillem/appie-go"
)

type orderCommand struct {
	List orderListCommand `command:"list" description:"List orders (open by default)"`
	Show orderShowCommand `command:"show" description:"Show contents of an order"`
	Add  orderAddCommand  `command:"add" description:"Add a product to an order"`
	Rm   orderRmCommand   `command:"rm" description:"Remove a product from an order"`
}

func (cmd *orderCommand) Execute(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("unknown argument %q, did you mean: appie order show %s", args[0], args[0])
	}
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	fulfillments, err := client.GetFulfillments(ctx, appie.FulfillmentOpen, 0)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}

	if len(fulfillments) == 0 {
		fmt.Println("No open orders")
		return nil
	}

	printFulfillmentList(fulfillments)
	return nil
}

func printFulfillmentList(fulfillments []appie.Fulfillment) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	fmt.Fprintf(w, "\t%s\t%s\t%s\t%s\t\n", "Order", "Status", "Delivery", "Total")
	for _, f := range fulfillments {
		delivery := formatNLDate(f.Delivery.Slot.Date)
		if f.Delivery.Slot.StartTime != "" {
			delivery += fmt.Sprintf("  %s-%s", f.Delivery.Slot.StartTime, f.Delivery.Slot.EndTime)
		}
		fmt.Fprintf(w, "\t%d\t%s\t%s\t%.2f\t\n", f.OrderID, f.Status, delivery, f.TotalPrice)
	}
	w.Flush()
}

// formatNLDate formats an ISO date (2006-01-02) as Dutch long form: "dinsdag 15 april 2025".
func formatNLDate(iso string) string {
	t, err := time.Parse("2006-01-02", iso)
	if err != nil {
		return iso
	}
	days := [...]string{"zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"}
	months := [...]string{"", "januari", "februari", "maart", "april", "mei", "juni",
		"juli", "augustus", "september", "oktober", "november", "december"}
	return fmt.Sprintf("%s %d %s %d", days[t.Weekday()], t.Day(), months[t.Month()], t.Year())
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

// list subcommand

type orderListCommand struct {
	Open   bool `long:"open" description:"Show open orders (default)"`
	Closed bool `long:"closed" description:"Show closed/delivered orders"`
	All    bool `long:"all" description:"Show both open and closed orders"`
	N      int  `short:"n" default:"25" description:"Number of orders to show"`
}

func (cmd *orderListCommand) Execute(args []string) error {
	count := 0
	for _, b := range []bool{cmd.Open, cmd.Closed, cmd.All} {
		if b {
			count++
		}
	}
	if count > 1 {
		return fmt.Errorf("--open, --closed, and --all are mutually exclusive")
	}

	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	if cmd.All {
		open, err := client.GetFulfillments(ctx, appie.FulfillmentOpen, cmd.N)
		if err != nil {
			return fmt.Errorf("failed to get open orders: %w", err)
		}
		closed, err := client.GetFulfillments(ctx, appie.FulfillmentClosed, cmd.N)
		if err != nil {
			return fmt.Errorf("failed to get closed orders: %w", err)
		}
		all := append(open, closed...)
		if len(all) == 0 {
			fmt.Println("No orders found")
			return nil
		}
		printFulfillmentList(all)
		return nil
	}

	status := appie.FulfillmentOpen
	emptyMsg := "No open orders"
	if cmd.Closed {
		status = appie.FulfillmentClosed
		emptyMsg = "No closed orders"
	}

	fulfillments, err := client.GetFulfillments(ctx, status, cmd.N)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}
	if len(fulfillments) == 0 {
		fmt.Println(emptyMsg)
		return nil
	}
	printFulfillmentList(fulfillments)
	return nil
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

	detail, err := client.GetFulfillmentDetail(ctx, cmd.Args.OrderID)
	if err != nil {
		// Fall back to REST order detail for open/active orders
		return showOrderViaREST(ctx, client, cmd.Args.OrderID)
	}

	printFulfillmentDetail(detail)
	return nil
}

func showOrderViaREST(ctx context.Context, client *appie.Client, orderID int) error {
	fulfillments, err := client.GetFulfillments(ctx, appie.FulfillmentOpen, 0)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}

	order, err := client.GetOrderDetails(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to get order details: %w", err)
	}

	client.SetOrderID(orderID)
	if summary, err := client.GetOrder(ctx); err == nil {
		order.TotalPrice = summary.TotalPrice
		order.TotalDiscount = summary.TotalDiscount
	}

	f := findFulfillment(fulfillments, order.ID)
	return printOrder(order, f)
}

func printFulfillmentDetail(detail *appie.FulfillmentDetail) {
	fmt.Printf("Order %d", detail.OrderID)
	if detail.ClosingDateTime != "" {
		fmt.Printf("  (closed %s)", detail.ClosingDateTime)
	} else if detail.Delivery.Status != "" {
		fmt.Printf("  %s", detail.Delivery.Status)
	}
	fmt.Println()

	if detail.Delivery.Slot.Date != "" {
		delivery := formatNLDate(detail.Delivery.Slot.Date)
		if detail.Delivery.Slot.StartTime != "" {
			delivery += fmt.Sprintf("  %s - %s", detail.Delivery.Slot.StartTime, detail.Delivery.Slot.EndTime)
		}
		fmt.Printf("Delivery: %s  %s\n", detail.Delivery.Method, delivery)

		if detail.Delivery.Address.Street != "" {
			fmt.Printf("Address:  %s %d%s, %s %s\n",
				detail.Delivery.Address.Street,
				detail.Delivery.Address.HouseNumber,
				detail.Delivery.Address.HouseNumberExtra,
				detail.Delivery.Address.PostalCode,
				detail.Delivery.Address.City,
			)
		}
		if detail.InvoiceID != "" {
			fmt.Printf("Invoice:  %s\n", detail.InvoiceID)
		}
	}
	fmt.Println()

	if len(detail.OrderLines) == 0 {
		fmt.Println("No items")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	var total float64
	for _, ol := range detail.OrderLines {
		if ol.Product == nil {
			continue
		}
		linePrice := float64(ol.AllocatedQuantity) * ol.Product.CurrentPrice
		total += linePrice

		qty := ol.AllocatedQuantity
		if qty != ol.Quantity {
			fmt.Fprintf(w, "  %d\t%s\t%s\t%d/%d\t%6.2f\n",
				ol.Product.ID, ol.Product.Title, ol.Product.SalesUnitSize,
				qty, ol.Quantity, linePrice)
		} else {
			fmt.Fprintf(w, "  %d\t%s\t%s\t%d\t%6.2f\n",
				ol.Product.ID, ol.Product.Title, ol.Product.SalesUnitSize,
				qty, linePrice)
		}
	}

	fmt.Fprintf(w, "\t\t\t\t──────\n")
	fmt.Fprintf(w, "\t\t\t%d items\t%6.2f\n", len(detail.OrderLines), total)
	w.Flush()
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

	fulfillments, err := client.GetFulfillments(ctx, appie.FulfillmentOpen, 0)
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

	fulfillments, err := client.GetFulfillments(ctx, appie.FulfillmentOpen, 0)
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
