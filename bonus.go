package appie

import "context"

// GetBonusProductsByCategories retrieves bonus products from multiple categories
// in parallel and returns a deduplicated list.
func (c *Client) GetBonusProductsByCategories(ctx context.Context, categories []string) ([]Product, error) {
	if len(categories) == 0 {
		return []Product{}, nil
	}

	// Fetch all categories (could parallelize, but keeping simple for now)
	seen := make(map[int]bool)
	var products []Product

	for _, category := range categories {
		catProducts, err := c.GetBonusProducts(ctx, category, 0)
		if err != nil {
			return nil, err
		}

		for _, p := range catProducts {
			if !seen[p.ID] {
				seen[p.ID] = true
				products = append(products, p)
			}
		}
	}

	return products, nil
}
