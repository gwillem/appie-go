package appie

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// searchResponse matches the API response for product search
type searchResponse struct {
	Products []productResponse `json:"products"`
	Page     struct {
		Number        int `json:"number"`
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
	} `json:"page"`
}

type productResponse struct {
	WebshopID            int             `json:"webshopId"`
	HqID                 int             `json:"hqId"`
	Title                string          `json:"title"`
	Brand                string          `json:"brand"`
	SalesUnitSize        string          `json:"salesUnitSize"`
	UnitPriceDescription string          `json:"unitPriceDescription"`
	Images               []imageResponse `json:"images"`
	CurrentPrice         float64         `json:"currentPrice"`
	PriceBeforeBonus     float64         `json:"priceBeforeBonus"`
	IsBonus              bool            `json:"isBonus"`
	BonusMechanism       string          `json:"bonusMechanism"`
	MainCategory         string          `json:"mainCategory"`
	SubCategory          string          `json:"subCategory"`
	NutriScore           string          `json:"nutriscore"`
	AvailableOnline      bool            `json:"availableOnline"`
	IsPreviouslyBought   bool            `json:"isPreviouslyBought"`
	IsOrderable          bool            `json:"isOrderable"`
	PropertyIcons        []string        `json:"propertyIcons"`
}

type imageResponse struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// GraphQL query for fetching product nutritional info via tradeItem
const fetchProductNutritionQuery = `query FetchProduct($productId: Int!) {
  product(id: $productId) {
    id
    tradeItem {
      nutritions {
        nutrients {
          type
          name
          value
        }
      }
    }
  }
}`

type productNutritionResponse struct {
	Product struct {
		ID        int `json:"id"`
		TradeItem *struct {
			Nutritions []struct {
				Nutrients []struct {
					Type  string `json:"type"`
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"nutrients"`
			} `json:"nutritions"`
		} `json:"tradeItem"`
	} `json:"product"`
}

func (p *productResponse) toProduct() Product {
	var images []Image
	for _, img := range p.Images {
		images = append(images, Image{
			URL:    img.URL,
			Width:  img.Width,
			Height: img.Height,
		})
	}

	price := p.CurrentPrice
	if price == 0 {
		price = p.PriceBeforeBonus
	}

	return Product{
		ID:                   p.WebshopID,
		WebshopID:            strconv.Itoa(p.WebshopID),
		Title:                p.Title,
		Brand:                p.Brand,
		Category:             p.MainCategory,
		SubCategory:          p.SubCategory,
		Price:                Price{Now: price, Was: p.PriceBeforeBonus},
		Images:               images,
		NutriScore:           p.NutriScore,
		IsBonus:              p.IsBonus,
		BonusMechanism:       p.BonusMechanism,
		IsAvailable:          p.AvailableOnline,
		IsOrderable:          p.IsOrderable,
		IsPreviouslyBought:   p.IsPreviouslyBought,
		UnitSize:             p.SalesUnitSize,
		UnitPriceDescription: p.UnitPriceDescription,
		PropertyIcons:        p.PropertyIcons,
	}
}

// SearchProducts searches for products by query string and returns up to limit results.
// If limit is 0 or negative, defaults to 30 products.
//
// Example:
//
//	products, err := client.SearchProducts(ctx, "melk", 10)
func (c *Client) SearchProducts(ctx context.Context, query string, limit int) ([]Product, error) {
	if limit <= 0 {
		limit = 30
	}

	params := url.Values{}
	params.Set("query", query)
	params.Set("page", "0")
	params.Set("size", strconv.Itoa(limit))
	params.Set("sortOn", "RELEVANCE")

	path := "/mobile-services/product/search/v2?" + params.Encode()

	var result searchResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("search products failed: %w", err)
	}

	products := make([]Product, 0, len(result.Products))
	for _, p := range result.Products {
		products = append(products, p.toProduct())
	}

	return products, nil
}

// GetProduct retrieves a single product by its webshopId.
// For nutritional information, use GetProductFull instead.
func (c *Client) GetProduct(ctx context.Context, productID int) (*Product, error) {
	path := fmt.Sprintf("/mobile-services/product/detail/v4/fir/%d", productID)

	var result struct {
		ProductID   int             `json:"productId"`
		ProductCard productResponse `json:"productCard"`
	}

	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get product failed: %w", err)
	}

	product := result.ProductCard.toProduct()
	return &product, nil
}

// GetProductFull retrieves a product with full details including nutritional info.
// This makes an additional GraphQL call for nutritional data.
func (c *Client) GetProductFull(ctx context.Context, productID int) (*Product, error) {
	product, err := c.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}

	nutritionInfo, err := c.fetchNutritionalInfo(ctx, productID)
	if err == nil && nutritionInfo != nil {
		product.NutritionalInfo = nutritionInfo
	}

	return product, nil
}

// fetchNutritionalInfo fetches nutritional data for a product via GraphQL.
func (c *Client) fetchNutritionalInfo(ctx context.Context, productID int) ([]NutritionalInfo, error) {
	variables := map[string]any{
		"productId": productID,
	}

	var resp productNutritionResponse
	if err := c.doGraphQL(ctx, fetchProductNutritionQuery, variables, &resp); err != nil {
		return nil, err
	}

	if resp.Product.TradeItem == nil || len(resp.Product.TradeItem.Nutritions) == 0 {
		return nil, nil
	}

	var nutritionalInfo []NutritionalInfo
	for _, nutrition := range resp.Product.TradeItem.Nutritions {
		for _, n := range nutrition.Nutrients {
			nutritionalInfo = append(nutritionalInfo, NutritionalInfo{
				Amount: n.Value, // Contains value with unit, e.g., "15.5 g"
				Unit:   n.Name,  // Display name, e.g., "Fat"
				Type:   NutrientType(n.Type),
			})
		}
	}

	return nutritionalInfo, nil
}

// GetProductsByIDs retrieves multiple products by their webshopIds in a single request.
// Products are returned in the same order as the input IDs.
func (c *Client) GetProductsByIDs(ctx context.Context, productIDs []int) ([]Product, error) {
	if len(productIDs) == 0 {
		return nil, nil
	}

	params := url.Values{}
	for _, id := range productIDs {
		params.Add("ids", strconv.Itoa(id))
	}
	params.Set("sortOn", "INPUT_PRODUCT_IDS")

	path := "/mobile-services/product/search/v2/products?" + params.Encode()

	var result struct {
		Products []productResponse `json:"products"`
	}
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get products by ids failed: %w", err)
	}

	products := make([]Product, 0, len(result.Products))
	for _, p := range result.Products {
		products = append(products, p.toProduct())
	}

	return products, nil
}

// bonusSectionResponse matches the API response for bonus section.
type bonusSectionResponse struct {
	SectionType          string `json:"sectionType"`
	SectionDescription   string `json:"sectionDescription"`
	BonusGroupOrProducts []struct {
		Product    *productResponse `json:"product,omitempty"`
		BonusGroup *struct {
			ID                  string            `json:"id"`
			DiscountDescription string            `json:"discountDescription"`
			Products            []productResponse `json:"products"`
		} `json:"bonusGroup,omitempty"`
	} `json:"bonusGroupOrProducts"`
}

// GetBonusProducts retrieves products currently on bonus (promotion).
// Filter by category (e.g., "Vlees", "Zuivel") or pass empty string for all.
// If limit is 0 or negative, defaults to 30 products.
func (c *Client) GetBonusProducts(ctx context.Context, category string, limit int) ([]Product, error) {
	if limit <= 0 {
		limit = 30
	}

	params := url.Values{}
	params.Set("application", "AHWEBSHOP")
	params.Set("date", time.Now().Format("2006-01-02"))
	params.Set("promotionType", "NATIONAL")
	if category != "" {
		params.Set("category", category)
	}

	path := "/mobile-services/bonuspage/v2/section?" + params.Encode()

	var result bonusSectionResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get bonus products failed: %w", err)
	}

	var products []Product
	for _, item := range result.BonusGroupOrProducts {
		if item.Product != nil {
			products = append(products, item.Product.toProduct())
		}
		if item.BonusGroup != nil {
			for _, p := range item.BonusGroup.Products {
				products = append(products, p.toProduct())
			}
		}
		if limit > 0 && len(products) >= limit {
			break
		}
	}

	return products, nil
}

// GetSpotlightBonusProducts retrieves featured/highlighted bonus products.
// These are typically the best or most promoted deals of the week.
func (c *Client) GetSpotlightBonusProducts(ctx context.Context) ([]Product, error) {
	params := url.Values{}
	params.Set("application", "AHWEBSHOP")
	params.Set("date", time.Now().Format("2006-01-02"))

	path := "/mobile-services/bonuspage/v2/section/spotlight?" + params.Encode()

	var result bonusSectionResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get spotlight bonus products failed: %w", err)
	}

	var products []Product
	for _, item := range result.BonusGroupOrProducts {
		if item.Product != nil {
			products = append(products, item.Product.toProduct())
		}
		if item.BonusGroup != nil {
			for _, p := range item.BonusGroup.Products {
				products = append(products, p.toProduct())
			}
		}
	}

	return products, nil
}
