package appie

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// listResponse matches the API response for favorite lists (v3).
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

// shoppingListItem is the v2 request body format for adding items.
type shoppingListItem struct {
	Description   string `json:"description"`
	ProductID     int    `json:"productId,omitempty"`
	Quantity      int    `json:"quantity"`
	Type          string `json:"type"`
	OriginCode    string `json:"originCode"`
	SearchTerm    string `json:"searchTerm,omitempty"`
	StrikeThrough bool   `json:"strikeThrough"`
}

// GetShoppingLists retrieves all favorite lists (v3) for the authenticated user.
// The API quirk requires a productId parameter, but returns all lists regardless.
// Pass 0 to use a default product ID (recommended).
func (c *Client) GetShoppingLists(ctx context.Context, productID int) ([]ShoppingList, error) {
	if productID <= 0 {
		productID = 1 // Default product ID - API requires it but returns all lists
	}
	path := fmt.Sprintf("/mobile-services/lists/v3/lists?productId=%d", productID)

	var result []listResponse
	if err := c.DoRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
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

// GetShoppingListItems retrieves all items for a specific list (v2 GraphQL).
func (c *Client) GetShoppingListItems(ctx context.Context, listID string) ([]ListItem, error) {
	const query = `query FavoriteListV2($ids: [String!]!) {
  favoriteListV2(ids: $ids) {
    id
    description
    totalSize
    items {
      id
      productId
      quantity
    }
  }
}`

	variables := map[string]any{
		"ids": []string{strings.ToUpper(listID)},
	}

	type listItemResult struct {
		ID        string `json:"id"`
		ProductID int    `json:"productId"`
		Quantity  int    `json:"quantity"`
	}
	type listResult struct {
		ID          string           `json:"id"`
		Description string           `json:"description"`
		TotalSize   int              `json:"totalSize"`
		Items       []listItemResult `json:"items"`
	}

	// API may return a single object or array; try array first.
	var arrResult struct {
		FavoriteListV2 []listResult `json:"favoriteListV2"`
	}
	if err := c.DoGraphQL(ctx, query, variables, &arrResult); err != nil {
		return nil, fmt.Errorf("get shopping list items failed: %w", err)
	}

	if len(arrResult.FavoriteListV2) == 0 {
		return nil, fmt.Errorf("list %s not found", listID)
	}

	raw := arrResult.FavoriteListV2[0].Items
	items := make([]ListItem, 0, len(raw))
	for _, item := range raw {
		items = append(items, ListItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  max(item.Quantity, 1),
		})
	}

	return items, nil
}

// getShoppingList retrieves the first (default) favorite list with its items populated.
func (c *Client) getShoppingList(ctx context.Context) (*ShoppingList, error) {
	lists, err := c.GetShoppingLists(ctx, 0)
	if err != nil {
		return nil, err
	}

	if len(lists) == 0 {
		return nil, fmt.Errorf("no shopping lists found")
	}

	items, err := c.GetShoppingListItems(ctx, lists[0].ID)
	if err != nil {
		return nil, err
	}
	lists[0].Items = items

	return &lists[0], nil
}

// AddToShoppingList adds products to the main shopping list (v2).
// This uses PATCH /shoppinglist/v2/items.
func (c *Client) AddToShoppingList(ctx context.Context, items []ListItem) error {
	v2Items := make([]shoppingListItem, 0, len(items))
	for _, item := range items {
		v2 := shoppingListItem{
			Quantity:      max(item.Quantity, 1),
			StrikeThrough: false,
		}
		if item.ProductID > 0 {
			v2.ProductID = item.ProductID
			v2.Type = "SHOPPABLE"
			v2.OriginCode = "PRD"
			v2.Description = item.Name
			v2.SearchTerm = item.Name
		} else {
			v2.Type = "SHOPPABLE"
			v2.OriginCode = "PRD"
			v2.Description = item.Name
		}
		v2Items = append(v2Items, v2)
	}

	body := map[string]any{
		"items": v2Items,
	}

	if err := c.DoRequest(ctx, http.MethodPatch, "/mobile-services/shoppinglist/v2/items", body, nil); err != nil {
		return fmt.Errorf("add to shopping list failed: %w", err)
	}

	return nil
}

// AddProductToShoppingList adds a product to the main shopping list.
func (c *Client) AddProductToShoppingList(ctx context.Context, productID int, quantity int) error {
	return c.AddToShoppingList(ctx, []ListItem{{
		ProductID: productID,
		Quantity:  max(quantity, 1),
	}})
}

// AddFreeTextToShoppingList adds a free-text item (not linked to a product) to the main shopping list.
func (c *Client) AddFreeTextToShoppingList(ctx context.Context, name string, quantity int) error {
	return c.AddToShoppingList(ctx, []ListItem{{
		Name:     name,
		Quantity: max(quantity, 1),
	}})
}

// AddToFavoriteList adds products to a named favorite list (v3) using GraphQL.
// Each item's ProductID and Quantity are sent. Use GetShoppingLists to get list IDs.
func (c *Client) AddToFavoriteList(ctx context.Context, listID string, items []ListItem) error {
	const mutation = `mutation AddProductsToFavoriteList($favoriteListId: String!, $products: [FavoriteListProductMutation!]!) {
  favoriteListProductsAddV2(id: $favoriteListId, products: $products) {
    __typename
    status
    errorMessage
  }
}`

	products := make([]map[string]int, 0, len(items))
	for _, item := range items {
		products = append(products, map[string]int{
			"productId": item.ProductID,
			"quantity":  max(item.Quantity, 1),
		})
	}

	variables := map[string]any{
		"favoriteListId": strings.ToUpper(listID),
		"products":       products,
	}

	var result struct {
		FavoriteListProductsAddV2 struct {
			Status       string `json:"status"`
			ErrorMessage string `json:"errorMessage"`
		} `json:"favoriteListProductsAddV2"`
	}

	if err := c.DoGraphQL(ctx, mutation, variables, &result); err != nil {
		return fmt.Errorf("add to favorite list failed: %w", err)
	}

	if result.FavoriteListProductsAddV2.Status != "SUCCESS" {
		return fmt.Errorf("add to favorite list failed: %s", result.FavoriteListProductsAddV2.ErrorMessage)
	}

	return nil
}

// RemoveFromFavoriteList removes products from a named favorite list using GraphQL.
// It looks up item IDs by product ID, then calls favoriteListProductsDeleteV2.
func (c *Client) RemoveFromFavoriteList(ctx context.Context, listID string, productIDs []int) error {
	items, err := c.GetShoppingListItems(ctx, listID)
	if err != nil {
		return err
	}

	want := make(map[int]bool, len(productIDs))
	for _, id := range productIDs {
		want[id] = true
	}

	var itemIDs []string
	for _, item := range items {
		if want[item.ProductID] {
			itemIDs = append(itemIDs, item.ID)
		}
	}

	if len(itemIDs) == 0 {
		return fmt.Errorf("none of the specified products found in list")
	}

	const mutation = `mutation DeleteProductsFromFavoriteList($favoriteListId: String!, $itemIds: [String!]!) {
  favoriteListProductsDeleteV2(id: $favoriteListId, itemIds: $itemIds) {
    status
    errorMessage
  }
}`

	variables := map[string]any{
		"favoriteListId": strings.ToUpper(listID),
		"itemIds":        itemIDs,
	}

	var result struct {
		FavoriteListProductsDeleteV2 struct {
			Status       string `json:"status"`
			ErrorMessage string `json:"errorMessage"`
		} `json:"favoriteListProductsDeleteV2"`
	}

	if err := c.DoGraphQL(ctx, mutation, variables, &result); err != nil {
		return fmt.Errorf("remove from favorite list failed: %w", err)
	}

	if result.FavoriteListProductsDeleteV2.Status != "SUCCESS" {
		return fmt.Errorf("remove from favorite list failed: %s", result.FavoriteListProductsDeleteV2.ErrorMessage)
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
	if err := c.DoRequest(ctx, http.MethodPatch, path, body, nil); err != nil {
		return fmt.Errorf("check shopping list item failed: %w", err)
	}

	return nil
}

// ClearShoppingList removes all items from the shopping list.
func (c *Client) ClearShoppingList(ctx context.Context) error {
	list, err := c.getShoppingList(ctx)
	if err != nil {
		return err
	}

	var productIDs []int
	for _, item := range list.Items {
		if item.ProductID > 0 {
			productIDs = append(productIDs, item.ProductID)
		}
	}

	if len(productIDs) == 0 {
		return nil
	}

	return c.RemoveFromFavoriteList(ctx, list.ID, productIDs)
}

// ShoppingListToOrder adds all unchecked product items from the shopping list to the order.
// Free-text items (without ProductID) are skipped. This is useful for quickly
// converting your shopping list into an order.
func (c *Client) ShoppingListToOrder(ctx context.Context) error {
	list, err := c.getShoppingList(ctx)
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
