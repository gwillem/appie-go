package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	appie "github.com/gwillem/appie-go"
)

// trimMillis strips sub-second precision from a timestamp string (e.g. ".000" in "T14:30:00.000").
func trimMillis(s string) string {
	if i := strings.IndexByte(s, '.'); i != -1 {
		return s[:i]
	}
	return s
}

type receiptCommand struct {
	N    int `short:"n" default:"20" description:"Number of recent receipts to show"`
	Args struct {
		TransactionID string `positional-arg-name:"transaction-id"`
	} `positional-args:"yes"`
}

func (cmd *receiptCommand) Execute(args []string) error {
	client, err := appie.NewWithConfig(globalOpts.Config)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if !client.IsAuthenticated() {
		return fmt.Errorf("not authenticated, run 'appie login' first")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if cmd.Args.TransactionID != "" {
		return showReceipt(ctx, client, cmd.Args.TransactionID)
	}
	return listReceipts(ctx, client, cmd.N)
}

func listReceipts(ctx context.Context, client *appie.Client, n int) error {
	receipts, err := client.GetReceipts(ctx)
	if err != nil {
		return fmt.Errorf("failed to get receipts: %w", err)
	}

	if len(receipts) == 0 {
		fmt.Println("No receipts found")
		return nil
	}

	limit := min(n, len(receipts))
	for _, r := range receipts[:limit] {
		fmt.Printf("%-20s %s %6.2f\n", r.TransactionID, trimMillis(r.Date), r.TotalAmount)
	}

	return nil
}

func showReceipt(ctx context.Context, client *appie.Client, id string) error {
	receipts, err := client.GetReceipts(ctx)
	if err != nil {
		return fmt.Errorf("failed to get receipts: %w", err)
	}

	var meta *appie.Receipt
	for i, r := range receipts {
		if r.TransactionID == id {
			meta = &receipts[i]
			break
		}
	}

	receipt, err := client.GetReceipt(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get receipt: %w", err)
	}

	fmt.Printf("Receipt %s\n", receipt.TransactionID)
	if meta != nil {
		fmt.Printf("Date:  %s\n", trimMillis(meta.Date))
	}
	fmt.Println()

	var subtotal float64
	for _, item := range receipt.Items {
		subtotal += item.Amount
		if item.Quantity > 1 {
			fmt.Printf("  %dx %-30s %6.2f\n", item.Quantity, item.Description, item.Amount)
		} else {
			fmt.Printf("     %-30s %6.2f\n", item.Description, item.Amount)
		}
	}

	if len(receipt.Discounts) > 0 {
		fmt.Println()
		for _, d := range receipt.Discounts {
			subtotal += d.Amount
			fmt.Printf("     %-30s %6.2f\n", d.Name, d.Amount)
		}
	}

	fmt.Printf("     %-30s ------\n", "")
	for _, p := range receipt.Payments {
		fmt.Printf("     %-30s %6.2f\n", p.Method, p.Amount)
	}

	return nil
}
