package appie

import (
	"context"
	"fmt"
)

const storesSearchQuery = `query StoresSearch($filter: StoresFilterInput) {
	storesSearch(filter: $filter, limit: 5) {
		result {
			id
			name
			storeType
			address { street houseNumber houseNumberExtra postalCode city }
		}
	}
}`

const bargainItemsQuery = `query BargainItems($storeId: String!) {
	bargainItems(storeId: $storeId) {
		product {
			id
			title
			brand
			salesUnitSize
		}
		categoryTitle
		markdown {
			markdownType
			markdownExpirationDate
			markdownPercentage
		}
		stock
		bargainPrice {
			priceWas
			priceNow
		}
	}
}`

// Store represents an Albert Heijn store location.
type Store struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	StoreType string       `json:"storeType"`
	Address   StoreAddress `json:"address"`
}

// StoreAddress represents a store's address.
type StoreAddress struct {
	Street           string `json:"street"`
	HouseNumber      string `json:"houseNumber"`
	HouseNumberExtra string `json:"houseNumberExtra,omitempty"`
	PostalCode       string `json:"postalCode"`
	City             string `json:"city"`
}

// Bargain represents a last-chance discounted product at a specific store.
type Bargain struct {
	Product            BargainProduct `json:"product"`
	Category           string         `json:"category"`
	MarkdownType       string         `json:"markdownType"`
	MarkdownPercentage float64        `json:"markdownPercentage"`
	ExpirationDate     string         `json:"expirationDate"`
	Stock              int            `json:"stock"`
	PriceWas           string         `json:"priceWas"`
	PriceNow           string         `json:"priceNow"`
}

// BargainProduct contains the basic product info for a bargain item.
type BargainProduct struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Brand    string `json:"brand"`
	UnitSize string `json:"unitSize"`
}

type storesSearchResponse struct {
	StoresSearch struct {
		Result []Store `json:"result"`
	} `json:"storesSearch"`
}

type bargainItemsResponse struct {
	BargainItems []struct {
		Product struct {
			ID            int    `json:"id"`
			Title         string `json:"title"`
			Brand         string `json:"brand"`
			SalesUnitSize string `json:"salesUnitSize"`
		} `json:"product"`
		CategoryTitle string `json:"categoryTitle"`
		Markdown      struct {
			MarkdownType           string  `json:"markdownType"`
			MarkdownExpirationDate string  `json:"markdownExpirationDate"`
			MarkdownPercentage     float64 `json:"markdownPercentage"`
		} `json:"markdown"`
		Stock        int `json:"stock"`
		BargainPrice struct {
			PriceWas string `json:"priceWas"`
			PriceNow string `json:"priceNow"`
		} `json:"bargainPrice"`
	} `json:"bargainItems"`
}

// SearchStores finds AH stores near a postal code.
func (c *Client) SearchStores(ctx context.Context, postalCode string) ([]Store, error) {
	vars := map[string]any{
		"filter": map[string]any{"postalCode": postalCode},
	}
	var resp storesSearchResponse
	if err := c.DoGraphQL(ctx, storesSearchQuery, vars, &resp); err != nil {
		return nil, fmt.Errorf("store search failed: %w", err)
	}
	return resp.StoresSearch.Result, nil
}

// GetBargains retrieves last-chance discounted products (laatste kans koopjes) for a store.
func (c *Client) GetBargains(ctx context.Context, storeID int) ([]Bargain, error) {
	vars := map[string]any{"storeId": fmt.Sprintf("%d", storeID)}
	var resp bargainItemsResponse
	if err := c.DoGraphQL(ctx, bargainItemsQuery, vars, &resp); err != nil {
		return nil, fmt.Errorf("get bargains failed: %w", err)
	}

	bargains := make([]Bargain, len(resp.BargainItems))
	for i, item := range resp.BargainItems {
		bargains[i] = Bargain{
			Product: BargainProduct{
				ID:       item.Product.ID,
				Title:    item.Product.Title,
				Brand:    item.Product.Brand,
				UnitSize: item.Product.SalesUnitSize,
			},
			Category:           item.CategoryTitle,
			MarkdownType:       item.Markdown.MarkdownType,
			MarkdownPercentage: item.Markdown.MarkdownPercentage,
			ExpirationDate:     item.Markdown.MarkdownExpirationDate,
			Stock:              item.Stock,
			PriceWas:           item.BargainPrice.PriceWas,
			PriceNow:           item.BargainPrice.PriceNow,
		}
	}
	return bargains, nil
}
