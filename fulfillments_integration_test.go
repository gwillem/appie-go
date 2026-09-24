//go:build integration

package appie

import (
	"context"
	"testing"
)

// TestGetFulfillmentsOpenGraphQL verifies that the default call (no options)
// reaches the GraphQL endpoint and that the restored fields are populated.
func TestGetFulfillmentsOpenGraphQL(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	fulfillments, err := client.GetFulfillments(ctx)
	if err != nil {
		t.Fatalf("GetFulfillments failed: %v", err)
	}

	t.Logf("Found %d open fulfillment(s)", len(fulfillments))
	for _, f := range fulfillments {
		t.Logf("  order %d: status=%q transactionCompleted=%v modifiable=%v",
			f.OrderID, f.Status, f.TransactionCompleted, f.Modifiable)

		if f.OrderID == 0 {
			t.Error("expected non-zero OrderID")
		}
	}
}

// TestGetFulfillmentsClosedGraphQL verifies the CLOSED status path reaches GraphQL
// and that the restored fields contain real values.
func TestGetFulfillmentsClosedGraphQL(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	fulfillments, err := client.GetFulfillments(ctx, WithStatus(FulfillmentClosed), WithSize(5))
	if err != nil {
		t.Fatalf("GetFulfillments(CLOSED) failed: %v", err)
	}
	if len(fulfillments) == 0 {
		t.Skip("no closed fulfillments found; cannot verify restored fields")
	}

	t.Logf("Found %d closed fulfillment(s) (showing up to 5)", len(fulfillments))
	for _, f := range fulfillments {
		t.Logf("  order %d: status=%q transactionCompleted=%v modifiable=%v reopenable=%v",
			f.OrderID, f.Status, f.TransactionCompleted, f.Modifiable, f.Reopenable)

		if f.OrderID == 0 {
			t.Error("expected non-zero OrderID")
		}
		if f.TotalPrice == 0 {
			t.Errorf("order %d: expected non-zero TotalPrice for a closed order", f.OrderID)
		}
		// transactionCompleted and modifiable are returned but their semantics
		// are AH-internal; log rather than assert specific values.
	}
}

// TestGetFulfillmentDetailIntegration verifies that GetFulfillmentDetail returns
// a fully populated order including delivery address and order lines.
func TestGetFulfillmentDetailIntegration(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// Find a closed order to use as the test subject.
	fulfillments, err := client.GetFulfillments(ctx, WithStatus(FulfillmentClosed), WithSize(1))
	if err != nil {
		t.Fatalf("GetFulfillments(CLOSED) failed: %v", err)
	}
	if len(fulfillments) == 0 {
		t.Skip("no closed fulfillments available for detail test")
	}

	orderID := fulfillments[0].OrderID
	t.Logf("Fetching detail for order %d", orderID)

	detail, err := client.GetFulfillmentDetail(ctx, orderID)
	if err != nil {
		t.Fatalf("GetFulfillmentDetail(%d) failed: %v", orderID, err)
	}

	if detail.OrderID != orderID {
		t.Errorf("got OrderID %d, want %d", detail.OrderID, orderID)
	}
	if detail.Delivery.Slot.Date == "" {
		t.Error("expected non-empty Delivery.Slot.Date")
	}
	if detail.Delivery.Address.PostalCode == "" {
		t.Error("expected non-empty Delivery.Address.PostalCode")
	}
	if len(detail.OrderLines) == 0 {
		t.Error("expected at least one order line")
	}

	t.Logf("  delivery: %s %s-%s", detail.Delivery.Slot.Date, detail.Delivery.Slot.StartTime, detail.Delivery.Slot.EndTime)
	t.Logf("  address: %s %d, %s", detail.Delivery.Address.Street, detail.Delivery.Address.HouseNumber, detail.Delivery.Address.City)
	t.Logf("  invoice: %s, order lines: %d", detail.InvoiceID, len(detail.OrderLines))

	var withProduct, withoutProduct int
	for _, ol := range detail.OrderLines {
		if ol.Product != nil {
			withProduct++
		} else {
			withoutProduct++
		}
	}
	t.Logf("  lines with product: %d, without: %d", withProduct, withoutProduct)
}
