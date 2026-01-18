package appie

import (
	"context"
	"fmt"
	"net/http"
)

// listResponse matches the API response for shopping lists.
type listResponse struct {
	ID                 string `json:"id"`
	Description        string `json:"description"`
	ItemCount          int    `json:"itemCount"`
	HasFavoriteProduct bool   `json:"hasFavoriteProduct"`
	ProductImages      [][]struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		URL    string `json:"url"`
	} `json:"productImages"`
}

// GetShoppingLists retrieves all shopping lists for the authenticated user.
// The API quirk requires a productId parameter, but returns all lists regardless.
// Pass 0 to use a default product ID (recommended).
func (c *Client) GetShoppingLists(ctx context.Context, productID int) ([]ShoppingList, error) {
	if productID <= 0 {
		productID = 1 // Default product ID - API requires it but returns all lists
	}
	path := fmt.Sprintf("/mobile-services/lists/v3/lists?productId=%d", productID)

	var result []listResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get shopping lists failed: %w", err)
	}

	lists := make([]ShoppingList, 0, len(result))
	for _, r := range result {
		lists = append(lists, ShoppingList{
			ID:        r.ID,
			Name:      r.Description,
			ItemCount: r.ItemCount,
		})
	}

	return lists, nil
}

// GetShoppingList retrieves the first (default) shopping list.
// Use GetShoppingLists if you need to access multiple lists.
func (c *Client) GetShoppingList(ctx context.Context) (*ShoppingList, error) {
	lists, err := c.GetShoppingLists(ctx, 0)
	if err != nil {
		return nil, err
	}

	if len(lists) == 0 {
		return nil, fmt.Errorf("no shopping lists found")
	}

	return &lists[0], nil
}

// AddToShoppingList adds items to the default shopping list.
// Items can be products (with ProductID) or free-text entries (with Name).
func (c *Client) AddToShoppingList(ctx context.Context, items []ListItem) error {
	body := map[string]any{
		"items": items,
	}

	if err := c.doRequest(ctx, http.MethodPost, "/mobile-services/lists/v3/lists/items", body, nil); err != nil {
		return fmt.Errorf("add to shopping list failed: %w", err)
	}

	return nil
}

// AddProductToShoppingList adds a product to the default shopping list.
// This is a convenience wrapper around AddToShoppingList.
func (c *Client) AddProductToShoppingList(ctx context.Context, productID int, quantity int) error {
	if quantity <= 0 {
		quantity = 1
	}

	items := []ListItem{{
		ProductID: productID,
		Quantity:  quantity,
	}}

	return c.AddToShoppingList(ctx, items)
}

// AddFreeTextToShoppingList adds a free-text item (not linked to a product) to the list.
// Useful for items like "bread" or "eggs" without specifying a specific product.
func (c *Client) AddFreeTextToShoppingList(ctx context.Context, name string, quantity int) error {
	if quantity <= 0 {
		quantity = 1
	}

	items := []ListItem{{
		Name:     name,
		Quantity: quantity,
	}}

	return c.AddToShoppingList(ctx, items)
}

// RemoveFromShoppingList removes an item from the shopping list.
func (c *Client) RemoveFromShoppingList(ctx context.Context, itemID string) error {
	path := fmt.Sprintf("/mobile-services/lists/v3/lists/items/%s", itemID)
	if err := c.doRequest(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("remove from shopping list failed: %w", err)
	}

	return nil
}

// CheckShoppingListItem marks an item as checked (picked) or unchecked.
// Checked items are typically displayed differently in the app UI.
func (c *Client) CheckShoppingListItem(ctx context.Context, itemID string, checked bool) error {
	body := map[string]any{
		"checked": checked,
	}

	path := fmt.Sprintf("/mobile-services/lists/v3/lists/items/%s", itemID)
	if err := c.doRequest(ctx, http.MethodPatch, path, body, nil); err != nil {
		return fmt.Errorf("check shopping list item failed: %w", err)
	}

	return nil
}

// ClearShoppingList removes all items from the shopping list.
func (c *Client) ClearShoppingList(ctx context.Context) error {
	list, err := c.GetShoppingList(ctx)
	if err != nil {
		return err
	}

	for _, item := range list.Items {
		if err := c.RemoveFromShoppingList(ctx, item.ID); err != nil {
			return fmt.Errorf("failed to remove item %s: %w", item.ID, err)
		}
	}

	return nil
}

// ShoppingListToOrder adds all unchecked product items from the shopping list to the order.
// Free-text items (without ProductID) are skipped. This is useful for quickly
// converting your shopping list into an order.
func (c *Client) ShoppingListToOrder(ctx context.Context) error {
	list, err := c.GetShoppingList(ctx)
	if err != nil {
		return err
	}

	var orderItems []OrderItem
	for _, item := range list.Items {
		if !item.Checked && item.ProductID > 0 {
			orderItems = append(orderItems, OrderItem{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			})
		}
	}

	if len(orderItems) == 0 {
		return nil
	}

	return c.AddToOrder(ctx, orderItems)
}
