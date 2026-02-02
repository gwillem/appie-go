package appie

import (
	"context"
	"testing"
)

func TestGetBonusProductsByCategories(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// Use single category to avoid rate limiting issues with rapid requests
	categories := []string{"Vlees"}
	products, err := client.GetBonusProductsByCategories(ctx, categories)
	if err != nil {
		t.Fatalf("failed to get bonus products by categories: %v", err)
	}

	if len(products) == 0 {
		t.Error("expected at least one product")
	}

	t.Logf("Found %d bonus products across %d categories", len(products), len(categories))
	for i, p := range products {
		if i >= 10 {
			break
		}
		t.Logf("  - %s: €%.2f (was €%.2f) [%s]", p.Title, p.Price.Now, p.Price.Was, p.Category)
	}
}

func TestGetBonusProductsByCategoriesEmpty(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	products, err := client.GetBonusProductsByCategories(ctx, nil)
	if err != nil {
		t.Fatalf("unexpected error for empty categories: %v", err)
	}

	if products == nil {
		t.Error("expected empty slice, got nil")
	}
}

func TestGetBonusProductsByCategoriesDeduplicates(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// Request same category twice - should deduplicate results
	categories := []string{"Vlees", "Vlees"}
	products, err := client.GetBonusProductsByCategories(ctx, categories)
	if err != nil {
		t.Fatalf("failed to get bonus products: %v", err)
	}

	// Check for duplicate IDs
	seen := make(map[int]bool)
	for _, p := range products {
		if seen[p.ID] {
			t.Errorf("duplicate product ID: %d (%s)", p.ID, p.Title)
		}
		seen[p.ID] = true
	}
}
