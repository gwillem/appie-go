package appie

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchProductsFiltered(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify it's the REST search endpoint
		if r.URL.Path != "/mobile-services/product/search/v2" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("query") != "kaas" {
			t.Errorf("expected query 'kaas', got %q", r.URL.Query().Get("query"))
		}

		json.NewEncoder(w).Encode(searchResponse{
			Products: []productResponse{
				{WebshopID: 123, Title: "AH Kaas", Brand: "AH", MainCategory: "Kaas", SalesUnitSize: "200 g",
					CurrentPrice: 2.99, PriceBeforeBonus: 3.49, IsBonus: true, BonusMechanism: "2 VOOR 5.50", IsOrderable: true},
				{WebshopID: 456, Title: "Old Amsterdam", Brand: "Old Amsterdam", MainCategory: "Kaas", SalesUnitSize: "300 g",
					CurrentPrice: 5.99, PriceBeforeBonus: 5.99, IsBonus: false, IsOrderable: true},
				{WebshopID: 789, Title: "AH Belegen", Brand: "AH", MainCategory: "Kaas", SalesUnitSize: "250 g",
					CurrentPrice: 4.49, PriceBeforeBonus: 4.99, IsBonus: true, BonusMechanism: "25% KORTING", IsOrderable: true},
			},
		})
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))
	ctx := context.Background()

	// With bonus filter: should only return bonus products
	products, err := client.SearchProductsFiltered(ctx, SearchOptions{
		Query: "kaas",
		Limit: 10,
		Bonus: true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(products) != 2 {
		t.Fatalf("expected 2 bonus products, got %d", len(products))
	}
	if products[0].ID != 123 {
		t.Errorf("expected product ID 123, got %d", products[0].ID)
	}
	if products[0].BonusMechanism != "2 VOOR 5.50" {
		t.Errorf("expected '2 VOOR 5.50', got %q", products[0].BonusMechanism)
	}
	if products[1].ID != 789 {
		t.Errorf("expected product ID 789, got %d", products[1].ID)
	}
	if products[1].BonusMechanism != "25% KORTING" {
		t.Errorf("expected '25%% KORTING', got %q", products[1].BonusMechanism)
	}
}

func TestSearchProductsFilteredNoBonus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Without bonus filter, size should be the requested limit (not over-fetched)
		if r.URL.Query().Get("size") != "5" {
			t.Errorf("expected size=5, got %q", r.URL.Query().Get("size"))
		}
		json.NewEncoder(w).Encode(searchResponse{
			Products: []productResponse{
				{WebshopID: 123, Title: "AH Kaas", CurrentPrice: 2.99, PriceBeforeBonus: 2.99},
			},
		})
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))
	ctx := context.Background()

	products, err := client.SearchProductsFiltered(ctx, SearchOptions{Query: "kaas", Limit: 5})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(products) != 1 {
		t.Fatalf("expected 1 product, got %d", len(products))
	}
}

func TestSearchProductsBonusPagination(t *testing.T) {
	// limit=3, so pageSize=3*5=15. We need totalElements > 15 to trigger pagination.
	var requestCount int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		page := r.URL.Query().Get("page")

		var products []productResponse
		switch page {
		case "0":
			// First page: only 1 bonus in 15 non-bonus products
			products = make([]productResponse, 15)
			for i := range products {
				products[i] = productResponse{WebshopID: i + 1, Title: fmt.Sprintf("P%d", i+1), CurrentPrice: float64(i + 1)}
			}
			products[3].IsBonus = true
			products[3].BonusMechanism = "X"
		case "1":
			// Second page: 2 bonus
			products = []productResponse{
				{WebshopID: 16, Title: "P16", CurrentPrice: 16, IsBonus: true, BonusMechanism: "Y"},
				{WebshopID: 17, Title: "P17", CurrentPrice: 17, IsBonus: false},
				{WebshopID: 18, Title: "P18", CurrentPrice: 18, IsBonus: true, BonusMechanism: "Z"},
			}
		}
		json.NewEncoder(w).Encode(searchResponse{
			Products: products,
			Page: struct {
				Number        int `json:"number"`
				Size          int `json:"size"`
				TotalElements int `json:"totalElements"`
				TotalPages    int `json:"totalPages"`
			}{Size: 15, TotalElements: 30, TotalPages: 2},
		})
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))
	ctx := context.Background()

	products, err := client.SearchProductsFiltered(ctx, SearchOptions{Query: "test", Limit: 3, Bonus: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(products) != 3 {
		t.Fatalf("expected 3 bonus products, got %d", len(products))
	}
	if requestCount != 2 {
		t.Errorf("expected 2 requests (pagination), got %d", requestCount)
	}
	if products[0].BonusMechanism != "X" || products[1].BonusMechanism != "Y" || products[2].BonusMechanism != "Z" {
		t.Errorf("unexpected mechanisms: %q, %q, %q", products[0].BonusMechanism, products[1].BonusMechanism, products[2].BonusMechanism)
	}
}

func TestGetProductUsesFetchProductDetails(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		var body struct {
			Query     string         `json:"query"`
			Variables map[string]any `json:"variables"`
		}
		json.NewDecoder(r.Body).Decode(&body)
		if body.Variables["productId"] != float64(51768) {
			t.Errorf("expected productId 51768, got %v", body.Variables["productId"])
		}
		fmt.Fprint(w, `{"data":{"product":{
			"id":51768,"title":"AH Boerenkool","brand":"AH","category":"Groente, aardappelen/Boerenkool",
			"summary":"Boerenkool voor stamppot.","icons":["NUTRISCORE_A","LOCALLY_PRODUCED"],"salesUnitSize":"300 g",
			"priceV2":{"now":{"amount":1.99},"was":{"amount":1.99},
				"unitInfo":{"description":"KG","price":{"formattedV2":"€6.63"}},
				"discount":{"description":"1 + 1 gratis"}},
			"availability":{"online":{"status":"AVAILABLE"},"isOrderable":true},
			"imagePack":[{"small":{"url":"s.webp","width":200,"height":200},"large":{"url":"l.webp","width":800,"height":800}}]}}}`)
	}))
	defer server.Close()

	product, err := New(WithBaseURL(server.URL)).GetProduct(context.Background(), 51768)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	checks := map[string][2]any{
		"ID":                   {product.ID, 51768},
		"Title":                {product.Title, "AH Boerenkool"},
		"Category":             {product.Category, "Groente, aardappelen"},
		"SubCategory":          {product.SubCategory, "Boerenkool"},
		"Price.Now":            {product.Price.Now, 1.99},
		"IsBonus":              {product.IsBonus, true},
		"BonusMechanism":       {product.BonusMechanism, "1 + 1 gratis"},
		"IsAvailable":          {product.IsAvailable, true},
		"IsOrderable":          {product.IsOrderable, true},
		"UnitSize":             {product.UnitSize, "300 g"},
		"UnitPriceDescription": {product.UnitPriceDescription, "per kg €6.63"},
		"NutriScore":           {product.NutriScore, "A"},
		"Images":               {len(product.Images), 2},
	}
	for name, c := range checks {
		if c[0] != c[1] {
			t.Errorf("%s: expected %v, got %v", name, c[1], c[0])
		}
	}
}

func TestGetProductNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"data":{"product":null}}`)
	}))
	defer server.Close()

	if _, err := New(WithBaseURL(server.URL)).GetProduct(context.Background(), 1); err == nil {
		t.Error("expected an error for a missing product")
	}
}
