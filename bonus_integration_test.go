//go:build integration

package appie

import (
	"context"
	"testing"
)

func TestGetBonusProducts(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	products, err := client.GetBonusProducts(ctx)
	if err != nil {
		t.Fatalf("failed to get bonus products: %v", err)
	}

	if len(products) == 0 {
		t.Fatal("expected at least one bonus product")
	}

	t.Logf("Found %d bonus products across all categories", len(products))
	for i, p := range products {
		t.Logf("  [%d] %s (ID: %d)", i, p.Title, p.ID)
		t.Logf("       €%.2f (was €%.2f), Mechanism: %q, Category: %s", p.Price.Now, p.Price.Was, p.BonusMechanism, p.Category)
	}
}
