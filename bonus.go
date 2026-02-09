package appie

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// bonusMetadataResponse matches the API response for bonus metadata.
type bonusMetadataResponse struct {
	Periods []struct {
		BonusStartDate string `json:"bonusStartDate"`
		BonusEndDate   string `json:"bonusEndDate"`
		Tabs           []struct {
			Description     string `json:"description"`
			URLMetadataList []struct {
				URL         string `json:"url"`
				Count       int    `json:"count"`
				BonusType   string `json:"bonusType"`
				Description string `json:"description"`
			} `json:"urlMetadataList"`
		} `json:"tabs"`
	} `json:"periods"`
}

// bonusSectionResponse matches the API response for bonus section.
type bonusSectionResponse struct {
	SectionType          string `json:"sectionType"`
	SectionDescription   string `json:"sectionDescription"`
	BonusGroupOrProducts []struct {
		Product    *productResponse    `json:"product,omitempty"`
		BonusGroup *bonusGroupResponse `json:"bonusGroup,omitempty"`
	} `json:"bonusGroupOrProducts"`
}

type bonusGroupResponse struct {
	ID                  string            `json:"id"`
	SegmentDescription  string            `json:"segmentDescription"`
	DiscountDescription string            `json:"discountDescription"`
	Category            string            `json:"category"`
	Images              []imageResponse   `json:"images"`
	Products            []productResponse `json:"products"`
	ExampleFromPrice    float64           `json:"exampleFromPrice"`
	ExampleForPrice     float64           `json:"exampleForPrice"`
}

func (bg *bonusGroupResponse) toProduct() Product {
	var images []Image
	for _, img := range bg.Images {
		images = append(images, Image{
			URL:    img.URL,
			Width:  img.Width,
			Height: img.Height,
		})
	}

	return Product{
		Title:          bg.SegmentDescription,
		Category:       bg.Category,
		BonusMechanism: bg.DiscountDescription,
		IsBonus:        true,
		Price: Price{
			Now: bg.ExampleForPrice,
			Was: bg.ExampleFromPrice,
		},
		Images: images,
	}
}

// getBonusMetadata retrieves the list of valid bonus category names from the current period.
func (c *Client) getBonusMetadata(ctx context.Context) ([]string, error) {
	path := "/mobile-services/bonuspage/v3/metadata"

	var result bonusMetadataResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get bonus metadata failed: %w", err)
	}

	var categories []string
	for _, period := range result.Periods {
		for _, tab := range period.Tabs {
			for _, meta := range tab.URLMetadataList {
				if meta.BonusType == "NATIONAL" {
					categories = append(categories, meta.Description)
				}
			}
		}
	}
	return categories, nil
}

// getBonusSection retrieves bonus products for a single category.
func (c *Client) getBonusSection(ctx context.Context, category string) ([]Product, error) {
	params := url.Values{}
	params.Set("application", "AHWEBSHOP")
	params.Set("date", time.Now().Format("2006-01-02"))
	params.Set("promotionType", "NATIONAL")
	params.Set("category", category)

	path := "/mobile-services/bonuspage/v2/section?" + params.Encode()

	var result bonusSectionResponse
	if err := c.doRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get bonus products failed (category=%s): %w", category, err)
	}

	return collectBonusProducts(result), nil
}

// GetBonusProducts retrieves all products currently on bonus (promotion)
// across all categories. Results are deduplicated by product title (since
// group-level bonus entries have no product ID).
func (c *Client) GetBonusProducts(ctx context.Context) ([]Product, error) {
	categories, err := c.getBonusMetadata(ctx)
	if err != nil {
		return nil, err
	}

	seen := make(map[string]bool)
	var products []Product

	for _, category := range categories {
		catProducts, err := c.getBonusSection(ctx, category)
		if err != nil {
			return nil, err
		}

		for _, p := range catProducts {
			key := p.Title
			if !seen[key] {
				seen[key] = true
				products = append(products, p)
			}
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

	return collectBonusProducts(result), nil
}

// collectBonusProducts extracts products from a bonus section response.
// Individual products are returned directly. For bonus groups, if the group
// contains individual products those are returned; otherwise the group itself
// is converted to a Product entry.
func collectBonusProducts(result bonusSectionResponse) []Product {
	var products []Product
	for _, item := range result.BonusGroupOrProducts {
		if item.Product != nil {
			products = append(products, item.Product.toProduct())
		}
		if item.BonusGroup != nil {
			if len(item.BonusGroup.Products) > 0 {
				for _, p := range item.BonusGroup.Products {
					products = append(products, p.toProduct())
				}
			} else {
				products = append(products, item.BonusGroup.toProduct())
			}
		}
	}
	return products
}
