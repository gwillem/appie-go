package appie

import (
	"context"
	"os"
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

func TestNewClient(t *testing.T) {
	client := New()
	if client == nil {
		t.Fatal("expected client, got nil")
	}
	if client.baseURL != defaultBaseURL {
		t.Errorf("expected baseURL %s, got %s", defaultBaseURL, client.baseURL)
	}
}

func TestNewWithConfig(t *testing.T) {
	// Create a temp config file
	tmpFile, err := os.CreateTemp("", "appie-test-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(`{"access_token": "test-token", "refresh_token": "test-refresh"}`)
	if err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	client, err := NewWithConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if client.AccessToken() != "test-token" {
		t.Errorf("expected access token 'test-token', got '%s'", client.AccessToken())
	}
	if client.RefreshTokenValue() != "test-refresh" {
		t.Errorf("expected refresh token 'test-refresh', got '%s'", client.RefreshTokenValue())
	}
}

func TestNewWithConfigMissing(t *testing.T) {
	client, err := NewWithConfig("/nonexistent/path.json")
	if err != nil {
		t.Fatalf("unexpected error for missing config: %v", err)
	}
	if client == nil {
		t.Fatal("expected client even with missing config")
	}
}

func TestLoginURL(t *testing.T) {
	client := New()
	url := client.LoginURL()
	expected := "https://login.ah.nl/login?client_id=appie-ios&response_type=code&redirect_uri=appie://login-exit"
	if url != expected {
		t.Errorf("expected %s, got %s", expected, url)
	}
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

	err := client.RefreshToken(ctx)
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

	// Save updated tokens
	if client.configPath != "" {
		if err := client.SaveConfig(); err != nil {
			t.Logf("warning: failed to save config: %v", err)
		}
	}
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

func TestGetBonusProducts(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// Category is required - use "Fruit, verse sappen" or similar
	products, err := client.GetBonusProducts(ctx, "Vlees", 10)
	if err != nil {
		t.Fatalf("failed to get bonus products: %v", err)
	}

	t.Logf("Found %d bonus products in Vlees category", len(products))
	for i, p := range products {
		if i >= 5 {
			break
		}
		t.Logf("  - %s: €%.2f (was €%.2f)", p.Title, p.Price.Now, p.Price.Was)
	}
}

func TestGetOrder(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	order, err := client.GetOrder(ctx)
	if err != nil {
		t.Fatalf("failed to get order: %v", err)
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

func TestGetFeatureFlags(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	flags, err := client.GetFeatureFlags(ctx)
	if err != nil {
		t.Fatalf("failed to get feature flags: %v", err)
	}

	t.Logf("Feature flags: %d total", len(flags))

	// Check some known flags
	knownFlags := []string{"dark-mode", "ah-premium", "my-list", "push-notifications", "nutriscore"}
	for _, flag := range knownFlags {
		t.Logf("  %s: %d%% (enabled: %v)", flag, flags.Rollout(flag), flags.IsEnabled(flag))
	}
}
