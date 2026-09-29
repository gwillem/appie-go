package appie

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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

func TestGetFulfillmentsByStatus(t *testing.T) {
	var gotStatus string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" {
			t.Errorf("unexpected path: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		var req graphQLRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("decode request: %v", err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if !strings.Contains(req.Query, "$status: FulfillmentStatus!") {
			t.Errorf("query does not declare status variable: %s", req.Query)
		}
		gotStatus, _ = req.Variables["status"].(string)

		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"orderFulfillments": map[string]any{
					"result": []map[string]any{
						{
							"orderId":              384664324,
							"statusCode":           5,
							"statusDescription":    "Delivered",
							"shoppingType":         "DELIVERY",
							"transactionCompleted": true,
							"modifiable":           false,
							"totalPrice": map[string]any{
								"totalPrice": map[string]any{"amount": 103.72},
							},
							"delivery": map[string]any{
								"status": "DELIVERED",
								"method": "HOME",
								"slot": map[string]any{
									"date":        "2026-06-16",
									"dateDisplay": "dinsdag 16 juni",
									"timeDisplay": "19:00 - 21:00",
									"startTime":   "19:00",
									"endTime":     "21:00",
								},
							},
						},
					},
				},
			},
		})
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))
	fulfillments, err := client.GetFulfillmentsByStatus(context.Background(), FulfillmentStatusClosed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotStatus != "CLOSED" {
		t.Fatalf("expected status CLOSED, got %q", gotStatus)
	}
	if len(fulfillments) != 1 {
		t.Fatalf("expected 1 fulfillment, got %d", len(fulfillments))
	}
	if fulfillments[0].OrderID != 384664324 {
		t.Errorf("expected order 384664324, got %d", fulfillments[0].OrderID)
	}
	if fulfillments[0].Status != "DELIVERED" {
		t.Errorf("expected status DELIVERED, got %q", fulfillments[0].Status)
	}
	if fulfillments[0].TotalPrice != 103.72 {
		t.Errorf("expected total 103.72, got %.2f", fulfillments[0].TotalPrice)
	}
}

func TestGetFulfillmentsByStatusRejectsInvalidStatus(t *testing.T) {
	client := New(WithTokens("test", "test"))
	_, err := client.GetFulfillmentsByStatus(context.Background(), FulfillmentStatus("INVALID"))
	if err == nil {
		t.Fatal("expected invalid status error")
	}
}

func TestGetOrderSubmissionInfo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("appie-current-order-id"); got != "716976811" {
			t.Fatalf("current order header = %q, want 716976811", got)
		}
		req, _ := readGraphQLRequest(t, r)
		if !strings.Contains(req.Query, "orderValueLimits") || !strings.Contains(req.Query, "checkoutValidateOrder") {
			t.Fatalf("query does not include checkout readiness fields: %s", req.Query)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"order": map[string]any{
					"id":                 716976811,
					"state":              "REOPENED",
					"submitted":          false,
					"lastUserChangeTime": "2026-06-29T05:49:57Z",
					"price": map[string]any{
						"priceTotalPayable": map[string]any{"amount": 6.57},
					},
				},
				"orderValueLimits": map[string]any{
					"minimumOrderValue": map[string]any{"amount": 50.0, "deadline": "2026-06-29T16:00:00Z"},
					"maximumOrderValue": map[string]any{"amount": 999999.0},
					"submittable":       true,
				},
				"checkoutValidateOrder": map[string]any{
					"errors": []map[string]any{
						{
							"__typename": "CheckoutErrorResponse",
							"code":       "ORDER_ATP_FAILED",
							"message":    "The order ATP check failed",
							"data": []map[string]any{
								{
									"__typename": "CheckoutErrorData",
									"errorType":  "UNKNOWN",
									"orderLines": []map[string]any{
										{
											"count": 1, "available": 0, "limitType": "STOCK_LIMIT",
											"product": map[string]any{"id": 578190, "title": "AH Borrelnoten Shanghai", "unitSize": "300 g"},
										},
									},
								},
							},
						},
					},
					"atpError": map[string]any{
						"__typename": "CheckoutATPError",
						"stockLimits": []map[string]any{
							{
								"count": 1, "available": 0, "limitType": "STOCK_LIMIT",
								"product": map[string]any{"id": 578190, "title": "AH Borrelnoten Shanghai", "unitSize": "300 g"},
							},
						},
						"orderLimits": []any{},
					},
				},
			},
		})
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))
	info, err := client.GetOrderSubmissionInfo(context.Background(), 716976811)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info.State != "REOPENED" {
		t.Fatalf("state = %q, want REOPENED", info.State)
	}
	if info.TotalPrice != 6.57 {
		t.Fatalf("total = %.2f, want 6.57", info.TotalPrice)
	}
	if !info.ValueLimits.Submittable {
		t.Fatal("expected submittable")
	}
	if info.ValueLimits.MinimumOrderValue.Amount != 50 {
		t.Fatalf("minimum = %.2f, want 50", info.ValueLimits.MinimumOrderValue.Amount)
	}
	if info.ValidationErrors != 1 || !info.HasATPError {
		t.Fatalf("validation summary = %d errors, ATP %t", info.ValidationErrors, info.HasATPError)
	}
	if got := info.CheckoutErrors[0].Code; got != "ORDER_ATP_FAILED" {
		t.Fatalf("checkout error code = %q", got)
	}
	if info.ATPError == nil || len(info.ATPError.StockLimits) != 1 {
		t.Fatalf("ATP error = %#v", info.ATPError)
	}
	line := info.ATPError.StockLimits[0]
	if line.Product == nil || line.Product.ID != 578190 || line.Available != 0 || line.Count != 1 {
		t.Fatalf("stock-limit line = %#v", line)
	}
}

func TestGetDefaultDCTCard(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		req, _ := readGraphQLRequest(t, r)
		if !strings.Contains(req.Query, "paymentsGetDCTCards") {
			t.Fatalf("unexpected query: %s", req.Query)
		}
		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"paymentsGetDCTCards": []map[string]any{
					{
						"cardId":    "card-secondary",
						"cardAlias": "secondary",
						"default":   false,
						"status":    "ACTIVE",
					},
					{
						"cardId":    "card-default",
						"cardAlias": "default",
						"default":   true,
						"status":    "ACTIVE",
					},
				},
			},
		})
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))
	card, err := client.GetDefaultDCTCard(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if card.CardID != "card-default" {
		t.Fatalf("card ID = %q, want card-default", card.CardID)
	}
}

func TestSubmitOrderUsesDCTPayloadV4(t *testing.T) {
	var sawSubmit bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		req, _ := readGraphQLRequest(t, r)
		switch {
		case strings.Contains(req.Query, "OrderSubmissionInfo"):
			json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{
					"order": map[string]any{
						"id":                 716976811,
						"state":              "REOPENED",
						"submitted":          false,
						"lastUserChangeTime": "2026-06-29T05:49:57Z",
						"price": map[string]any{
							"priceTotalPayable": map[string]any{"amount": 6.57},
						},
					},
					"orderValueLimits": map[string]any{
						"minimumOrderValue": map[string]any{"amount": 50.0},
						"maximumOrderValue": map[string]any{"amount": 999999.0},
						"submittable":       true,
					},
					"checkoutValidateOrder": map[string]any{"errors": []any{}, "atpError": nil},
				},
			})
		case strings.Contains(req.Query, "DCTCards"):
			json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{
					"paymentsGetDCTCards": []map[string]any{
						{
							"cardId":    "card-default",
							"cardAlias": "0 30",
							"default":   true,
							"status":    "ACTIVE",
						},
					},
				},
			})
		case strings.Contains(req.Query, "CheckoutConfirmOrder"):
			sawSubmit = true
			if !strings.Contains(req.Query, "CheckoutConfirmOrderPayloadV4") {
				t.Fatalf("submit query does not use V4 payload: %s", req.Query)
			}
			orderInfo, ok := req.Variables["orderInfo"].(map[string]any)
			if !ok {
				t.Fatalf("orderInfo variable missing or wrong type: %#v", req.Variables["orderInfo"])
			}
			if got := orderInfo["channel"]; got != "IOS" {
				t.Fatalf("channel = %v, want IOS", got)
			}
			if got := orderInfo["orderLastModified"]; got != "2026-06-29T05:49:57Z" {
				t.Fatalf("orderLastModified = %v", got)
			}
			if got := orderInfo["paymentMethod"]; got != "DCT" {
				t.Fatalf("paymentMethod = %v, want DCT", got)
			}
			dct, ok := orderInfo["dct"].(map[string]any)
			if !ok {
				t.Fatalf("dct variable missing or wrong type: %#v", orderInfo["dct"])
			}
			if got := dct["cardId"]; got != "card-default" {
				t.Fatalf("cardId = %v, want card-default", got)
			}
			json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{
					"checkoutConfirmOrderV4": map[string]any{
						"status":       "SUCCESS",
						"errorMessage": "",
						"errors":       []any{},
						"atpError":     nil,
						"data": map[string]any{
							"order": map[string]any{
								"id":        716976811,
								"state":     "SUBMITTED",
								"submitted": true,
							},
							"payments": []map[string]any{
								{"mutation": map[string]any{"status": "SUCCESS"}},
							},
						},
					},
				},
			})
		default:
			t.Fatalf("unexpected query: %s", req.Query)
		}
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))
	result, err := client.SubmitOrder(context.Background(), 716976811, OrderSubmitOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !sawSubmit {
		t.Fatal("submit mutation was not called")
	}
	if result.OrderState != "SUBMITTED" {
		t.Fatalf("state = %q, want SUBMITTED", result.OrderState)
	}
	if len(result.PaymentStatuses) != 1 || result.PaymentStatuses[0] != "SUCCESS" {
		t.Fatalf("payment statuses = %#v", result.PaymentStatuses)
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

func TestAddToOrderBySearch(t *testing.T) {
	var addedItems []struct {
		ProductID int `json:"productId"`
		Quantity  int `json:"quantity"`
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/mobile-services/product/search/v2" && r.Method == http.MethodGet:
			json.NewEncoder(w).Encode(map[string]any{
				"products": []map[string]any{
					{
						"webshopId":        12345,
						"title":            "AH Halfvolle melk",
						"brand":            "AH",
						"salesUnitSize":    "1 L",
						"currentPrice":     1.15,
						"priceBeforeBonus": 1.15,
						"availableOnline":  true,
						"isOrderable":      true,
					},
				},
				"page": map[string]any{
					"totalElements": 1,
				},
			})

		case r.URL.Path == "/mobile-services/order/v1/items" && r.Method == http.MethodPut:
			var body struct {
				Items []struct {
					ProductID int `json:"productId"`
					Quantity  int `json:"quantity"`
				} `json:"items"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("decode body: %v", err)
				http.Error(w, "bad request", 400)
				return
			}
			addedItems = body.Items
			w.WriteHeader(http.StatusOK)

		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))
	ctx := context.Background()

	// Search → should find 1 product
	products, err := client.SearchProducts(ctx, "halfvolle melk", 5)
	if err != nil {
		t.Fatalf("search: %v", err)
	}
	if len(products) != 1 {
		t.Fatalf("expected 1 product, got %d", len(products))
	}
	if products[0].ID != 12345 {
		t.Errorf("expected product ID 12345, got %d", products[0].ID)
	}

	// Add to order
	err = client.AddToOrder(ctx, []OrderItem{{ProductID: products[0].ID, Quantity: 3}})
	if err != nil {
		t.Fatalf("add to order: %v", err)
	}

	if len(addedItems) != 1 {
		t.Fatalf("expected 1 item added, got %d", len(addedItems))
	}
	if addedItems[0].ProductID != 12345 {
		t.Errorf("expected product ID 12345, got %d", addedItems[0].ProductID)
	}
	if addedItems[0].Quantity != 3 {
		t.Errorf("expected quantity 3, got %d", addedItems[0].Quantity)
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
