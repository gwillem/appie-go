package appie

import (
	"context"
	"fmt"
)

const fetchPosReceiptsQuery = `query FetchPosReceipts($offset: Int!, $limit: Int!) {
	posReceiptsPage(pagination: {offset: $offset, limit: $limit}) {
		posReceipts {
			id
			dateTime
			totalAmount {
				amount
			}
		}
	}
}`

const fetchPosReceiptDetailsQuery = `query FetchReceipt($id: String!) {
	posReceiptDetails(id: $id) {
		id
		memberId
		products {
			id
			quantity
			name
			price {
				amount
			}
			amount {
				amount
			}
		}
		discounts {
			name
			amount {
				amount
			}
		}
		payments {
			method
			amount {
				amount
			}
		}
	}
}`

type posReceiptsResponse struct {
	PosReceiptsPage struct {
		PosReceipts []struct {
			ID          string `json:"id"`
			DateTime    string `json:"dateTime"`
			TotalAmount struct {
				Amount float64 `json:"amount"`
			} `json:"totalAmount"`
		} `json:"posReceipts"`
	} `json:"posReceiptsPage"`
}

type posReceiptDetailsResponse struct {
	PosReceiptDetails struct {
		ID       string `json:"id"`
		Products []struct {
			ID       int    `json:"id"`
			Quantity int    `json:"quantity"`
			Name     string `json:"name"`
			Price    *struct {
				Amount float64 `json:"amount"`
			} `json:"price"`
			Amount struct {
				Amount float64 `json:"amount"`
			} `json:"amount"`
		} `json:"products"`
		Discounts []struct {
			Name   string `json:"name"`
			Amount struct {
				Amount float64 `json:"amount"`
			} `json:"amount"`
		} `json:"discounts"`
		Payments []struct {
			Method string `json:"method"`
			Amount struct {
				Amount float64 `json:"amount"`
			} `json:"amount"`
		} `json:"payments"`
	} `json:"posReceiptDetails"`
}

// GetReceipts retrieves the list of in-store receipts (kassabonnen) for the authenticated user.
func (c *Client) GetReceipts(ctx context.Context) ([]Receipt, error) {
	vars := map[string]any{
		"offset": 0,
		"limit":  100,
	}

	var resp posReceiptsResponse
	if err := c.DoGraphQL(ctx, fetchPosReceiptsQuery, vars, &resp); err != nil {
		return nil, fmt.Errorf("get receipts failed: %w", err)
	}

	posReceipts := resp.PosReceiptsPage.PosReceipts
	receipts := make([]Receipt, 0, len(posReceipts))
	for _, r := range posReceipts {
		receipts = append(receipts, Receipt{
			TransactionID: r.ID,
			Date:          r.DateTime,
			TotalAmount:   r.TotalAmount.Amount,
		})
	}

	return receipts, nil
}

// GetReceipt retrieves the details of a specific in-store receipt by ID.
func (c *Client) GetReceipt(ctx context.Context, id string) (*Receipt, error) {
	vars := map[string]any{
		"id": id,
	}

	var resp posReceiptDetailsResponse
	if err := c.DoGraphQL(ctx, fetchPosReceiptDetailsQuery, vars, &resp); err != nil {
		return nil, fmt.Errorf("get receipt failed: %w", err)
	}

	details := resp.PosReceiptDetails
	items := make([]ReceiptItem, 0, len(details.Products))
	for _, p := range details.Products {
		var unitPrice float64
		if p.Price != nil {
			unitPrice = p.Price.Amount
		}
		items = append(items, ReceiptItem{
			Description: p.Name,
			Quantity:    p.Quantity,
			Amount:      p.Amount.Amount,
			UnitPrice:   unitPrice,
			ProductID:   p.ID,
		})
	}

	discounts := make([]ReceiptDiscount, 0, len(details.Discounts))
	for _, d := range details.Discounts {
		discounts = append(discounts, ReceiptDiscount{
			Name:   d.Name,
			Amount: d.Amount.Amount,
		})
	}

	payments := make([]ReceiptPayment, 0, len(details.Payments))
	for _, p := range details.Payments {
		payments = append(payments, ReceiptPayment{
			Method: p.Method,
			Amount: p.Amount.Amount,
		})
	}

	return &Receipt{
		TransactionID: details.ID,
		Items:         items,
		Discounts:     discounts,
		Payments:      payments,
	}, nil
}
