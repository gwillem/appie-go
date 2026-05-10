package appie

import (
	"context"
	"fmt"
	"strings"
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

	posIDs := make([]int, 0, len(items))
	for _, it := range items {
		posIDs = append(posIDs, it.ProductID)
	}
	if mapping, err := c.ConvertPOSIDs(ctx, posIDs); err != nil {
		c.logger.Printf("ConvertPOSIDs: %v (continuing without webshop ids)", err)
	} else {
		for i := range items {
			items[i].WebshopID = mapping[items[i].ProductID]
		}
	}

	return &Receipt{
		TransactionID: details.ID,
		Items:         items,
		Discounts:     discounts,
		Payments:      payments,
	}, nil
}

// ConvertPOSIDs maps POS-receipt productIds (the integer found on
// PosReceiptProduct.id / ReceiptItem.ProductID) to AH webshop product ids
// (the "wi<id>" used on ah.nl). Resolution happens in a single GraphQL
// request using aliased subfields, regardless of how many ids are passed.
//
// The returned map only contains entries that resolved successfully.
// Unknown ids, ids for which the API returned its -1 "no conversion"
// sentinel, and ids whose alias came back null are silently omitted, so
// callers can use a plain `out[id]` lookup with a `0` zero-value meaning
// "unresolved". The map is always non-nil (including on error), so it is
// safe to read from even when the caller chooses to ignore err.
//
// Only feed this method ids that actually came from PosReceiptProduct.id.
// The underlying resolver does not validate the source id space; passing
// arbitrary integers can return real but unrelated webshop ids.
func (c *Client) ConvertPOSIDs(ctx context.Context, ids []int) (map[int]int, error) {
	out := make(map[int]int)
	if len(ids) == 0 {
		return out, nil
	}

	seen := make(map[int]struct{}, len(ids))
	unique := make([]int, 0, len(ids))
	for _, id := range ids {
		if id <= 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		unique = append(unique, id)
	}
	if len(unique) == 0 {
		return out, nil
	}

	// Aliases must be static identifiers, so the ids are inlined into
	// the query string rather than passed as variables. They originate
	// from the AH API itself (PosReceiptProduct.id) and the dedupe loop
	// above guarantees they are positive integers — no injection surface.
	var b strings.Builder
	b.WriteString("query Convert {")
	for i, id := range unique {
		fmt.Fprintf(&b, " p%d: productConvertId(sourceId: %d)", i, id)
	}
	b.WriteString(" }")

	// productConvertId is declared as Int (nullable) in the schema —
	// decode through *int so a null alias becomes a normal "unresolved"
	// instead of a JSON unmarshal failure that would torch the batch.
	var resp map[string]*int
	if err := c.DoGraphQL(ctx, b.String(), nil, &resp); err != nil {
		return out, fmt.Errorf("convert pos ids: %w", err)
	}

	for i, id := range unique {
		v := resp[fmt.Sprintf("p%d", i)]
		if v == nil || *v <= 0 {
			continue
		}
		out[id] = *v
	}
	return out, nil
}
