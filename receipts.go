package appie

import (
	"context"
	"fmt"
	"net/http"
)

// receiptsResponse matches the API response for GET /mobile-services/v1/receipts
type receiptsResponse struct {
	Receipts []receiptSummary `json:"receipts"`
}

type receiptSummary struct {
	TransactionID string  `json:"transactionId"`
	DateTime      string  `json:"datetime"`
	StoreID       int     `json:"storeId"`
	StoreName     string  `json:"storeName"`
	Total         float64 `json:"total"`
}

// receiptDetailResponse matches the API response for GET /mobile-services/v2/receipts/{id}
type receiptDetailResponse struct {
	TransactionID string              `json:"transactionId"`
	DateTime      string              `json:"datetime"`
	StoreID       int                 `json:"storeId"`
	StoreName     string              `json:"storeName"`
	Total         float64             `json:"total"`
	ReceiptItems  []receiptItemDetail `json:"receiptItems"`
}

type receiptItemDetail struct {
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Amount      float64 `json:"amount"`
	UnitPrice   float64 `json:"unitPrice"`
	ProductID   int     `json:"productId,omitempty"`
}

// GetReceipts retrieves the list of in-store receipts (kassabonnen) for the authenticated user.
func (c *Client) GetReceipts(ctx context.Context) ([]Receipt, error) {
	var result receiptsResponse
	if err := c.DoRequest(ctx, http.MethodGet, "/mobile-services/v1/receipts", nil, &result); err != nil {
		return nil, fmt.Errorf("get receipts failed: %w", err)
	}

	receipts := make([]Receipt, 0, len(result.Receipts))
	for _, r := range result.Receipts {
		receipts = append(receipts, Receipt{
			TransactionID: r.TransactionID,
			Date:          r.DateTime,
			StoreID:       r.StoreID,
			StoreName:     r.StoreName,
			TotalAmount:   r.Total,
		})
	}

	return receipts, nil
}

// GetReceipt retrieves the details of a specific in-store receipt by transaction ID.
func (c *Client) GetReceipt(ctx context.Context, transactionID string) (*Receipt, error) {
	path := fmt.Sprintf("/mobile-services/v2/receipts/%s", transactionID)

	var result receiptDetailResponse
	if err := c.DoRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get receipt failed: %w", err)
	}

	items := make([]ReceiptItem, 0, len(result.ReceiptItems))
	for _, item := range result.ReceiptItems {
		items = append(items, ReceiptItem{
			Description: item.Description,
			Quantity:    item.Quantity,
			Amount:      item.Amount,
			UnitPrice:   item.UnitPrice,
			ProductID:   item.ProductID,
		})
	}

	return &Receipt{
		TransactionID: result.TransactionID,
		Date:          result.DateTime,
		StoreID:       result.StoreID,
		StoreName:     result.StoreName,
		TotalAmount:   result.Total,
		Items:         items,
	}, nil
}
