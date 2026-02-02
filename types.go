package appie

import "time"

// Token represents the authentication tokens returned by the API.
// These tokens are used for API authorization and should be stored securely.
type Token struct {
	// AccessToken is the bearer token used for API requests.
	AccessToken string `json:"access_token"`
	// RefreshToken is used to obtain a new access token when it expires.
	RefreshToken string `json:"refresh_token"`
	// MemberID is the authenticated user's member ID.
	MemberID string `json:"member_id,omitempty"`
	// ExpiresIn is the token lifetime in seconds (typically ~7 days).
	ExpiresIn int `json:"expires_in"`
}

// Config holds the client configuration stored in .appie.json.
// This is used for persisting authentication state across sessions.
type Config struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	MemberID     string    `json:"member_id,omitempty"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
}

// NutrientType represents the type of nutrient in nutritional information.
type NutrientType string

// NutritionalInfo represents nutritional information for a product.
type NutritionalInfo struct {
	// Amount is the quantity of this nutrient (e.g., "15", "2.5").
	Amount string `json:"amount,omitempty"`
	// Unit is the measurement unit (e.g., "g", "kcal", "kJ").
	Unit string `json:"unit,omitempty"`
	// Type identifies the nutrient (e.g., "ENERGY", "PROTEIN", "FAT").
	Type NutrientType `json:"type,omitempty"`
}

// Product represents an AH product with pricing and availability information.
type Product struct {
	// ID is the webshop product ID used for ordering.
	ID int `json:"id"`
	// WebshopID is the string representation of the ID.
	WebshopID string `json:"webshopId,omitempty"`
	// Title is the product name.
	Title string `json:"title"`
	// Brand is the product brand (e.g., "AH Biologisch").
	Brand string `json:"brand,omitempty"`
	// Category is the main product category.
	Category string `json:"category,omitempty"`
	// SubCategory is the secondary product category.
	SubCategory string `json:"subCategory,omitempty"`
	// ShortDescription is a brief product description.
	ShortDescription string `json:"shortDescription,omitempty"`
	// Price contains current and previous pricing.
	Price Price `json:"price"`
	// Images contains product images in various sizes.
	Images []Image `json:"images,omitempty"`
	// NutriScore is the nutritional score (A-E).
	NutriScore string `json:"nutriScore,omitempty"`
	// NutritionalInfo contains detailed nutritional values for the product.
	NutritionalInfo []NutritionalInfo `json:"nutritionalInfo,omitempty"`
	// IsBonus indicates if the product is currently on promotion.
	IsBonus bool `json:"isBonus"`
	// BonusMechanism describes the type of bonus (e.g., "25% korting").
	BonusMechanism string `json:"bonusMechanism,omitempty"`
	// IsAvailable indicates if the product can be ordered online.
	IsAvailable bool `json:"isAvailable"`
	// IsOrderable indicates if the product can currently be ordered.
	IsOrderable bool `json:"isOrderable"`
	// IsPreviouslyBought indicates if the user has bought this before.
	IsPreviouslyBought bool `json:"isPreviouslyBought"`
	// UnitSize is the package size (e.g., "500 g", "1 L").
	UnitSize string `json:"unitSize,omitempty"`
	// UnitPriceDescription describes price per unit (e.g., "per kg €5.99").
	UnitPriceDescription string `json:"unitPriceDescription,omitempty"`
	// PropertyIcons contains icons like "vegan", "bio", etc.
	PropertyIcons []string `json:"propertyIcons,omitempty"`
}

// Price represents product pricing in EUR.
type Price struct {
	// Now is the current price.
	Now float64 `json:"now"`
	// Was is the price before any discount (0 if no discount).
	Was float64 `json:"was,omitempty"`
	// UnitSize for price comparison purposes.
	UnitSize string `json:"unitSize,omitempty"`
}

// Image represents a product image. Images are available in multiple sizes
// (typically 48, 80, 200, 400, 800 pixels).
type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

// Order represents a shopping order (cart). Orders in AH can have various states:
// NEW, REOPENED, PROCESSING, DELIVERED, etc.
type Order struct {
	// ID is the numeric order ID.
	ID string `json:"id"`
	// Hash is used for order verification in API calls.
	Hash string `json:"hash,omitempty"`
	// State is the current order state (e.g., "NEW", "REOPENED").
	State string `json:"state,omitempty"`
	// Items contains all products in the order.
	Items []OrderItem `json:"items"`
	// TotalCount is the number of unique items.
	TotalCount int `json:"totalCount"`
	// TotalPrice is the total order value in EUR.
	TotalPrice float64 `json:"totalPrice"`
	// LastUpdated is when the order was last modified.
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
}

// OrderItem represents a product and quantity in an order.
type OrderItem struct {
	// ProductID is the webshop product ID.
	ProductID int `json:"productId"`
	// Quantity is the number of items (0 to remove).
	Quantity int `json:"quantity"`
	// Product contains product details when retrieved with the order.
	Product *Product `json:"product,omitempty"`
}

// OrderSummary provides totals and discount information for an order.
type OrderSummary struct {
	TotalItems    int     `json:"totalItems"`
	TotalPrice    float64 `json:"totalPrice"`
	TotalDiscount float64 `json:"totalDiscount,omitempty"`
	DeliveryCost  float64 `json:"deliveryCost,omitempty"`
}

// ShoppingList represents a user's shopping list. Users can have multiple lists.
type ShoppingList struct {
	// ID is a UUID identifying the list.
	ID string `json:"id"`
	// Name is the user-defined list name.
	Name string `json:"name,omitempty"`
	// ItemCount is the number of items in the list.
	ItemCount int `json:"itemCount,omitempty"`
	// Items contains the list items (may not always be populated).
	Items []ListItem `json:"items,omitempty"`
}

// ListItem represents an item in a shopping list. Items can be either
// linked to a product (by ProductID) or free-text entries (by Name).
type ListItem struct {
	// ID is the unique item identifier within the list.
	ID string `json:"id"`
	// Name is the item description (for free-text items).
	Name string `json:"name"`
	// ProductID links to a product (0 for free-text items).
	ProductID int `json:"productId,omitempty"`
	// Quantity is the desired amount.
	Quantity int `json:"quantity"`
	// Checked indicates if the item has been picked/crossed off.
	Checked bool `json:"checked"`
	// Product contains product details when available.
	Product *Product `json:"product,omitempty"`
}

// Member represents basic member profile information.
type Member struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// Address represents a Dutch postal address.
type Address struct {
	Street           string `json:"street"`
	HouseNumber      int    `json:"houseNumber"`
	HouseNumberExtra string `json:"houseNumberExtra,omitempty"` // e.g., "A", "2-hoog"
	PostalCode       string `json:"postalCode"`                 // Dutch format: "1234AB"
	City             string `json:"city"`
	CountryCode      string `json:"countryCode,omitempty"` // e.g., "NL"
}

// MemberFull represents the complete member profile including address,
// loyalty cards, and customer segmentation data.
type MemberFull struct {
	ID              string   `json:"id"`
	FirstName       string   `json:"firstName"`
	LastName        string   `json:"lastName"`
	Email           string   `json:"email"`
	Gender          string   `json:"gender,omitempty"`      // "MALE", "FEMALE", or empty
	DateOfBirth     string   `json:"dateOfBirth,omitempty"` // Format: "YYYY-MM-DD"
	PhoneNumber     string   `json:"phoneNumber,omitempty"` // Dutch format with country code
	Address         Address  `json:"address"`
	BonusCardNumber string   `json:"bonusCardNumber,omitempty"` // 13-digit card number
	GallCardNumber  string   `json:"gallCardNumber,omitempty"`  // Gall & Gall card
	Audiences       []string `json:"audiences,omitempty"`       // Customer segments
}

// BonusCard represents the AH Bonuskaart (loyalty card) information.
type BonusCard struct {
	// CardNumber is the 13-digit bonus card number.
	CardNumber string `json:"cardNumber"`
	// IsActive indicates if a bonus card is linked to the account.
	IsActive bool `json:"isActive"`
}

// SearchResult represents the result of a product search with pagination.
type SearchResult struct {
	Products   []Product `json:"products"`
	TotalCount int       `json:"totalCount"`
	Page       int       `json:"page"`
	PageSize   int       `json:"pageSize"`
}

// Fulfillment represents a scheduled order with delivery information.
// A fulfillment is an order that has been submitted for delivery.
type Fulfillment struct {
	// OrderID is the numeric order identifier.
	OrderID int `json:"orderId"`
	// Status is the delivery status (e.g., "REOPENED", "SUBMITTED").
	Status string `json:"status"`
	// StatusDescription is a human-readable status.
	StatusDescription string `json:"statusDescription"`
	// ShoppingType is the order type (e.g., "DELIVERY", "PICKUP").
	ShoppingType string `json:"shoppingType"`
	// TotalPrice is the order total in EUR.
	TotalPrice float64 `json:"totalPrice"`
	// Delivery contains delivery slot and address information.
	Delivery FulfillmentDelivery `json:"delivery"`
}

// FulfillmentDelivery contains delivery details for a fulfillment.
type FulfillmentDelivery struct {
	// Status is the delivery status.
	Status string `json:"status"`
	// Method is the delivery method.
	Method string `json:"method"`
	// Slot contains the delivery time window.
	Slot DeliverySlot `json:"slot"`
	// Address is the delivery address.
	Address Address `json:"address"`
}

// DeliverySlot represents a delivery time window.
type DeliverySlot struct {
	// Date is the delivery date (YYYY-MM-DD).
	Date string `json:"date"`
	// DateDisplay is the formatted date (e.g., "dinsdag 20 januari").
	DateDisplay string `json:"dateDisplay"`
	// TimeDisplay is the formatted time window (e.g., "18:00 - 20:00").
	TimeDisplay string `json:"timeDisplay"`
	// StartTime is the slot start time (HH:MM).
	StartTime string `json:"startTime"`
	// EndTime is the slot end time (HH:MM).
	EndTime string `json:"endTime"`
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

// Receipt represents an in-store purchase receipt (kassabon).
type Receipt struct {
	// TransactionID uniquely identifies this receipt.
	TransactionID string `json:"transactionId"`
	// Date is the purchase date/time.
	Date string `json:"date"`
	// StoreName is the store where the purchase was made.
	StoreName string `json:"storeName"`
	// StoreID is the store identifier.
	StoreID int `json:"storeId,omitempty"`
	// TotalAmount is the total purchase amount in EUR.
	TotalAmount float64 `json:"totalAmount"`
	// Items contains the purchased products (only in detailed receipt).
	Items []ReceiptItem `json:"items,omitempty"`
}

// ReceiptItem represents a single item on a receipt.
type ReceiptItem struct {
	// Description is the product name/description.
	Description string `json:"description"`
	// Quantity is the number of items purchased.
	Quantity int `json:"quantity"`
	// Amount is the total price for this line item.
	Amount float64 `json:"amount"`
	// UnitPrice is the price per unit.
	UnitPrice float64 `json:"unitPrice,omitempty"`
	// ProductID is the webshop product ID if available.
	ProductID int `json:"productId,omitempty"`
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
