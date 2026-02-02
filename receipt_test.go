//go:build integration

package appie

import (
	"context"
	"testing"
)

func TestGetReceipts(t *testing.T) {
	t.Skip("TODO: receipt backend service temporarily unavailable (503)")

	client := testClient(t)
	ctx := context.Background()

	receipts, err := client.GetReceipts(ctx)
	if err != nil {
		t.Fatalf("failed to get receipts: %v", err)
	}

	t.Logf("Found %d receipts", len(receipts))
	for i, r := range receipts {
		if i >= 5 {
			t.Logf("  ... and %d more", len(receipts)-5)
			break
		}
		t.Logf("  - %s: %s at %s - %.2f EUR", r.TransactionID, r.Date, r.StoreName, r.TotalAmount)
	}
}

func TestGetReceipt(t *testing.T) {
	t.Skip("TODO: receipt backend service temporarily unavailable (503)")

	client := testClient(t)
	ctx := context.Background()

	// First get the list to find a valid transaction ID
	receipts, err := client.GetReceipts(ctx)
	if err != nil {
		t.Fatalf("failed to get receipts: %v", err)
	}

	if len(receipts) == 0 {
		t.Skip("no receipts available to test")
	}

	transactionID := receipts[0].TransactionID
	t.Logf("Fetching receipt details for transaction: %s", transactionID)

	receipt, err := client.GetReceipt(ctx, transactionID)
	if err != nil {
		t.Fatalf("failed to get receipt %s: %v", transactionID, err)
	}

	t.Logf("Receipt: %s", receipt.TransactionID)
	t.Logf("  Date: %s", receipt.Date)
	t.Logf("  Store: %s", receipt.StoreName)
	t.Logf("  Total: %.2f EUR", receipt.TotalAmount)
	t.Logf("  Items: %d", len(receipt.Items))

	for i, item := range receipt.Items {
		if i >= 5 {
			t.Logf("  ... and %d more items", len(receipt.Items)-5)
			break
		}
		t.Logf("    - %s: %.2f EUR (qty: %d)", item.Description, item.Amount, item.Quantity)
	}
}
