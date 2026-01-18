package appie

import "time"

// Token represents the authentication tokens returned by the API.
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	MemberID     string `json:"member_id,omitempty"`
	ExpiresIn    int    `json:"expires_in"`
}

// Config holds the client configuration stored in .appie.json.
type Config struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	MemberID     string `json:"member_id,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
}

// Product represents an AH product.
type Product struct {
	ID              int      `json:"id"`
	WebshopID       string   `json:"webshopId,omitempty"`
	Title           string   `json:"title"`
	Brand           string   `json:"brand,omitempty"`
	Category        string   `json:"category,omitempty"`
	ShortDescription string  `json:"shortDescription,omitempty"`
	Price           Price    `json:"price"`
	Images          []Image  `json:"images,omitempty"`
	NutriScore      string   `json:"nutriScore,omitempty"`
	IsBonus         bool     `json:"isBonus"`
	IsAvailable     bool     `json:"isAvailable"`
	UnitSize        string   `json:"unitSize,omitempty"`
	UnitPriceDescription string `json:"unitPriceDescription,omitempty"`
}

// Price represents product pricing.
type Price struct {
	Now      float64 `json:"now"`
	Was      float64 `json:"was,omitempty"`
	UnitSize string  `json:"unitSize,omitempty"`
}

// Image represents a product image.
type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

// Order represents a shopping order.
type Order struct {
	ID            string      `json:"id"`
	Hash          string      `json:"hash,omitempty"`
	State         string      `json:"state,omitempty"`
	Items         []OrderItem `json:"items"`
	TotalCount    int         `json:"totalCount"`
	TotalPrice    float64     `json:"totalPrice"`
	LastUpdated   time.Time   `json:"lastUpdated,omitempty"`
}

// OrderItem represents an item in an order.
type OrderItem struct {
	ProductID int     `json:"productId"`
	Quantity  int     `json:"quantity"`
	Product   *Product `json:"product,omitempty"`
}

// OrderSummary represents the order summary/totals.
type OrderSummary struct {
	TotalItems    int     `json:"totalItems"`
	TotalPrice    float64 `json:"totalPrice"`
	TotalDiscount float64 `json:"totalDiscount,omitempty"`
	DeliveryCost  float64 `json:"deliveryCost,omitempty"`
}

// ShoppingList represents a shopping list.
type ShoppingList struct {
	ID        string     `json:"id"`
	Name      string     `json:"name,omitempty"`
	ItemCount int        `json:"itemCount,omitempty"`
	Items     []ListItem `json:"items,omitempty"`
}

// ListItem represents an item in a shopping list.
type ListItem struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	ProductID int      `json:"productId,omitempty"`
	Quantity  int      `json:"quantity"`
	Checked   bool     `json:"checked"`
	Product   *Product `json:"product,omitempty"`
}

// Member represents a member profile.
type Member struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// Address represents a physical address.
type Address struct {
	Street           string `json:"street"`
	HouseNumber      int    `json:"houseNumber"`
	HouseNumberExtra string `json:"houseNumberExtra,omitempty"`
	PostalCode       string `json:"postalCode"`
	City             string `json:"city"`
	CountryCode      string `json:"countryCode,omitempty"`
}

// MemberFull represents the full member profile with all details.
type MemberFull struct {
	ID              string   `json:"id"`
	FirstName       string   `json:"firstName"`
	LastName        string   `json:"lastName"`
	Email           string   `json:"email"`
	Gender          string   `json:"gender,omitempty"`
	DateOfBirth     string   `json:"dateOfBirth,omitempty"`
	PhoneNumber     string   `json:"phoneNumber,omitempty"`
	Address         Address  `json:"address"`
	BonusCardNumber string   `json:"bonusCardNumber,omitempty"`
	GallCardNumber  string   `json:"gallCardNumber,omitempty"`
	Audiences       []string `json:"audiences,omitempty"`
}

// BonusCard represents the bonus card info.
type BonusCard struct {
	CardNumber string `json:"cardNumber"`
	IsActive   bool   `json:"isActive"`
}

// SearchResult represents the result of a product search.
type SearchResult struct {
	Products   []Product `json:"products"`
	TotalCount int       `json:"totalCount"`
	Page       int       `json:"page"`
	PageSize   int       `json:"pageSize"`
}

// GraphQL request/response types

type graphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

type graphQLResponse[T any] struct {
	Data   T              `json:"data"`
	Errors []graphQLError `json:"errors,omitempty"`
}

type graphQLError struct {
	Message string `json:"message"`
	Path    []any  `json:"path,omitempty"`
}

// API error response
type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *apiError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return e.Code
}
