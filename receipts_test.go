package appie

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetReceipts(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		var req graphQLRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatal(err)
		}

		resp := graphQLResponse[json.RawMessage]{
			Data: json.RawMessage(`{
				"posReceiptsPage": {
					"posReceipts": [
						{
							"id": "txn-001",
							"dateTime": "2025-01-15T14:30:00",
							"totalAmount": {"amount": 42.50, "formattedV2": "€ 42,50"}
						},
						{
							"id": "txn-002",
							"dateTime": "2025-01-10T09:15:00",
							"totalAmount": {"amount": 18.99, "formattedV2": "€ 18,99"}
						}
					]
				}
			}`),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL))
	receipts, err := client.GetReceipts(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(receipts) != 2 {
		t.Fatalf("got %d receipts, want 2", len(receipts))
	}

	if receipts[0].TransactionID != "txn-001" {
		t.Errorf("got transaction ID %q, want %q", receipts[0].TransactionID, "txn-001")
	}
	if receipts[0].Date != "2025-01-15T14:30:00" {
		t.Errorf("got date %q, want %q", receipts[0].Date, "2025-01-15T14:30:00")
	}
	if receipts[0].TotalAmount != 42.50 {
		t.Errorf("got total %.2f, want 42.50", receipts[0].TotalAmount)
	}
}

func TestGetReceipt(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		var req graphQLRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatal(err)
		}

		resp := graphQLResponse[json.RawMessage]{
			Data: json.RawMessage(`{
				"posReceiptDetails": {
					"id": "txn-001",
					"products": [
						{
							"id": 12345,
							"quantity": 2,
							"name": "AH Halfvolle melk",
							"price": {"amount": 1.29},
							"amount": {"amount": 2.58}
						},
						{
							"id": 67890,
							"quantity": 1,
							"name": "AH Pindakaas",
							"price": null,
							"amount": {"amount": 3.49}
						}
					],
					"discounts": [
						{
							"name": "BONUS MELK",
							"amount": {"amount": -1.00}
						}
					],
					"payments": [
						{
							"method": "PINNEN",
							"amount": {"amount": 5.07}
						}
					]
				}
			}`),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL))
	receipt, err := client.GetReceipt(context.Background(), "txn-001")
	if err != nil {
		t.Fatal(err)
	}

	if receipt.TransactionID != "txn-001" {
		t.Errorf("got transaction ID %q, want %q", receipt.TransactionID, "txn-001")
	}
	if len(receipt.Items) != 2 {
		t.Fatalf("got %d items, want 2", len(receipt.Items))
	}
	if receipt.Items[0].Description != "AH Halfvolle melk" {
		t.Errorf("got description %q, want %q", receipt.Items[0].Description, "AH Halfvolle melk")
	}
	if receipt.Items[0].Quantity != 2 {
		t.Errorf("got quantity %d, want 2", receipt.Items[0].Quantity)
	}
	if receipt.Items[0].Amount != 2.58 {
		t.Errorf("got amount %.2f, want 2.58", receipt.Items[0].Amount)
	}
	if receipt.Items[0].UnitPrice != 1.29 {
		t.Errorf("got unit price %.2f, want 1.29", receipt.Items[0].UnitPrice)
	}
	if receipt.Items[1].UnitPrice != 0 {
		t.Errorf("got unit price %.2f for null price, want 0", receipt.Items[1].UnitPrice)
	}

	if len(receipt.Discounts) != 1 {
		t.Fatalf("got %d discounts, want 1", len(receipt.Discounts))
	}
	if receipt.Discounts[0].Name != "BONUS MELK" {
		t.Errorf("got discount name %q, want %q", receipt.Discounts[0].Name, "BONUS MELK")
	}
	if receipt.Discounts[0].Amount != -1.00 {
		t.Errorf("got discount amount %.2f, want -1.00", receipt.Discounts[0].Amount)
	}

	if len(receipt.Payments) != 1 {
		t.Fatalf("got %d payments, want 1", len(receipt.Payments))
	}
	if receipt.Payments[0].Method != "PINNEN" {
		t.Errorf("got payment method %q, want %q", receipt.Payments[0].Method, "PINNEN")
	}
	if receipt.Payments[0].Amount != 5.07 {
		t.Errorf("got payment amount %.2f, want 5.07", receipt.Payments[0].Amount)
	}
}
