package appie

import (
	"context"
	"fmt"
)

// FulfillmentStatus represents the delivery status of an order fulfillment.
type FulfillmentStatus string

const (
	// Granular delivery status constants returned by the AH API.
	FulfillmentSubmitted        FulfillmentStatus = "SUBMITTED"
	FulfillmentSubmittedWithETA FulfillmentStatus = "SUBMITTED_WITH_SHORTREC_ETA"
	FulfillmentDelivered        FulfillmentStatus = "DELIVERED"
	FulfillmentCancelled        FulfillmentStatus = "CANCELLED"

	// FulfillmentOpen and FulfillmentClosed are the GraphQL filter values accepted
	// by the API. They map to groups of granular statuses:
	//   OPEN   → SUBMITTED, SUBMITTED_WITH_SHORTREC_ETA
	//   CLOSED → DELIVERED, CANCELLED
	// These are kept as the default filter values and for backwards compatibility.
	FulfillmentOpen   FulfillmentStatus = "OPEN"
	FulfillmentClosed FulfillmentStatus = "CLOSED"
)

// FulfillmentOption configures a GetFulfillments call.
type FulfillmentOption func(*fulfillmentQuery)

type fulfillmentQuery struct {
	status FulfillmentStatus
	size   int
}

// WithStatus filters fulfillments by the given order status.
func WithStatus(s FulfillmentStatus) FulfillmentOption {
	return func(q *fulfillmentQuery) { q.status = s }
}

// WithSize limits the number of fulfillments returned (0 uses the API default).
func WithSize(n int) FulfillmentOption {
	return func(q *fulfillmentQuery) { q.size = n }
}

const fetchFulfillmentsQuery = `query FetchOrderFulfillments($status: FulfillmentStatus, $size: Int) {
  orderFulfillments(status: $status, size: $size) {
    result {
      __typename
      ...OrderFulfillment
    }
  }
}

fragment OrderFulfillment on Fulfillment {
  orderId
  transactionCompleted
  modifiable
  delivery {
    deliveryMessage
    method
    slot {
      date
      startTime
      endTime
    }
    homeShopCenterId
    status
  }
  reopenable
  isSubscriptionOrder
  totalPrice {
    totalPrice {
      amount
    }
  }
}`

const fetchFulfillmentDetailQuery = `query FetchFullOrderFulfillment($orderId: Int!) {
  orderFulfillment(orderId: $orderId) {
    __typename
    ...FullOrderFulfillment
  }
}

fragment OrderProductFragment on Product {
  priceV2 {
    now {
      amount
    }
    was {
      amount
    }
  }
  brand
  listPrice {
    amount
  }
  hqId
  imagePack {
    medium {
      width
      height
      url
    }
  }
  category
  salesUnitSize
  title
  id
}

fragment OrderLineFragment on OrderLineV2 {
  quantity
  allocatedQuantity
  product {
    __typename
    ...OrderProductFragment
  }
}

fragment FullOrderFulfillment on Fulfillment {
  orderId
  closingDateTime
  cancellable
  delivery {
    slot {
      date
      startTime
      endTime
    }
    homeShopCenterId
    status
    method
    address {
      postalCode
      city
      countryCode
      houseNumber
      houseNumberExtra
      street
    }
  }
  costOverview {
    invoiceId
  }
  reopenable
  orderLines(sortType: TAXONOMY) {
    __typename
    ...OrderLineFragment
  }
}`

type fulfillmentsGQLResponse struct {
	OrderFulfillments struct {
		Result []fulfillmentGQLResult `json:"result"`
	} `json:"orderFulfillments"`
}

type fulfillmentGQLResult struct {
	OrderID              int  `json:"orderId"`
	TransactionCompleted bool `json:"transactionCompleted"`
	Modifiable           bool `json:"modifiable"`
	Reopenable           bool `json:"reopenable"`
	IsSubscriptionOrder  bool `json:"isSubscriptionOrder"`
	TotalPrice           struct {
		TotalPrice struct {
			Amount float64 `json:"amount"`
		} `json:"totalPrice"`
	} `json:"totalPrice"`
	Delivery struct {
		DeliveryMessage  string `json:"deliveryMessage"`
		Method           string `json:"method"`
		HomeShopCenterID int    `json:"homeShopCenterId"`
		Status           string `json:"status"`
		Slot             struct {
			Date      string `json:"date"`
			StartTime string `json:"startTime"`
			EndTime   string `json:"endTime"`
		} `json:"slot"`
	} `json:"delivery"`
}

type fulfillmentDetailGQLResponse struct {
	OrderFulfillment struct {
		OrderID         int    `json:"orderId"`
		ClosingDateTime string `json:"closingDateTime"`
		Cancellable     bool   `json:"cancellable"`
		Reopenable      bool   `json:"reopenable"`
		Delivery        struct {
			Slot struct {
				Date      string `json:"date"`
				StartTime string `json:"startTime"`
				EndTime   string `json:"endTime"`
			} `json:"slot"`
			HomeShopCenterID int    `json:"homeShopCenterId"`
			Status           string `json:"status"`
			Method           string `json:"method"`
			Address          struct {
				PostalCode       string `json:"postalCode"`
				City             string `json:"city"`
				CountryCode      string `json:"countryCode"`
				HouseNumber      int    `json:"houseNumber"`
				HouseNumberExtra string `json:"houseNumberExtra"`
				Street           string `json:"street"`
			} `json:"address"`
		} `json:"delivery"`
		// CostOverview sometimes is nil for orders, so we specify it as a pointer to allow it to be missing.
		CostOverview *struct {
			InvoiceID string `json:"invoiceId"`
		} `json:"costOverview"`
		OrderLines []struct {
			Quantity          int `json:"quantity"`
			AllocatedQuantity int `json:"allocatedQuantity"`
			Product           *struct {
				ID      int    `json:"id"`
				Title   string `json:"title"`
				Brand   string `json:"brand"`
				HqID    int    `json:"hqId"`
				PriceV2 struct {
					Now *struct {
						Amount float64 `json:"amount"`
					} `json:"now"`
					Was *struct {
						Amount float64 `json:"amount"`
					} `json:"was"`
				} `json:"priceV2"`
				ListPrice *struct {
					Amount float64 `json:"amount"`
				} `json:"listPrice"`
				ImagePack []struct {
					Medium *struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"medium"`
				} `json:"imagePack"`
				Category      string `json:"category"`
				SalesUnitSize string `json:"salesUnitSize"`
			} `json:"product"`
		} `json:"orderLines"`
	} `json:"orderFulfillment"`
}

// GetFulfillments retrieves order fulfillments via GraphQL. Without options it
// returns all open orders (SUBMITTED and SUBMITTED_WITH_SHORTREC_ETA).
// Use WithStatus to filter by a specific status and WithSize to limit results.
func (c *Client) GetFulfillments(ctx context.Context, opts ...FulfillmentOption) ([]Fulfillment, error) {
	q := fulfillmentQuery{status: FulfillmentOpen}
	for _, opt := range opts {
		opt(&q)
	}

	// Map granular delivery status constants to the OPEN/CLOSED filter the
	// GraphQL API accepts. The broad filter constants pass through unchanged.
	gqlStatus := string(q.status)
	switch q.status {
	case FulfillmentSubmitted, FulfillmentSubmittedWithETA:
		gqlStatus = "OPEN"
	case FulfillmentDelivered, FulfillmentCancelled:
		gqlStatus = "CLOSED"
	}

	vars := map[string]any{"status": gqlStatus}
	if q.size > 0 {
		vars["size"] = q.size
	}

	var resp fulfillmentsGQLResponse
	if err := c.DoGraphQL(ctx, fetchFulfillmentsQuery, vars, &resp); err != nil {
		return nil, fmt.Errorf("get fulfillments failed: %w", err)
	}

	results := resp.OrderFulfillments.Result
	fulfillments := make([]Fulfillment, 0, len(results))
	for _, r := range results {
		fulfillments = append(fulfillments, Fulfillment{
			OrderID:             r.OrderID,
			Status:              r.Delivery.Status,
			ShoppingType:        r.Delivery.Method,
			TotalPrice:          r.TotalPrice.TotalPrice.Amount,
			TransactionCompleted: r.TransactionCompleted,
			Modifiable:          r.Modifiable,
			Reopenable:          r.Reopenable,
			IsSubscriptionOrder: r.IsSubscriptionOrder,
			DeliveryMessage:     r.Delivery.DeliveryMessage,
			Delivery: FulfillmentDelivery{
				Status: r.Delivery.Status,
				Method: r.Delivery.Method,
				Slot: DeliverySlot{
					Date:      r.Delivery.Slot.Date,
					StartTime: r.Delivery.Slot.StartTime,
					EndTime:   r.Delivery.Slot.EndTime,
				},
			},
		})
	}
	return fulfillments, nil
}

// GetFulfillmentDetail retrieves the full details of a specific order,
// including all order lines with product information.
func (c *Client) GetFulfillmentDetail(ctx context.Context, orderID int) (*FulfillmentDetail, error) {
	vars := map[string]any{
		"orderId": orderID,
	}

	var resp fulfillmentDetailGQLResponse
	if err := c.DoGraphQL(ctx, fetchFulfillmentDetailQuery, vars, &resp); err != nil {
		return nil, fmt.Errorf("get fulfillment detail failed: %w", err)
	}

	r := resp.OrderFulfillment
	detail := &FulfillmentDetail{
		OrderID:         r.OrderID,
		ClosingDateTime: r.ClosingDateTime,
		Cancellable:     r.Cancellable,
		Reopenable:      r.Reopenable,
		Delivery: FulfillmentDelivery{
			Status: r.Delivery.Status,
			Method: r.Delivery.Method,
			Slot: DeliverySlot{
				Date:      r.Delivery.Slot.Date,
				StartTime: r.Delivery.Slot.StartTime,
				EndTime:   r.Delivery.Slot.EndTime,
			},
			Address: Address{
				Street:           r.Delivery.Address.Street,
				HouseNumber:      r.Delivery.Address.HouseNumber,
				HouseNumberExtra: r.Delivery.Address.HouseNumberExtra,
				PostalCode:       r.Delivery.Address.PostalCode,
				City:             r.Delivery.Address.City,
				CountryCode:      r.Delivery.Address.CountryCode,
			},
		},
	}
	if r.CostOverview != nil {
		detail.InvoiceID = r.CostOverview.InvoiceID
	}

	for _, ol := range r.OrderLines {
		line := OrderLine{
			Quantity:          ol.Quantity,
			AllocatedQuantity: ol.AllocatedQuantity,
		}
		if ol.Product != nil {
			p := &OrderLineProduct{
				ID:            ol.Product.ID,
				Title:         ol.Product.Title,
				Brand:         ol.Product.Brand,
				SalesUnitSize: ol.Product.SalesUnitSize,
				Category:      ol.Product.Category,
			}
			if ol.Product.PriceV2.Now != nil {
				p.CurrentPrice = ol.Product.PriceV2.Now.Amount
			}
			if ol.Product.PriceV2.Was != nil {
				p.WasPrice = ol.Product.PriceV2.Was.Amount
			}
			if ol.Product.ListPrice != nil {
				p.ListPrice = ol.Product.ListPrice.Amount
			}
			if len(ol.Product.ImagePack) > 0 && ol.Product.ImagePack[0].Medium != nil {
				p.ImageURL = ol.Product.ImagePack[0].Medium.URL
			}
			line.Product = p
		}
		detail.OrderLines = append(detail.OrderLines, line)
	}

	return detail, nil
}

// GetOrderHistory is a convenience wrapper for GetFulfillments with status CLOSED.
// size controls how many orders to return; 0 uses a large default (200).
func (c *Client) GetOrderHistory(ctx context.Context, size int) ([]Fulfillment, error) {
	if size <= 0 {
		size = 200
	}
	return c.GetFulfillments(ctx, WithStatus(FulfillmentClosed), WithSize(size))
}

// GetOrderHistoryDetail is a convenience wrapper for GetFulfillmentDetail.
func (c *Client) GetOrderHistoryDetail(ctx context.Context, orderID int) (*FulfillmentDetail, error) {
	return c.GetFulfillmentDetail(ctx, orderID)
}
