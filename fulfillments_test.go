package appie

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFulfillments(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		var req graphQLRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatal(err)
		}

		// Verify status variable is passed through
		status, ok := req.Variables["status"].(string)
		if !ok {
			t.Fatal("expected status variable")
		}

		resp := graphQLResponse[json.RawMessage]{
			Data: json.RawMessage(`{
				"orderFulfillments": {
					"result": [
						{
							"orderId": 387946303,
							"reopenable": false,
							"isSubscriptionOrder": false,
							"totalPrice": {
								"totalPrice": {"amount": 78.54}
							},
							"delivery": {
								"deliveryMessage": "Bezorgd op dinsdag 15 april",
								"method": "HOME",
								"homeShopCenterId": 1234,
								"status": "DELIVERED",
								"slot": {"date": "2025-04-15", "startTime": "18:00", "endTime": "20:00"}
							}
						},
						{
							"orderId": 387000100,
							"reopenable": true,
							"isSubscriptionOrder": true,
							"totalPrice": {
								"totalPrice": {"amount": 45.20}
							},
							"delivery": {
								"deliveryMessage": "Bezorgd op maandag 7 april",
								"method": "PICKUP",
								"homeShopCenterId": 5678,
								"status": "DELIVERED",
								"slot": {"date": "2025-04-07", "startTime": "", "endTime": ""}
							}
						}
					]
				}
			}`),
		}
		_ = status
		json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))

	t.Run("closed", func(t *testing.T) {
		fulfillments, err := client.GetFulfillments(context.Background(), FulfillmentClosed, 25)
		if err != nil {
			t.Fatal(err)
		}
		if len(fulfillments) != 2 {
			t.Fatalf("got %d fulfillments, want 2", len(fulfillments))
		}

		f := fulfillments[0]
		if f.OrderID != 387946303 {
			t.Errorf("got OrderID %d, want 387946303", f.OrderID)
		}
		if f.TotalPrice != 78.54 {
			t.Errorf("got TotalPrice %.2f, want 78.54", f.TotalPrice)
		}
		if f.Delivery.Slot.Date != "2025-04-15" {
			t.Errorf("got Date %q, want 2025-04-15", f.Delivery.Slot.Date)
		}
		if f.Delivery.Method != "HOME" {
			t.Errorf("got Method %q, want HOME", f.Delivery.Method)
		}
		if f.Delivery.Slot.StartTime != "18:00" {
			t.Errorf("got StartTime %q, want 18:00", f.Delivery.Slot.StartTime)
		}
		if f.DeliveryMessage != "Bezorgd op dinsdag 15 april" {
			t.Errorf("got DeliveryMessage %q", f.DeliveryMessage)
		}

		f2 := fulfillments[1]
		if !f2.IsSubscriptionOrder {
			t.Error("expected IsSubscriptionOrder=true")
		}
		if !f2.Reopenable {
			t.Error("expected Reopenable=true")
		}
	})

	t.Run("open", func(t *testing.T) {
		fulfillments, err := client.GetFulfillments(context.Background(), FulfillmentOpen, 0)
		if err != nil {
			t.Fatal(err)
		}
		if len(fulfillments) != 2 {
			t.Fatalf("got %d fulfillments, want 2", len(fulfillments))
		}
	})
}

func TestGetFulfillmentDetail(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		var req graphQLRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatal(err)
		}

		orderID, ok := req.Variables["orderId"].(float64)
		if !ok || int(orderID) != 387946303 {
			t.Errorf("expected orderId 387946303, got %v", req.Variables["orderId"])
		}

		resp := graphQLResponse[json.RawMessage]{
			Data: json.RawMessage(`{
				"orderFulfillment": {
					"orderId": 387946303,
					"closingDateTime": "2025-04-14T22:59:00Z",
					"cancellable": false,
					"reopenable": false,
					"delivery": {
						"slot": {
							"date": "2025-04-15",
							"startTime": "18:00",
							"endTime": "20:00"
						},
						"homeShopCenterId": 1234,
						"status": "DELIVERED",
						"method": "HOME",
						"address": {
							"postalCode": "1234AB",
							"city": "Amsterdam",
							"countryCode": "NL",
							"houseNumber": 42,
							"houseNumberExtra": "B",
							"street": "Keizersgracht"
						}
					},
					"costOverview": {
						"invoiceId": "3879463-00123"
					},
					"orderLines": [
						{
							"quantity": 2,
							"allocatedQuantity": 2,
							"product": {
								"id": 164358,
								"title": "AH Oranje zoete aardappel",
								"brand": "AH",
								"hqId": 164358,
								"priceV2": {
									"now": {"amount": 3.79},
									"was": null
								},
								"listPrice": {"amount": 3.79},
								"imagePack": [
									{
										"medium": {
											"url": "https://static.ah.nl/dam/product/164358_medium.jpg",
											"width": 200,
											"height": 200
										}
									}
								],
								"category": "Groente, aardappelen",
								"salesUnitSize": "1 kg"
							}
						},
						{
							"quantity": 1,
							"allocatedQuantity": 1,
							"product": {
								"id": 371880,
								"title": "Optimel Drinkyoghurt aardbei",
								"brand": "Optimel",
								"hqId": 371880,
								"priceV2": {
									"now": {"amount": 1.19},
									"was": {"amount": 1.59}
								},
								"listPrice": {"amount": 1.59},
								"imagePack": null,
								"category": "Zuivel, eieren",
								"salesUnitSize": "1 L"
							}
						},
						{
							"quantity": 3,
							"allocatedQuantity": 2,
							"product": null
						}
					]
				}
			}`),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))
	detail, err := client.GetFulfillmentDetail(context.Background(), 387946303)
	if err != nil {
		t.Fatal(err)
	}

	if detail.OrderID != 387946303 {
		t.Errorf("got OrderID %d, want 387946303", detail.OrderID)
	}
	if detail.ClosingDateTime != "2025-04-14T22:59:00Z" {
		t.Errorf("got ClosingDateTime %q", detail.ClosingDateTime)
	}
	if detail.InvoiceID != "3879463-00123" {
		t.Errorf("got InvoiceID %q, want 3879463-00123", detail.InvoiceID)
	}
	if detail.Delivery.Slot.StartTime != "18:00" {
		t.Errorf("got StartTime %q, want 18:00", detail.Delivery.Slot.StartTime)
	}
	if detail.Delivery.Address.City != "Amsterdam" {
		t.Errorf("got City %q, want Amsterdam", detail.Delivery.Address.City)
	}
	if detail.Delivery.Address.HouseNumberExtra != "B" {
		t.Errorf("got HouseNumberExtra %q, want B", detail.Delivery.Address.HouseNumberExtra)
	}

	if len(detail.OrderLines) != 3 {
		t.Fatalf("got %d order lines, want 3", len(detail.OrderLines))
	}

	ol := detail.OrderLines[0]
	if ol.Quantity != 2 {
		t.Errorf("line 0: got Quantity %d, want 2", ol.Quantity)
	}
	if ol.Product == nil {
		t.Fatal("line 0: expected product")
	}
	if ol.Product.ID != 164358 {
		t.Errorf("line 0: got product ID %d, want 164358", ol.Product.ID)
	}
	if ol.Product.Title != "AH Oranje zoete aardappel" {
		t.Errorf("line 0: got Title %q", ol.Product.Title)
	}
	if ol.Product.CurrentPrice != 3.79 {
		t.Errorf("line 0: got CurrentPrice %.2f, want 3.79", ol.Product.CurrentPrice)
	}
	if ol.Product.ImageURL != "https://static.ah.nl/dam/product/164358_medium.jpg" {
		t.Errorf("line 0: got ImageURL %q", ol.Product.ImageURL)
	}

	// Discounted product
	ol2 := detail.OrderLines[1]
	if ol2.Product.CurrentPrice != 1.19 {
		t.Errorf("line 1: got CurrentPrice %.2f, want 1.19", ol2.Product.CurrentPrice)
	}
	if ol2.Product.WasPrice != 1.59 {
		t.Errorf("line 1: got WasPrice %.2f, want 1.59", ol2.Product.WasPrice)
	}

	// Nil product (unavailable)
	ol3 := detail.OrderLines[2]
	if ol3.Product != nil {
		t.Error("line 2: expected nil product")
	}
	if ol3.Quantity != 3 {
		t.Errorf("line 2: got Quantity %d, want 3", ol3.Quantity)
	}
	if ol3.AllocatedQuantity != 2 {
		t.Errorf("line 2: got AllocatedQuantity %d, want 2", ol3.AllocatedQuantity)
	}
}
