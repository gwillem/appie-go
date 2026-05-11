package appie

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// orderDetailsResponse matches the API response for order details grouped by taxonomy.
type orderDetailsResponse struct {
	OrderID      int    `json:"orderId"`
	OrderState   string `json:"orderState"`
	DeliveryDate string `json:"deliveryDate"`
	DeliveryType string `json:"deliveryType"`
	DeliveryTime struct {
		StartDateTime string `json:"startDateTime"`
		EndDateTime   string `json:"endDateTime"`
	} `json:"deliveryTimePeriod"`
	GroupedProducts []struct {
		TaxonomyName    string `json:"taxonomyName"`
		OrderedProducts []struct {
			Amount   int `json:"amount"`
			Quantity int `json:"quantity"`
			Product  struct {
				WebshopID        int      `json:"webshopId"`
				Title            string   `json:"title"`
				Brand            string   `json:"brand"`
				SalesUnitSize    string   `json:"salesUnitSize"`
				PriceBeforeBonus float64  `json:"priceBeforeBonus"`
				CurrentPrice     *float64 `json:"currentPrice"` // discounted price, nil when discount is per-group (e.g. "2e gratis")
				IsBonus          bool     `json:"isBonus"`
				BonusMechanism   string   `json:"bonusMechanism"`
			} `json:"product"`
		} `json:"orderedProducts"`
	} `json:"groupedProductsInTaxonomy"`
}

func (r *orderDetailsResponse) toOrder() Order {
	var items []OrderItem
	for _, group := range r.GroupedProducts {
		for _, op := range group.OrderedProducts {
			price := Price{Now: op.Product.PriceBeforeBonus}
			if op.Product.CurrentPrice != nil {
				price = Price{Now: *op.Product.CurrentPrice, Was: op.Product.PriceBeforeBonus}
			}
			items = append(items, OrderItem{
				ProductID: op.Product.WebshopID,
				Quantity:  op.Quantity,
				Product: &Product{
					ID:             op.Product.WebshopID,
					Title:          op.Product.Title,
					Brand:          op.Product.Brand,
					UnitSize:       op.Product.SalesUnitSize,
					Price:          price,
					IsBonus:        op.Product.IsBonus,
					BonusMechanism: op.Product.BonusMechanism,
				},
			})
		}
	}

	return Order{
		ID:         strconv.Itoa(r.OrderID),
		State:      r.OrderState,
		Items:      items,
		TotalCount: len(items),
	}
}

// orderSummaryResponse matches the API response for active order summary.
type orderSummaryResponse struct {
	ID           int    `json:"id"`
	Hash         string `json:"hash,omitempty"`
	State        string `json:"state"`
	ShoppingType string `json:"shoppingType"`
	TotalPrice   struct {
		PriceBeforeDiscount float64 `json:"priceBeforeDiscount"`
		PriceAfterDiscount  float64 `json:"priceAfterDiscount"`
		PriceDiscount       float64 `json:"priceDiscount"`
		PriceTotalPayable   float64 `json:"priceTotalPayable"`
	} `json:"totalPrice"`
	DeliveryInformation struct {
		DeliveryDate      string `json:"deliveryDate"`
		DeliveryStartTime string `json:"deliveryStartTime"`
		DeliveryEndTime   string `json:"deliveryEndTime"`
		Address           struct {
			Street      string `json:"street"`
			HouseNumber int    `json:"houseNumber"`
			ZipCode     int    `json:"zipCode"`
			City        string `json:"city"`
		} `json:"address"`
	} `json:"deliveryInformation"`
	OrderedProducts []struct {
		Amount   int `json:"amount"`
		Quantity int `json:"quantity"`
		Product  struct {
			WebshopID int    `json:"webshopId"`
			Title     string `json:"title"`
			Brand     string `json:"brand"`
			Images    []struct {
				URL string `json:"url"`
			} `json:"images"`
		} `json:"product"`
	} `json:"orderedProducts"`
}

func (r *orderSummaryResponse) toOrder() Order {
	items := make([]OrderItem, 0, len(r.OrderedProducts))
	for _, op := range r.OrderedProducts {
		var images []Image
		for _, img := range op.Product.Images {
			images = append(images, Image{URL: img.URL})
		}
		items = append(items, OrderItem{
			ProductID: op.Product.WebshopID,
			Quantity:  op.Quantity,
			Product: &Product{
				ID:     op.Product.WebshopID,
				Title:  op.Product.Title,
				Brand:  op.Product.Brand,
				Images: images,
			},
		})
	}

	return Order{
		ID:            strconv.Itoa(r.ID),
		State:         r.State,
		Items:         items,
		TotalCount:    len(items),
		TotalPrice:    r.TotalPrice.PriceTotalPayable,
		TotalDiscount: r.TotalPrice.PriceDiscount,
	}
}

// GetOrder retrieves the current active order (shopping cart).
// The order ID is cached for use in subsequent order operations.
func (c *Client) GetOrder(ctx context.Context) (*Order, error) {
	var result orderSummaryResponse
	if err := c.DoRequest(ctx, http.MethodGet, "/mobile-services/order/v1/summaries/active?sortBy=DEFAULT", nil, &result); err != nil {
		return nil, fmt.Errorf("get order failed: %w", err)
	}

	// Store order ID for subsequent requests
	c.mu.Lock()
	c.orderID = strconv.Itoa(result.ID)
	c.orderHash = result.Hash
	c.mu.Unlock()

	fmt.Fprintf(os.Stderr, "[appie-go DEBUG] GetOrder: id=%s hash=%q state=%s\n", strconv.Itoa(result.ID), result.Hash, result.State)

	order := result.toOrder()
	return &order, nil
}

// GetOrderDetails retrieves the details of a specific order by ID.
// Unlike GetOrder, this works for any order (not just the active one) and
// returns products grouped by taxonomy (category).
func (c *Client) GetOrderDetails(ctx context.Context, orderID int) (*Order, error) {
	path := fmt.Sprintf("/mobile-services/order/v1/%d/details-grouped-by-taxonomy", orderID)

	var result orderDetailsResponse
	if err := c.DoRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get order details failed: %w", err)
	}

	order := result.toOrder()
	return &order, nil
}

// SetOrderID manually sets the order ID for subsequent API requests.
// This is useful when selecting a specific order from fulfillments.
func (c *Client) SetOrderID(id int) {
	c.mu.Lock()
	c.orderID = strconv.Itoa(id)
	c.mu.Unlock()
}

// AddToOrder adds or updates items in the current order.
// If an item already exists, its quantity is updated. Set quantity to 0 to remove.
//
// Example:
//
//	err := client.AddToOrder(ctx, []appie.OrderItem{
//	    {ProductID: 123456, Quantity: 2},
//	})
func (c *Client) AddToOrder(ctx context.Context, items []OrderItem) error {
	type itemRequest struct {
		ProductID     int    `json:"productId"`
		Quantity      int    `json:"quantity"`
		OriginCode    string `json:"originCode"`
		Description   string `json:"description"`
		Strikethrough bool   `json:"strikethrough"`
	}

	// Merge duplicates: the API rejects requests with duplicate product IDs.
	merged := make(map[int]int, len(items))
	for _, item := range items {
		merged[item.ProductID] += item.Quantity
	}

	reqItems := make([]itemRequest, 0, len(merged))
	for pid, qty := range merged {
		reqItems = append(reqItems, itemRequest{
			ProductID:     pid,
			Quantity:      qty,
			OriginCode:    "PRD",
			Description:   "",
			Strikethrough: false,
		})
	}

	body := map[string]any{
		"items": reqItems,
	}

	c.mu.RLock()
	debugID := c.orderID
	debugHash := c.orderHash
	c.mu.RUnlock()
	fmt.Fprintf(os.Stderr, "[appie-go DEBUG] AddToOrder: orderID=%q hash=%q items=%+v\n", debugID, debugHash, reqItems)

	if err := c.DoRequest(ctx, http.MethodPut, "/mobile-services/order/v1/items?sortBy=DEFAULT", body, nil); err != nil {
		fmt.Fprintf(os.Stderr, "[appie-go DEBUG] AddToOrder ERROR: %v\n", err)
		return fmt.Errorf("add to order failed: %w", err)
	}
	fmt.Fprintf(os.Stderr, "[appie-go DEBUG] AddToOrder SUCCESS\n")

	return nil
}

const basketItemsUpdateMutation = `mutation UpdateMyListBasket($items: [BasketMutation!]!, $input: BasketInput) {
  basketItemsUpdate(items: $items, input: $input) {
    __typename
    status
  }
}`

// AddToBasket adds or updates items in the active basket using the GraphQL basketItemsUpdate mutation.
// This is the correct method for modifying a REOPENED fulfillment order.
func (c *Client) AddToBasket(ctx context.Context, items []OrderItem) error {
	type basketMutation struct {
		ID              int  `json:"id"`
		IsStrikethrough bool `json:"isStrikethrough"`
		NewPosition     int  `json:"newPosition"`
		Quantity        int  `json:"quantity"`
	}

	mutations := make([]basketMutation, 0, len(items))
	for _, item := range items {
		mutations = append(mutations, basketMutation{
			ID:              item.ProductID,
			IsStrikethrough: false,
			NewPosition:     0,
			Quantity:        item.Quantity,
		})
	}

	vars := map[string]any{
		"items": mutations,
		"input": nil,
	}

	type response struct {
		BasketItemsUpdate struct {
			Status string `json:"status"`
		} `json:"basketItemsUpdate"`
	}

	var resp response
	if err := c.DoGraphQL(ctx, basketItemsUpdateMutation, vars, &resp); err != nil {
		return fmt.Errorf("basket update failed: %w", err)
	}

	fmt.Fprintf(os.Stderr, "[appie-go DEBUG] AddToBasket status: %s\n", resp.BasketItemsUpdate.Status)
	return nil
}

const basketSummaryQuery = `query FetchBasketSummary {
  basket {
    summary {
      lastUserChangeDateTime
    }
  }
}`

// GetBasketLastModified returns the lastUserChangeDateTime from the basket summary.
// This timestamp is required for CheckoutConfirmOrder.
func (c *Client) GetBasketLastModified(ctx context.Context) (string, error) {
	type response struct {
		Basket struct {
			Summary struct {
				LastUserChangeDateTime string `json:"lastUserChangeDateTime"`
			} `json:"summary"`
		} `json:"basket"`
	}

	var resp response
	if err := c.DoGraphQL(ctx, basketSummaryQuery, nil, &resp); err != nil {
		return "", fmt.Errorf("get basket summary failed: %w", err)
	}

	fmt.Fprintf(os.Stderr, "[appie-go DEBUG] GetBasketLastModified: %s\n", resp.Basket.Summary.LastUserChangeDateTime)
	return resp.Basket.Summary.LastUserChangeDateTime, nil
}

const confirmOrderMutation = `mutation CheckoutConfirmOrder($orderId: Int!, $orderInfo: CheckoutConfirmOrderPayloadV4!) {
  checkoutConfirmOrderV4(orderId: $orderId, orderInfo: $orderInfo) {
    __typename
    status
  }
}`

// ConfirmOrder resubmits a reopened order using the checkout confirm mutation.
// orderLastModified must be the lastUserChangeDateTime from the basket summary.
func (c *Client) ConfirmOrder(ctx context.Context, orderID int, orderLastModified string) error {
	type orderInfo struct {
		Channel                       string `json:"channel"`
		OrderLastModified             string `json:"orderLastModified"`
		PaymentMethod                 string `json:"paymentMethod"`
		PurchaseStampsBookletsApplied int    `json:"purchaseStampsBookletsApplied"`
	}

	vars := map[string]any{
		"orderId": orderID,
		"orderInfo": orderInfo{
			Channel:                       "IOS",
			OrderLastModified:             orderLastModified,
			PaymentMethod:                 "PAY_AT_DELIVERY",
			PurchaseStampsBookletsApplied: 0,
		},
	}

	type response struct {
		CheckoutConfirmOrderV4 struct {
			Status string `json:"status"`
		} `json:"checkoutConfirmOrderV4"`
	}

	var resp response
	if err := c.DoGraphQL(ctx, confirmOrderMutation, vars, &resp); err != nil {
		return fmt.Errorf("confirm order failed: %w", err)
	}

	fmt.Fprintf(os.Stderr, "[appie-go DEBUG] ConfirmOrder status: %s\n", resp.CheckoutConfirmOrderV4.Status)
	return nil
}

// RemoveFromOrder removes an item from the current order by setting quantity to 0.
func (c *Client) RemoveFromOrder(ctx context.Context, productID int) error {
	return c.AddToOrder(ctx, []OrderItem{{ProductID: productID, Quantity: 0}})
}

// UpdateOrderItem updates the quantity of an item in the order.
func (c *Client) UpdateOrderItem(ctx context.Context, productID, quantity int) error {
	return c.AddToOrder(ctx, []OrderItem{{ProductID: productID, Quantity: quantity}})
}

// ClearOrder removes all items from the current order.
func (c *Client) ClearOrder(ctx context.Context) error {
	order, err := c.GetOrder(ctx)
	if err != nil {
		return err
	}

	if len(order.Items) == 0 {
		return nil
	}

	removals := make([]OrderItem, 0, len(order.Items))
	for _, item := range order.Items {
		removals = append(removals, OrderItem{ProductID: item.ProductID, Quantity: 0})
	}

	return c.AddToOrder(ctx, removals)
}

// GetOrderSummary retrieves the order summary/totals.
func (c *Client) GetOrderSummary(ctx context.Context) (*OrderSummary, error) {
	var result orderSummaryResponse
	if err := c.DoRequest(ctx, http.MethodGet, "/mobile-services/order/v1/summaries/active?sortBy=DEFAULT", nil, &result); err != nil {
		return nil, fmt.Errorf("get order summary failed: %w", err)
	}

	return &OrderSummary{
		TotalItems:    len(result.OrderedProducts),
		TotalPrice:    result.TotalPrice.PriceTotalPayable,
		TotalDiscount: result.TotalPrice.PriceDiscount,
	}, nil
}

const reopenOrderMutation = `mutation OrderReopen($id: Int!) {
  orderReopen(id: $id) {
    status
    errorMessage
  }
}`

// ReopenOrder reopens a submitted order so items can be added or modified.
// This is required before calling AddToOrder on a fulfillment-selected order.
func (c *Client) ReopenOrder(ctx context.Context, orderID int) error {
	type reopenResponse struct {
		OrderReopen struct {
			Status       string `json:"status"`
			ErrorMessage string `json:"errorMessage"`
		} `json:"orderReopen"`
	}

	var resp reopenResponse
	vars := map[string]any{"id": orderID}
	if err := c.DoGraphQL(ctx, reopenOrderMutation, vars, &resp); err != nil {
		return fmt.Errorf("reopen order failed: %w", err)
	}

	if resp.OrderReopen.Status != "SUCCESS" {
		return fmt.Errorf("reopen order failed: %s", resp.OrderReopen.ErrorMessage)
	}

	// The reopen mutation doesn't return a hash. Set the order ID manually,
	// then call GetOrder to fetch the active summary which includes the hash.
	c.SetOrderID(orderID)

	if _, err := c.GetOrder(ctx); err != nil {
		return fmt.Errorf("reopen order succeeded but failed to fetch hash: %w", err)
	}

	return nil
}

const revertOrderMutation = `mutation OrderRevert($id: Int!) {
  orderRevert(id: $id) {
    status
    errorMessage
  }
}`

// RevertOrder reverts a reopened order back to its submitted state.
// Use this after ReopenOrder to deactivate the order as the active one.
func (c *Client) RevertOrder(ctx context.Context, orderID int) error {
	type revertResponse struct {
		OrderRevert struct {
			Status       string `json:"status"`
			ErrorMessage string `json:"errorMessage"`
		} `json:"orderRevert"`
	}

	var resp revertResponse
	vars := map[string]any{"id": orderID}
	if err := c.DoGraphQL(ctx, revertOrderMutation, vars, &resp); err != nil {
		return fmt.Errorf("revert order failed: %w", err)
	}

	if resp.OrderRevert.Status != "SUCCESS" {
		return fmt.Errorf("revert order failed: %s", resp.OrderRevert.ErrorMessage)
	}

	// Clear client-side order state
	c.mu.Lock()
	c.orderID = ""
	c.orderHash = ""
	c.mu.Unlock()

	return nil
}

const fulfillmentsQuery = `query OrderFulfillments {
  orderFulfillments(status: OPEN) {
    result {
      orderId
      statusCode
      statusDescription
      shoppingType
      transactionCompleted
      modifiable
      totalPrice {
        totalPrice { amount }
      }
      delivery {
        status
        method
        slot {
          date
          dateDisplay
          timeDisplay
          startTime
          endTime
        }
        address {
          street
          houseNumber
          houseNumberExtra
          city
          postalCode
        }
      }
    }
  }
}`

type fulfillmentsResponse struct {
	OrderFulfillments struct {
		Result []fulfillmentResult `json:"result"`
	} `json:"orderFulfillments"`
}

type fulfillmentResult struct {
	OrderID              int    `json:"orderId"`
	StatusCode           int    `json:"statusCode"`
	StatusDescription    string `json:"statusDescription"`
	ShoppingType         string `json:"shoppingType"`
	TransactionCompleted bool   `json:"transactionCompleted"`
	Modifiable           bool   `json:"modifiable"`
	TotalPrice           struct {
		TotalPrice struct {
			Amount float64 `json:"amount"`
		} `json:"totalPrice"`
	} `json:"totalPrice"`
	Delivery FulfillmentDelivery `json:"delivery"`
}

// GetFulfillments retrieves all open (scheduled) order fulfillments.
// These are orders that have been submitted and are awaiting delivery.
func (c *Client) GetFulfillments(ctx context.Context) ([]Fulfillment, error) {
	var resp fulfillmentsResponse
	if err := c.DoGraphQL(ctx, fulfillmentsQuery, nil, &resp); err != nil {
		return nil, fmt.Errorf("get fulfillments failed: %w", err)
	}

	results := resp.OrderFulfillments.Result
	fulfillments := make([]Fulfillment, 0, len(results))

	for _, r := range results {
		fulfillments = append(fulfillments, Fulfillment{
			OrderID:              r.OrderID,
			Status:               r.Delivery.Status,
			StatusDescription:    r.StatusDescription,
			ShoppingType:         r.ShoppingType,
			TotalPrice:           r.TotalPrice.TotalPrice.Amount,
			TransactionCompleted: r.TransactionCompleted,
			Modifiable:           r.Modifiable,
			Delivery:             r.Delivery,
		})
	}

	return fulfillments, nil
}
