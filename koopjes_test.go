package appie

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchStores(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req graphQLRequest
		json.NewDecoder(r.Body).Decode(&req)

		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"storesSearch": map[string]any{
					"result": []map[string]any{
						{
							"id":        1527,
							"name":      "Albert Heijn Utrecht Vondellaan",
							"storeType": "REGULAR",
							"address": map[string]any{
								"street":      "Vondellaan",
								"houseNumber": "200",
								"city":        "Utrecht",
								"postalCode":  "3521GZ",
							},
						},
						{
							"id":        1042,
							"name":      "Albert Heijn Utrecht Nachtegaalstraat",
							"storeType": "REGULAR",
							"address": map[string]any{
								"street":      "Nachtegaalstraat",
								"houseNumber": "32",
								"city":        "Utrecht",
								"postalCode":  "3581DE",
							},
						},
					},
				},
			},
		})
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))
	client.accessToken = "test"

	stores, err := client.SearchStores(context.Background(), "3521GZ")
	if err != nil {
		t.Fatal(err)
	}

	if len(stores) != 2 {
		t.Fatalf("expected 2 stores, got %d", len(stores))
	}

	if stores[0].ID != 1527 {
		t.Errorf("expected store ID 1527, got %d", stores[0].ID)
	}
	if stores[0].Address.City != "Utrecht" {
		t.Errorf("expected city Utrecht, got %s", stores[0].Address.City)
	}
}

func TestGetBargains(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req graphQLRequest
		json.NewDecoder(r.Body).Decode(&req)

		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"bargainItems": []map[string]any{
					{
						"product": map[string]any{
							"id":            12345,
							"title":         "AH Roodbaarsfilet",
							"brand":         "AH",
							"salesUnitSize": "ca. 200 g",
						},
						"categoryTitle": "Vis",
						"markdown": map[string]any{
							"markdownType":           "PERCENTAGE",
							"markdownExpirationDate": "2026-03-08",
							"markdownPercentage":     35.0,
						},
						"stock": 3,
						"bargainPrice": map[string]any{
							"priceWas": "5.99",
							"priceNow": "3.89",
						},
					},
				},
			},
		})
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))
	client.accessToken = "test"

	bargains, err := client.GetBargains(context.Background(), 1527)
	if err != nil {
		t.Fatal(err)
	}

	if len(bargains) != 1 {
		t.Fatalf("expected 1 bargain, got %d", len(bargains))
	}

	b := bargains[0]
	if b.Product.Title != "AH Roodbaarsfilet" {
		t.Errorf("expected title AH Roodbaarsfilet, got %s", b.Product.Title)
	}
	if b.MarkdownPercentage != 35.0 {
		t.Errorf("expected 35%% markdown, got %g%%", b.MarkdownPercentage)
	}
	if b.PriceWas != "5.99" {
		t.Errorf("expected priceWas 5.99, got %s", b.PriceWas)
	}
	if b.PriceNow != "3.89" {
		t.Errorf("expected priceNow 3.89, got %s", b.PriceNow)
	}
	if b.Stock != 3 {
		t.Errorf("expected stock 3, got %d", b.Stock)
	}
	if b.ExpirationDate != "2026-03-08" {
		t.Errorf("expected expiration 2026-03-08, got %s", b.ExpirationDate)
	}
}
