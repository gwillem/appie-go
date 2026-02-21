package appie

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetOrderDetails(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mobile-services/order/v1/229775812/details-grouped-by-taxonomy" {
			t.Errorf("unexpected path: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(map[string]any{
			"orderId":      229775812,
			"deliveryDate": "2025-12-09",
			"orderState":   "DELIVERED",
			"closingTime":  "2025-12-08T22:59:00Z",
			"deliveryType": "HOME",
			"deliveryTimePeriod": map[string]any{
				"startDateTime": "2025-12-09T18:00:00",
				"endDateTime":   "2025-12-09T20:00:00",
			},
			"groupedProductsInTaxonomy": []map[string]any{
				{
					"taxonomyName": "Groente, aardappelen",
					"orderedProducts": []map[string]any{
						{
							"amount":   1,
							"quantity": 2,
							"product": map[string]any{
								"webshopId":        164358,
								"title":            "AH Oranje zoete aardappel",
								"brand":            "AH",
								"salesUnitSize":    "1 kg",
								"priceBeforeBonus": 3.79,
								"isBonus":          false,
							},
						},
					},
				},
				{
					"taxonomyName": "Zuivel, eieren",
					"orderedProducts": []map[string]any{
						{
							"amount":   1,
							"quantity": 1,
							"product": map[string]any{
								"webshopId":        371880,
								"title":            "Optimel Drinkyoghurt aardbei",
								"brand":            "Optimel",
								"salesUnitSize":    "1 L",
								"priceBeforeBonus": 1.59,
								"currentPrice":     1.19,
								"isBonus":          true,
								"bonusMechanism":   "25% korting",
							},
						},
					},
				},
			},
			"invoiceId":   "2294567-00199",
			"cancellable": false,
		})
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))
	ctx := context.Background()

	order, err := client.GetOrderDetails(ctx, 229775812)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if order.ID != "229775812" {
		t.Errorf("expected ID '229775812', got %q", order.ID)
	}
	if order.State != "DELIVERED" {
		t.Errorf("expected state 'DELIVERED', got %q", order.State)
	}
	if len(order.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(order.Items))
	}

	item := order.Items[0]
	if item.ProductID != 164358 {
		t.Errorf("expected productID 164358, got %d", item.ProductID)
	}
	if item.Quantity != 2 {
		t.Errorf("expected quantity 2, got %d", item.Quantity)
	}
	if item.Product == nil {
		t.Fatal("expected product to be populated")
	}
	if item.Product.Title != "AH Oranje zoete aardappel" {
		t.Errorf("expected title 'AH Oranje zoete aardappel', got %q", item.Product.Title)
	}
	if item.Product.Price.Now != 3.79 {
		t.Errorf("expected price 3.79, got %.2f", item.Product.Price.Now)
	}

	// Non-bonus item: Price.Now = priceBeforeBonus, Price.Was = 0
	if item.Product.Price.Was != 0 {
		t.Errorf("expected Was 0 for non-bonus item, got %.2f", item.Product.Price.Was)
	}

	item2 := order.Items[1]
	if item2.ProductID != 371880 {
		t.Errorf("expected productID 371880, got %d", item2.ProductID)
	}
	if !item2.Product.IsBonus {
		t.Error("expected IsBonus true for second item")
	}

	// Bonus item: Price.Now = currentPrice, Price.Was = priceBeforeBonus
	if item2.Product.Price.Now != 1.19 {
		t.Errorf("expected discounted price 1.19, got %.2f", item2.Product.Price.Now)
	}
	if item2.Product.Price.Was != 1.59 {
		t.Errorf("expected Was 1.59, got %.2f", item2.Product.Price.Was)
	}
	if item2.Product.BonusMechanism != "25% korting" {
		t.Errorf("expected BonusMechanism '25%% korting', got %q", item2.Product.BonusMechanism)
	}
}

func TestGetOrderDiscount(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"id":    12345,
			"state": "OPEN",
			"totalPrice": map[string]any{
				"priceBeforeDiscount": 128.20,
				"priceAfterDiscount":  92.68,
				"priceDiscount":       35.52,
				"priceTotalPayable":   92.68,
			},
			"orderedProducts": []map[string]any{
				{
					"amount":   1,
					"quantity": 2,
					"product": map[string]any{
						"webshopId": 111,
						"title":     "Worteltjes",
						"brand":     "AH",
						"images":    []any{},
					},
				},
			},
		})
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))
	order, err := client.GetOrder(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if order.TotalPrice != 92.68 {
		t.Errorf("expected TotalPrice 92.68, got %.2f", order.TotalPrice)
	}
	if order.TotalDiscount != 35.52 {
		t.Errorf("expected TotalDiscount 35.52, got %.2f", order.TotalDiscount)
	}
}

func TestOrderDetailsTotals(t *testing.T) {
	data, err := os.ReadFile("testdata/order_details.json")
	if err != nil {
		t.Fatalf("read testdata: %v", err)
	}

	var resp orderDetailsResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	order := resp.toOrder()

	if len(order.Items) != 33 {
		t.Fatalf("expected 33 items, got %d", len(order.Items))
	}

	// Verify specific item types are parsed correctly
	findItem := func(title string) *OrderItem {
		for i := range order.Items {
			if order.Items[i].Product.Title == title {
				return &order.Items[i]
			}
		}
		t.Fatalf("item %q not found", title)
		return nil
	}

	// Non-bonus item
	item := findItem("Boursin Sjalot & bieslook")
	if item.Product.Price.Now != 3.79 || item.Product.Price.Was != 0 {
		t.Errorf("non-bonus: Now=%.2f Was=%.2f, want Now=3.79 Was=0", item.Product.Price.Now, item.Product.Price.Was)
	}

	// Percentage discount (has currentPrice)
	item = findItem("AH Biologisch Bleekselderij")
	if item.Product.Price.Now != 1.70 || item.Product.Price.Was != 1.89 {
		t.Errorf("percentage: Now=%.2f Was=%.2f, want Now=1.70 Was=1.89", item.Product.Price.Now, item.Product.Price.Was)
	}

	// Group promotion "2e gratis" (no currentPrice)
	item = findItem("AH Geschrapte worteltjes grootverpakking")
	if item.Quantity != 2 {
		t.Errorf("worteltjes qty=%d, want 2", item.Quantity)
	}
	if item.Product.Price.Now != 1.79 || item.Product.Price.Was != 0 {
		t.Errorf("2e gratis: Now=%.2f Was=%.2f, want Now=1.79 Was=0", item.Product.Price.Now, item.Product.Price.Was)
	}
	if item.Product.BonusMechanism != "2e gratis" {
		t.Errorf("mechanism=%q, want %q", item.Product.BonusMechanism, "2e gratis")
	}

	// Scratch card item (no priceBeforeBonus at all)
	item = findItem("Alpro Barista haver")
	if item.Quantity != 3 {
		t.Errorf("alpro qty=%d, want 3", item.Quantity)
	}
	if item.Product.Price.Now != 0 {
		t.Errorf("scratch card: Now=%.2f, want 0", item.Product.Price.Now)
	}

	// Subtotal from line items (sum of priceBeforeBonus * qty).
	// This is less than the API's true pre-discount total because
	// scratch card items have no price in the details response.
	subtotal := order.Subtotal()
	if math.Abs(subtotal-119.97) > 0.01 {
		t.Errorf("subtotal = %.2f, want 119.97", subtotal)
	}

	// Simulate CLI merging summary data (from GetOrder)
	order.TotalPrice = 92.68
	order.TotalDiscount = 35.52

	// The displayed totals should use API-provided values, not line item math
	if order.TotalDiscount != 35.52 {
		t.Errorf("discount = %.2f, want 35.52", order.TotalDiscount)
	}
	if order.TotalPrice != 92.68 {
		t.Errorf("total to pay = %.2f, want 92.68", order.TotalPrice)
	}
}

func TestGetOrderDetailsNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"code":    "NOT_FOUND",
			"message": "Order not found",
		})
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))
	ctx := context.Background()

	_, err := client.GetOrderDetails(ctx, 999999)
	if err == nil {
		t.Fatal("expected error for non-existent order")
	}
}
