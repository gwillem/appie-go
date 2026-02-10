//go:build integration

package appie

import (
	"context"
	"slices"
	"testing"
	"time"
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

	if client.accessToken == "" {
		t.Error("expected access token to be set")
	}
}

func TestRefreshToken(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	oldToken := client.accessToken

	err := client.refreshAccessToken(ctx)
	if err != nil {
		t.Fatalf("failed to refresh token: %v", err)
	}

	if client.accessToken == "" {
		t.Error("expected new access token")
	}

	// Token should have changed
	if client.accessToken == oldToken {
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

func TestOrderRoundTrip(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// 1. List all future orders with €0 amount
	fulfillments, err := client.GetFulfillments(ctx)
	if err != nil {
		t.Fatalf("failed to get fulfillments: %v", err)
	}

	empty := slices.DeleteFunc(slices.Clone(fulfillments), func(f Fulfillment) bool {
		return f.TotalPrice != 0
	})
	if len(empty) == 0 {
		t.Skip("no empty (€0) fulfillments found")
	}

	for _, f := range empty {
		t.Logf("Empty order %d: %s %s (€%.2f)", f.OrderID, f.Delivery.Slot.Date, f.Delivery.Slot.TimeDisplay, f.TotalPrice)
	}

	// 2. Pick the one most distant in the future (latest date)
	slices.SortFunc(empty, func(a, b Fulfillment) int {
		if a.Delivery.Slot.Date != b.Delivery.Slot.Date {
			if a.Delivery.Slot.Date < b.Delivery.Slot.Date {
				return 1
			}
			return -1
		}
		if a.Delivery.Slot.StartTime < b.Delivery.Slot.StartTime {
			return 1
		}
		return -1
	})
	target := empty[0]
	t.Logf("Selected order %d (%s %s)", target.OrderID, target.Delivery.Slot.Date, target.Delivery.Slot.TimeDisplay)

	// Reopen and select the order
	if err := client.ReopenOrder(ctx, target.OrderID); err != nil {
		t.Fatalf("failed to reopen order: %v", err)
	}
	client.SetOrderID(target.OrderID)

	// 3. Search for "pindakaas" and add top result
	pindakaas, err := client.SearchProducts(ctx, "pindakaas", 1)
	if err != nil {
		t.Fatalf("failed to search pindakaas: %v", err)
	}
	if len(pindakaas) == 0 {
		t.Fatal("no results for pindakaas")
	}
	t.Logf("Adding: %s (ID: %d)", pindakaas[0].Title, pindakaas[0].ID)

	// 4. Search for "wc papier" and add top result
	wcpapier, err := client.SearchProducts(ctx, "wc papier", 1)
	if err != nil {
		t.Fatalf("failed to search wc papier: %v", err)
	}
	if len(wcpapier) == 0 {
		t.Fatal("no results for wc papier")
	}
	t.Logf("Adding: %s (ID: %d)", wcpapier[0].Title, wcpapier[0].ID)

	// 5. Get bonus products (spotlight, falling back to regular) and add first 3
	bonus, err := client.GetSpotlightBonusProducts(ctx)
	if err != nil {
		t.Fatalf("failed to get spotlight bonus products: %v", err)
	}
	if len(bonus) < 3 {
		t.Logf("Spotlight has %d products, falling back to regular bonus", len(bonus))
		bonus, err = client.GetBonusProducts(ctx)
		if err != nil {
			t.Fatalf("failed to get bonus products: %v", err)
		}
	}

	orderItems := []OrderItem{
		{ProductID: pindakaas[0].ID, Quantity: 1},
		{ProductID: wcpapier[0].ID, Quantity: 1},
	}
	if len(bonus) >= 3 {
		for _, p := range bonus[:3] {
			t.Logf("Adding bonus: %s (ID: %d) - €%.2f", p.Title, p.ID, p.Price.Now)
			orderItems = append(orderItems, OrderItem{ProductID: p.ID, Quantity: 1})
		}
	} else {
		// Bonus promos are empty because the server scopes them to the active
		// order's delivery date, which is too far in the future for any current
		// bonus period.
		t.Logf("Only %d bonus products available (order date too far in the future), skipping bonus add", len(bonus))
	}

	if err := client.AddToOrder(ctx, orderItems); err != nil {
		t.Fatalf("failed to add to order: %v", err)
	}

	// 6. Display order summary
	summary, err := client.GetOrderSummary(ctx)
	if err != nil {
		t.Fatalf("failed to get order summary: %v", err)
	}
	t.Logf("Summary: %d items, €%.2f total, €%.2f discount", summary.TotalItems, summary.TotalPrice, summary.TotalDiscount)

	// 7. List order and ensure non-empty
	order, err := client.GetOrder(ctx)
	if err != nil {
		t.Fatalf("failed to get order: %v", err)
	}
	if len(order.Items) == 0 {
		t.Fatal("expected order to have items after adding products")
	}
	for _, item := range order.Items {
		t.Logf("  - %s (qty: %d)", item.Product.Title, item.Quantity)
	}

	// 8. Clear order
	if err := client.ClearOrder(ctx); err != nil {
		t.Fatalf("failed to clear order: %v", err)
	}

	// 9. Verify order is empty again
	fulfillments, err = client.GetFulfillments(ctx)
	if err != nil {
		t.Fatalf("failed to get fulfillments after clear: %v", err)
	}

	for _, f := range fulfillments {
		if f.OrderID == target.OrderID && f.TotalPrice != 0 {
			t.Errorf("expected order %d to be empty after clear, got €%.2f", target.OrderID, f.TotalPrice)
		}
	}
	t.Log("Order cleared successfully")

	// 10. Revert order back to submitted state so it's no longer the active order.
	// Without this, the account stays "in" the future order, which affects bonus
	// promo visibility on the server side.
	if err := client.RevertOrder(ctx, target.OrderID); err != nil {
		t.Fatalf("failed to revert order: %v", err)
	}
	t.Log("Order reverted to submitted state")

	// 11. Verify the bonus period contains today, confirming we're back to
	// the current context and not stuck in a future order's bonus period.
	startDate, endDate, err := client.getBonusPeriod(ctx)
	if err != nil {
		t.Fatalf("failed to get bonus period: %v", err)
	}
	today := time.Now().Format("2006-01-02")
	if today < startDate || today > endDate {
		t.Errorf("today %s is outside bonus period %s..%s; account may be stuck in future order context", today, startDate, endDate)
	}
	t.Logf("Bonus period %s..%s contains today (%s)", startDate, endDate, today)
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

func TestGetBonusGroupProducts(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// First get bonus products and find a group (ID==0, BonusSegmentID!="")
	bonus, err := client.GetBonusProducts(ctx)
	if err != nil {
		t.Fatalf("failed to get bonus products: %v", err)
	}

	var segmentID string
	for _, p := range bonus {
		if p.ID == 0 && p.BonusSegmentID != "" {
			segmentID = p.BonusSegmentID
			t.Logf("Found bonus group: %s (segment %s)", p.Title, p.BonusSegmentID)
			break
		}
	}
	if segmentID == "" {
		t.Skip("no bonus groups found in current period")
	}

	// Resolve the group to individual products
	products, err := client.GetBonusGroupProducts(ctx, segmentID)
	if err != nil {
		t.Fatalf("failed to get bonus group products: %v", err)
	}

	if len(products) == 0 {
		t.Fatal("expected at least one product in bonus group")
	}

	t.Logf("Bonus group %s: %d products", segmentID, len(products))
	for i, p := range products {
		if i >= 5 {
			t.Logf("  ... and %d more", len(products)-5)
			break
		}
		t.Logf("  - %s (ID: %d) €%.2f %s", p.Title, p.ID, p.Price.Now, p.BonusMechanism)
	}

	// All resolved products should have real IDs
	for _, p := range products {
		if p.ID == 0 {
			t.Errorf("expected non-zero ID for resolved product %q", p.Title)
		}
	}
}
