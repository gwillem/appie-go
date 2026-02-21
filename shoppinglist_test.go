package appie

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetShoppingListItems(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/graphql" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		resp := graphQLResponse[json.RawMessage]{
			Data: json.RawMessage(`{
				"favoriteListV2": [{
					"id": "LIST-123",
					"description": "Boodschappen",
					"totalSize": 3,
					"items": [
						{"id": "item-1", "productId": 12345, "quantity": 2},
						{"id": "item-2", "productId": 67890, "quantity": 1},
						{"id": "item-3", "productId": 0, "quantity": 1}
					]
				}]
			}`),
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL))
	items, err := client.GetShoppingListItems(context.Background(), "list-123")
	if err != nil {
		t.Fatal(err)
	}

	if len(items) != 3 {
		t.Fatalf("got %d items, want 3", len(items))
	}

	if items[0].ID != "item-1" {
		t.Errorf("got ID %q, want %q", items[0].ID, "item-1")
	}
	if items[0].ProductID != 12345 {
		t.Errorf("got ProductID %d, want 12345", items[0].ProductID)
	}
	if items[0].Quantity != 2 {
		t.Errorf("got Quantity %d, want 2", items[0].Quantity)
	}
	if items[2].ProductID != 0 {
		t.Errorf("got ProductID %d for free-text item, want 0", items[2].ProductID)
	}
}

func TestGetShoppingLists(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mobile-services/lists/v3/lists" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		resp := []listResponse{
			{ID: "abc-123", Description: "Boodschappen", ItemCount: 5},
			{ID: "def-456", Description: "Weekmenu", ItemCount: 3},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL))
	lists, err := client.GetShoppingLists(context.Background(), 0)
	if err != nil {
		t.Fatal(err)
	}

	if len(lists) != 2 {
		t.Fatalf("got %d lists, want 2", len(lists))
	}

	if lists[0].Name != "Boodschappen" {
		t.Errorf("got Name %q, want %q", lists[0].Name, "Boodschappen")
	}
	if lists[0].ItemCount != 5 {
		t.Errorf("got ItemCount %d, want 5", lists[0].ItemCount)
	}
	if lists[1].ID != "def-456" {
		t.Errorf("got ID %q, want %q", lists[1].ID, "def-456")
	}
}
