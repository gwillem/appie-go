package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	appie "github.com/gwillem/appie-go"
)

type receiptCommand struct{}

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

	receipts, err := client.GetReceipts(ctx)
	if err != nil {
		return fmt.Errorf("failed to get receipts: %w", err)
	}

	if len(receipts) == 0 {
		fmt.Println("No receipts found")
		return nil
	}

	limit := min(20, len(receipts))
	for _, r := range receipts[:limit] {
		fmt.Printf("%s  %-30s  €%.2f\n", r.Date, r.StoreName, r.TotalAmount)
	}

	return nil
}
