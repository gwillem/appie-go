//go:build integration

package appie

import (
	"context"
	"testing"
)

func testClient(t *testing.T) *Client {
	t.Helper()

	// Try to load from .appie.json
	client, err := NewWithConfig(".appie.json")
	if err != nil {
		t.Skipf("failed to load config: %v", err)
	}

	if !client.IsAuthenticated() {
		t.Skip("no authentication tokens available")
	}

	return client
}

func TestGetAnonymousToken(t *testing.T) {
	client := New()
	ctx := context.Background()

	err := client.GetAnonymousToken(ctx)
	if err != nil {
		t.Fatalf("failed to get anonymous token: %v", err)
	}

	if client.AccessToken() == "" {
		t.Error("expected access token to be set")
	}
}

func TestRefreshToken(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	oldToken := client.AccessToken()

	err := client.refreshAccessToken(ctx)
	if err != nil {
		t.Fatalf("failed to refresh token: %v", err)
	}

	if client.AccessToken() == "" {
		t.Error("expected new access token")
	}

	// Token should have changed
	if client.AccessToken() == oldToken {
		t.Log("warning: access token did not change after refresh")
	}

	// Tokens are auto-saved when configPath is set
}

func TestSearchProducts(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	products, err := client.SearchProducts(ctx, "melk", 5)
	if err != nil {
		t.Fatalf("failed to search products: %v", err)
	}

	if len(products) == 0 {
		t.Error("expected at least one product")
	}

	for _, p := range products {
		t.Logf("Product: %s (ID: %d) - €%.2f", p.Title, p.ID, p.Price.Now)
	}
}

func TestGetProduct(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// Use a known product ID (AH Biologisch Rundergehakt)
	product, err := client.GetProduct(ctx, 436752)
	if err != nil {
		t.Fatalf("failed to get product: %v", err)
	}

	t.Logf("Product: %s (ID: %d) - €%.2f", product.Title, product.ID, product.Price.Now)
	t.Logf("  Brand: %s, Category: %s", product.Brand, product.Category)
	t.Logf("  NutriScore: %s, IsBonus: %v", product.NutriScore, product.IsBonus)
}

func TestGetProductFull(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	product, err := client.GetProductFull(ctx, 436752)
	if err != nil {
		t.Fatalf("failed to get full product: %v", err)
	}

	t.Logf("Product: %s (ID: %d)", product.Title, product.ID)
	t.Logf("  Nutritional info: %d entries", len(product.NutritionalInfo))
	for _, n := range product.NutritionalInfo {
		t.Logf("    %s (%s): %s", n.Name, n.Type, n.Value)
	}
}

func TestGetOrder(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	order, err := client.GetOrder(ctx)
	if err != nil {
		t.Skipf("skipping, no active order: %v", err)
	}

	t.Logf("Order ID: %s, State: %s, Items: %d, Total: €%.2f", order.ID, order.State, len(order.Items), order.TotalPrice)
	for _, item := range order.Items {
		t.Logf("  - %s (qty: %d)", item.Product.Title, item.Quantity)
	}
}

func TestGetShoppingList(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	lists, err := client.GetShoppingLists(ctx, 0)
	if err != nil {
		t.Fatalf("failed to get shopping lists: %v", err)
	}

	if len(lists) == 0 {
		t.Log("No shopping lists found")
		return
	}

	for _, list := range lists {
		t.Logf("Shopping list: %s (ID: %s, Items: %d)", list.Name, list.ID, list.ItemCount)
	}
}

func TestGetMember(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	member, err := client.GetMember(ctx)
	if err != nil {
		t.Fatalf("failed to get member: %v", err)
	}

	t.Logf("Member: %s %s (%s)", member.FirstName, member.LastName, member.Email)
}

func TestGetBonusCard(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	card, err := client.GetBonusCard(ctx)
	if err != nil {
		t.Fatalf("failed to get bonus card: %v", err)
	}

	t.Logf("Bonus card: %s (active: %v)", card.CardNumber, card.IsActive)
}

