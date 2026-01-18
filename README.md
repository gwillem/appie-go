# appie-go

Go client library for the Albert Heijn mobile API.

## Installation

```bash
go get github.com/gwillem/appie-go
```

## Quick Start

### Anonymous Access (No Login)

Browse products without authentication:

```go
package main

import (
    "context"
    "fmt"
    "log"

    appie "github.com/gwillem/appie-go"
)

func main() {
    client := appie.New()
    ctx := context.Background()

    // Get anonymous token for product browsing
    if err := client.GetAnonymousToken(ctx); err != nil {
        log.Fatal(err)
    }

    // Search for products
    products, err := client.SearchProducts(ctx, "hagelslag", 5)
    if err != nil {
        log.Fatal(err)
    }

    for _, p := range products {
        fmt.Printf("%s - €%.2f\n", p.Title, p.Price.Now)
    }
}
```

### Authenticated Access

For orders, shopping lists, and member data, you need to authenticate:

```go
// 1. Get login URL and open in browser
client := appie.New()
fmt.Println("Open this URL:", client.LoginURL())

// 2. After login, browser redirects to: appie://login-exit?code=AUTH_CODE
// Extract the code and exchange it:
if err := client.ExchangeCode(ctx, "AUTH_CODE"); err != nil {
    log.Fatal(err)
}

// 3. Save tokens for future use
client = appie.New(
    appie.WithConfigPath(".appie.json"),
    appie.WithTokens(client.AccessToken(), client.RefreshTokenValue()),
)
client.SaveConfig()
```

Or load from saved config:

```go
client, err := appie.NewWithConfig(".appie.json")
if err != nil {
    log.Fatal(err)
}

if !client.IsAuthenticated() {
    // Need to login first
}
```

## API Reference

### Products

```go
// Search products
products, _ := client.SearchProducts(ctx, "melk", 20)

// Get single product by ID
product, _ := client.GetProduct(ctx, 123456)

// Get multiple products by ID
products, _ := client.GetProductsByIDs(ctx, []int{123, 456, 789})

// Get bonus (promotional) products
bonusProducts, _ := client.GetBonusProducts(ctx, "", 30)        // All categories
bonusProducts, _ := client.GetBonusProducts(ctx, "Zuivel", 30)  // Specific category

// Get featured deals
spotlight, _ := client.GetSpotlightBonusProducts(ctx)
```

### Orders (Cart)

```go
// Get current order
order, _ := client.GetOrder(ctx)
fmt.Printf("Total: €%.2f (%d items)\n", order.TotalPrice, order.TotalCount)

// Add items to order
client.AddToOrder(ctx, []appie.OrderItem{
    {ProductID: 123456, Quantity: 2},
    {ProductID: 789012, Quantity: 1},
})

// Update quantity
client.UpdateOrderItem(ctx, 123456, 3)

// Remove item (set quantity to 0)
client.RemoveFromOrder(ctx, 123456)

// Get order summary
summary, _ := client.GetOrderSummary(ctx)
```

### Shopping Lists

```go
// Get all shopping lists
lists, _ := client.GetShoppingLists(ctx, 0)

// Get default shopping list
list, _ := client.GetShoppingList(ctx)

// Add product to list
client.AddProductToShoppingList(ctx, 123456, 2)

// Add free-text item
client.AddFreeTextToShoppingList(ctx, "brood", 1)

// Add multiple items
client.AddToShoppingList(ctx, []appie.ListItem{
    {ProductID: 123, Quantity: 1},
    {Name: "eieren", Quantity: 12},
})

// Check/uncheck item
client.CheckShoppingListItem(ctx, "item-uuid", true)

// Remove item
client.RemoveFromShoppingList(ctx, "item-uuid")

// Convert shopping list to order
client.ShoppingListToOrder(ctx)
```

### Member Profile

```go
// Basic info
member, _ := client.GetMember(ctx)
fmt.Printf("%s %s\n", member.FirstName, member.LastName)

// Full profile with address and cards
full, _ := client.GetMemberFull(ctx)
fmt.Printf("Bonus card: %s\n", full.BonusCardNumber)
fmt.Printf("Audiences: %v\n", full.Audiences)

// Just the bonus card
card, _ := client.GetBonusCard(ctx)
```

### Feature Flags

```go
flags, _ := client.GetFeatureFlags(ctx)

if flags.IsEnabled("dark-mode") {
    // Feature is enabled
}

// Get rollout percentage (0-100)
pct := flags.Rollout("new-checkout")
```

### Token Management

```go
// Refresh access token
client.RefreshToken(ctx)

// Check authentication status
if client.IsAuthenticated() {
    // Has valid token
}

// Get current tokens
accessToken := client.AccessToken()
refreshToken := client.RefreshTokenValue()

// Clear tokens
client.Logout()
```

## Configuration

The client can be configured with options:

```go
client := appie.New(
    appie.WithHTTPClient(customHTTPClient),
    appie.WithBaseURL("https://api.ah.nl"),
    appie.WithTokens(accessToken, refreshToken),
    appie.WithConfigPath(".appie.json"),
)
```

### Config File Format

```json
{
  "access_token": "eyJ...",
  "refresh_token": "eyJ...",
  "member_id": "12345678"
}
```

## Types

### Product

```go
type Product struct {
    ID                   int
    Title                string
    Brand                string
    Category             string
    Price                Price
    Images               []Image
    NutriScore           string    // A, B, C, D, E
    IsBonus              bool      // Currently on promotion
    IsAvailable          bool      // Available for online order
    UnitSize             string    // e.g., "500 g"
    UnitPriceDescription string    // e.g., "per kg €5.99"
}

type Price struct {
    Now float64  // Current price
    Was float64  // Price before discount (0 if no discount)
}
```

### Order

```go
type Order struct {
    ID         string
    State      string       // NEW, REOPENED, PROCESSING, DELIVERED
    Items      []OrderItem
    TotalCount int
    TotalPrice float64
}

type OrderItem struct {
    ProductID int
    Quantity  int
    Product   *Product  // Populated when retrieved with order
}
```

### ShoppingList

```go
type ShoppingList struct {
    ID        string
    Name      string
    ItemCount int
    Items     []ListItem
}

type ListItem struct {
    ID        string
    Name      string     // For free-text items
    ProductID int        // For product items (0 for free-text)
    Quantity  int
    Checked   bool
    Product   *Product
}
```

## Command-Line Tools

The repository includes several CLI tools in `cmd/`:

```bash
# Login and save credentials
go run ./cmd/login

# Dump feature flags
go run ./cmd/dump_feature_flags

# Dump member profile data
go run ./cmd/dump-member

# Dump GraphQL schema
go run ./cmd/dump-graphql
```

## Notes

- Tokens expire after ~7 days. Use `RefreshToken()` to get new tokens.
- The API uses both REST and GraphQL endpoints internally.
- Anonymous tokens work for product browsing but not for orders or member data.
- Rate limiting may apply; implement backoff for production use.

## License

AGPLv3
