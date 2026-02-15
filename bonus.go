package appie

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
		BonusSegmentID: bg.ID,
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
	if err := c.DoRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
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
	if err := c.DoRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
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
			key := fmt.Sprintf("%d:%s", p.ID, p.Title)
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
	if err := c.DoRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("get spotlight bonus products failed: %w", err)
	}

	return collectBonusProducts(result), nil
}

// GraphQL query for fetching products within a bonus group/promotion segment.
const fetchBonusGroupProductsQuery = `query FetchBonusPromotionWithProducts(
  $id: String,
  $periodStart: String,
  $periodEnd: String,
  $filterUnavailableProducts: Boolean,
  $forcePromotionVisibility: Boolean = true,
  $showAllPromotionSegments: Boolean = true
) {
  bonusPromotions(
    input: {
      id: $id
      periodStart: $periodStart
      periodEnd: $periodEnd
      filterUnavailableProducts: $filterUnavailableProducts
      forcePromotionVisibility: $forcePromotionVisibility
      showAllPromotionSegments: $showAllPromotionSegments
    }
  ) {
    id
    title
    productCount
    products {
      id
      title
      brand
      category
      salesUnitSize
      icons
      availability { isOrderable }
      priceV2(
        periodStart: $periodStart
        periodEnd: $periodEnd
        filterUnavailableProducts: $filterUnavailableProducts
        forcePromotionVisibility: true
      ) {
        now { amount }
        was { amount }
        promotionLabel {
          tiers {
            mechanism
            description
          }
        }
      }
      imagePack {
        large { url width height }
      }
    }
  }
}`

type bonusGroupProductsResponse struct {
	BonusPromotions []struct {
		ID           string                `json:"id"`
		Title        string                `json:"title"`
		ProductCount int                   `json:"productCount"`
		Products     []bonusGraphQLProduct `json:"products"`
	} `json:"bonusPromotions"`
}

type bonusGraphQLProduct struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	Brand         string   `json:"brand"`
	Category      string   `json:"category"`
	SalesUnitSize string   `json:"salesUnitSize"`
	Icons         []string `json:"icons"`
	Availability  struct {
		IsOrderable bool `json:"isOrderable"`
	} `json:"availability"`
	PriceV2 struct {
		Now struct {
			Amount float64 `json:"amount"`
		} `json:"now"`
		Was struct {
			Amount float64 `json:"amount"`
		} `json:"was"`
		PromotionLabel *struct {
			Tiers []struct {
				Mechanism   string `json:"mechanism"`
				Description string `json:"description"`
			} `json:"tiers"`
		} `json:"promotionLabel"`
	} `json:"priceV2"`
	ImagePack []struct {
		Large *struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"large"`
	} `json:"imagePack"`
}

func (p *bonusGraphQLProduct) toProduct() Product {
	var images []Image
	for _, pack := range p.ImagePack {
		if pack.Large != nil {
			images = append(images, Image{
				URL:    pack.Large.URL,
				Width:  pack.Large.Width,
				Height: pack.Large.Height,
			})
		}
	}

	var nutriScore string
	for _, icon := range p.Icons {
		if score, ok := strings.CutPrefix(icon, "NUTRISCORE_"); ok {
			nutriScore = score
			break
		}
	}

	var bonusMechanism string
	if p.PriceV2.PromotionLabel != nil && len(p.PriceV2.PromotionLabel.Tiers) > 0 {
		bonusMechanism = p.PriceV2.PromotionLabel.Tiers[0].Description
	}

	return Product{
		ID:             p.ID,
		WebshopID:      strconv.Itoa(p.ID),
		Title:          p.Title,
		Brand:          p.Brand,
		Category:       p.Category,
		Price:          Price{Now: p.PriceV2.Now.Amount, Was: p.PriceV2.Was.Amount},
		Images:         images,
		NutriScore:     nutriScore,
		IsBonus:        true,
		BonusMechanism: bonusMechanism,
		IsOrderable:    p.Availability.IsOrderable,
		IsAvailable:    p.Availability.IsOrderable,
		UnitSize:       p.SalesUnitSize,
	}
}

// getBonusPeriod retrieves the current bonus period dates from metadata.
func (c *Client) getBonusPeriod(ctx context.Context) (startDate, endDate string, err error) {
	path := "/mobile-services/bonuspage/v3/metadata"
	var result bonusMetadataResponse
	if err := c.DoRequest(ctx, http.MethodGet, path, nil, &result); err != nil {
		return "", "", fmt.Errorf("get bonus period failed: %w", err)
	}
	if len(result.Periods) == 0 {
		return "", "", fmt.Errorf("no bonus periods available")
	}
	p := result.Periods[0]
	return p.BonusStartDate, p.BonusEndDate, nil
}

// GetBonusGroupProducts retrieves the individual products within a bonus
// promotion group. Use this to resolve group-level promotions (e.g.,
// "Alle Hak*") that appear in GetBonusProducts/GetSpotlightBonusProducts
// with ID==0 and a BonusSegmentID.
//
// The segmentID is the bonus group identifier, available as BonusSegmentID
// on Product entries returned by GetBonusProducts.
func (c *Client) GetBonusGroupProducts(ctx context.Context, segmentID string) ([]Product, error) {
	startDate, endDate, err := c.getBonusPeriod(ctx)
	if err != nil {
		return nil, err
	}

	variables := map[string]any{
		"id":                        segmentID,
		"periodStart":               startDate,
		"periodEnd":                 endDate,
		"filterUnavailableProducts": true,
		"forcePromotionVisibility":  true,
		"showAllPromotionSegments":  true,
	}

	var resp bonusGroupProductsResponse
	if err := c.DoGraphQL(ctx, fetchBonusGroupProductsQuery, variables, &resp); err != nil {
		return nil, fmt.Errorf("get bonus group products failed: %w", err)
	}

	if len(resp.BonusPromotions) == 0 {
		return nil, nil
	}

	promo := resp.BonusPromotions[0]
	products := make([]Product, 0, len(promo.Products))
	for _, p := range promo.Products {
		products = append(products, p.toProduct())
	}

	return products, nil
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
