package appie

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

// receiptDetailsFixture is the GraphQL data block returned by the
// posReceiptDetails query in the resolver-aware tests below.
const receiptDetailsFixture = `{
	"posReceiptDetails": {
		"id": "txn-001",
		"products": [
			{"id": 767898, "quantity": 3, "name": "COCA-COLA",   "price": {"amount": 2.49}, "amount": {"amount": 7.47}},
			{"id": 823326, "quantity": 1, "name": "AH KWARK",    "price": {"amount": 1.49}, "amount": {"amount": 1.49}},
			{"id": 552022, "quantity": 1, "name": "AH ZALMFILET","price": {"amount": 5.99}, "amount": {"amount": 5.99}},
			{"id": 999999, "quantity": 1, "name": "UNRESOLVABLE","price": {"amount": 0.99}, "amount": {"amount": 0.99}}
		],
		"discounts": [],
		"payments":  []
	}
}`

// readGraphQLRequest reads the request body once and returns it both
// decoded as a graphQLRequest and as the original raw string, so the
// handler can switch on the query while still asserting on the wire form.
// It does not rewind r.Body — callers must not read it again.
func readGraphQLRequest(t *testing.T, r *http.Request) (graphQLRequest, string) {
	t.Helper()
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	var req graphQLRequest
	if err := json.Unmarshal(raw, &req); err != nil {
		t.Fatalf("decode body: %v", err)
	}
	return req, string(raw)
}

func TestGetReceiptResolvesWebshopID(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		req, _ := readGraphQLRequest(t, r)

		var data string
		switch {
		case strings.Contains(req.Query, "posReceiptDetails"):
			data = receiptDetailsFixture
		case strings.Contains(req.Query, "productConvertId"):
			// Mirror the production behaviour: 999999 has no
			// conversion (sentinel -1); the others resolve.
			data = `{"p0":171607,"p1":231878,"p2":121453,"p3":-1}`
		default:
			t.Fatalf("unexpected query: %s", req.Query)
		}
		json.NewEncoder(w).Encode(graphQLResponse[json.RawMessage]{Data: json.RawMessage(data)})
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL))
	receipt, err := client.GetReceipt(context.Background(), "txn-001")
	if err != nil {
		t.Fatal(err)
	}

	want := map[int]int{767898: 171607, 823326: 231878, 552022: 121453, 999999: 0}
	for _, item := range receipt.Items {
		if got := item.WebshopID; got != want[item.ProductID] {
			t.Errorf("ProductID %d: got WebshopID %d, want %d", item.ProductID, got, want[item.ProductID])
		}
	}
}

func TestConvertPOSIDsBatchShape(t *testing.T) {
	var captured string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, raw := readGraphQLRequest(t, r)
		captured = raw
		json.NewEncoder(w).Encode(graphQLResponse[json.RawMessage]{
			Data: json.RawMessage(`{"p0":171607,"p1":231878}`),
		})
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL))
	// Includes a duplicate (767898 twice) and a non-positive id (0)
	// that must both be filtered out before the request is built.
	mapping, err := client.ConvertPOSIDs(context.Background(), []int{767898, 823326, 767898, 0})
	if err != nil {
		t.Fatal(err)
	}

	if mapping[767898] != 171607 || mapping[823326] != 231878 {
		t.Errorf("unexpected mapping: %v", mapping)
	}

	// Exactly two aliased subfields (one per unique positive id), in
	// input order, no duplicates, no zero.
	wantFragments := []string{
		"p0: productConvertId(sourceId: 767898)",
		"p1: productConvertId(sourceId: 823326)",
	}
	for _, f := range wantFragments {
		if !strings.Contains(captured, f) {
			t.Errorf("expected query to contain %q, got %s", f, captured)
		}
	}
	if strings.Contains(captured, "p2:") {
		t.Errorf("expected exactly 2 aliases, got %s", captured)
	}
	if strings.Contains(captured, "sourceId: 0") {
		t.Errorf("zero id should have been filtered, got %s", captured)
	}
}

func TestConvertPOSIDsTreatsNullAsUnresolved(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		readGraphQLRequest(t, r)
		// productConvertId is `Int` (nullable) in the schema. A null
		// alias must not break the whole batch — it should fall back
		// to the same "unresolved" path as -1.
		json.NewEncoder(w).Encode(graphQLResponse[json.RawMessage]{
			Data: json.RawMessage(`{"p0":171607,"p1":null,"p2":-1}`),
		})
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL))
	mapping, err := client.ConvertPOSIDs(context.Background(), []int{767898, 111111, 222222})
	if err != nil {
		t.Fatalf("null alias should not produce an error, got: %v", err)
	}
	if mapping[767898] != 171607 {
		t.Errorf("expected 767898 -> 171607, got %d", mapping[767898])
	}
	if _, ok := mapping[111111]; ok {
		t.Errorf("null alias should be omitted from mapping, got %d", mapping[111111])
	}
	if _, ok := mapping[222222]; ok {
		t.Errorf("-1 sentinel should be omitted from mapping, got %d", mapping[222222])
	}
}

func TestConvertPOSIDsErrorIsBestEffort(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := readGraphQLRequest(t, r)
		if strings.Contains(req.Query, "posReceiptDetails") {
			json.NewEncoder(w).Encode(graphQLResponse[json.RawMessage]{
				Data: json.RawMessage(receiptDetailsFixture),
			})
			return
		}
		// Convert call fails — GetReceipt must still succeed.
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL))
	receipt, err := client.GetReceipt(context.Background(), "txn-001")
	if err != nil {
		t.Fatalf("GetReceipt should be best-effort, got error: %v", err)
	}
	if len(receipt.Items) != 4 {
		t.Fatalf("got %d items, want 4", len(receipt.Items))
	}
	for _, item := range receipt.Items {
		if item.WebshopID != 0 {
			t.Errorf("expected WebshopID=0 on convert failure, got %d for ProductID %d", item.WebshopID, item.ProductID)
		}
	}

	// Direct caller of ConvertPOSIDs that ignores err must not panic
	// or nil-deref on the returned map; it should be safely readable.
	mapping, err := client.ConvertPOSIDs(context.Background(), []int{1, 2})
	if err == nil {
		t.Fatalf("expected error from failing convert call")
	}
	if mapping == nil {
		t.Fatalf("ConvertPOSIDs returned nil map alongside error; should be non-nil for safe lookup")
	}
	if v := mapping[1]; v != 0 {
		t.Errorf("expected 0 lookup on empty map, got %d", v)
	}
}
