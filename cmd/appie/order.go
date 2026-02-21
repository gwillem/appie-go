package main

import (
	"cmp"
	"context"
	"fmt"
	"os"
	"os/signal"
	"slices"
	"strconv"
	"text/tabwriter"

	appie "github.com/gwillem/appie-go"
)

type orderCommand struct {
	List listCommand `command:"list" alias:"ls" description:"List all open orders"`
	Use  useCommand  `command:"use" description:"Set a different order as active (reopens if submitted)"`
	Add  addCommand  `command:"add" description:"Add a product to the active order"`
}

func (cmd *orderCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	fulfillments, err := client.GetFulfillments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}

	// Try active order first (summary has totals but no per-item prices)
	var orderID int
	var summary *appie.Order
	if s, err := client.GetOrder(ctx); err == nil {
		summary = s
		orderID, _ = strconv.Atoi(s.ID)
	} else if len(fulfillments) > 0 {
		orderID = fulfillments[0].OrderID
	} else {
		fmt.Println("No open orders")
		return nil
	}

	// Get detailed line items with pricing
	order, err := client.GetOrderDetails(ctx, orderID)
	if err != nil {
		return fmt.Errorf("failed to get order details: %w", err)
	}

	// Merge discount info from summary (details endpoint lacks totals)
	if summary != nil {
		order.TotalPrice = summary.TotalPrice
		order.TotalDiscount = summary.TotalDiscount
	}

	f := findFulfillment(fulfillments, order.ID)
	return printOrder(order, f)
}

func findFulfillment(fulfillments []appie.Fulfillment, orderID string) *appie.Fulfillment {
	for i, f := range fulfillments {
		if strconv.Itoa(f.OrderID) == orderID {
			return &fulfillments[i]
		}
	}
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
	for i, item := range order.Items {
		// Always show undiscounted price per line
		unitPrice := item.Product.Price.Now
		if item.Product.Price.Was > 0 {
			unitPrice = item.Product.Price.Was
		}
		linePrice := float64(item.Quantity) * unitPrice

		fmt.Fprintf(w, "%2d\t%s\t%d\t%6.2f\t%s\n", i+1, item.Product.Title, item.Quantity, linePrice, item.Product.BonusMechanism)
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

	fmt.Fprintf(w, "\t\t\t──────\t\n")
	if discount > 0 {
		fmt.Fprintf(w, "\t\t\t-%5.2f\tbonus\n", discount)
	}
	fmt.Fprintf(w, "\t\t%d items\t%6.2f\t\n", len(order.Items), total)
	return w.Flush()
}

// list subcommand

type listCommand struct{}

func (cmd *listCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	fulfillments, err := client.GetFulfillments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}

	if len(fulfillments) == 0 {
		fmt.Println("No open orders")
		return nil
	}

	var activeID string
	if active, err := client.GetOrder(ctx); err == nil {
		activeID = active.ID
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	fmt.Fprintf(w, "\t%s\t%s\t%s\t%s\t\n", "Order", "Status", "Delivery", "Total")
	for _, f := range fulfillments {
		delivery := f.Delivery.Slot.DateDisplay
		if f.Delivery.Slot.TimeDisplay != "" {
			delivery += "  " + f.Delivery.Slot.TimeDisplay
		}
		marker := " "
		if strconv.Itoa(f.OrderID) == activeID {
			marker = "*"
		}
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%.2f\t\n", marker, f.OrderID, f.Status, delivery, f.TotalPrice)
	}
	return w.Flush()
}

// use subcommand

type useCommand struct {
	Args struct {
		OrderID int `positional-arg-name:"order-id" required:"true"`
	} `positional-args:"yes"`
}

func (cmd *useCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	orderID := cmd.Args.OrderID

	fulfillments, err := client.GetFulfillments(ctx)
	if err != nil {
		return fmt.Errorf("failed to get orders: %w", err)
	}

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

	fmt.Printf("Active order: %d\n", orderID)
	return nil
}

// add subcommand

type addCommand struct {
	Args struct {
		Product string `positional-arg-name:"product" required:"true"`
	} `positional-args:"yes"`
	Quantity int `short:"n" long:"quantity" default:"1" description:"Quantity to add"`
}

func (cmd *addCommand) Execute(args []string) error {
	ctx, client, err := orderSetup()
	if err != nil {
		return err
	}

	// Fetch active order to populate order ID header
	if _, err := client.GetOrder(ctx); err != nil {
		return fmt.Errorf("no active order: %w", err)
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
			slices.SortFunc(products, func(a, b appie.Product) int {
				return cmp.Compare(a.Price.Now, b.Price.Now)
			})
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			for _, p := range products {
				fmt.Fprintf(w, "  %d\t%s\t%s\t€%.2f\n", p.ID, p.Title, p.UnitSize, p.Price.Now)
			}
			w.Flush()
			return fmt.Errorf("multiple matches for %q, specify product ID", product)
		}
		productID = products[0].ID
		fmt.Printf("Found: %s\n", products[0].Title)
	}

	if err := client.AddToOrder(ctx, []appie.OrderItem{{ProductID: productID, Quantity: qty}}); err != nil {
		return err
	}

	fmt.Printf("Added %dx %d\n", qty, productID)
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
