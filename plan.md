# Appie Library Plan

Go library for interacting with Albert Heijn's mobile API.

## API Endpoints

### Authentication

| Method | Endpoint                                                                                   | Description                             |
| ------ | ------------------------------------------------------------------------------------------ | --------------------------------------- |
| POST   | `api.ah.nl/mobile-auth/v1/auth/token/anonymous`                                            | Get anonymous token (no login required) |
| POST   | `api.ah.nl/mobile-auth/v1/auth/token`                                                      | Exchange auth code for tokens           |
| POST   | `api.ah.nl/mobile-auth/v1/auth/token/refresh`                                              | Refresh access token                    |
| GET    | `login.ah.nl/login?client_id=appie-ios&response_type=code&redirect_uri=appie://login-exit` | OAuth login page (requires captcha)     |

### Products

| Method | Endpoint                                           | Description                   |
| ------ | -------------------------------------------------- | ----------------------------- |
| GET    | `/mobile-services/product/search/v2`               | Search products               |
| GET    | `/mobile-services/product/detail/v4/fir/{id}`      | Get product details           |
| GET    | `/mobile-services/product/search/v2/products`      | Get products by IDs           |
| GET    | `/mobile-services/bonuspage/v2/section`            | Get bonus products by category|
| GET    | `/mobile-services/bonuspage/v2/section/spotlight`  | Get featured bonus products   |

### Order Management

| Method | Endpoint                                               | Description               |
| ------ | ------------------------------------------------------ | ------------------------- |
| GET    | `/mobile-services/order/v1/summaries/active`           | Get active order          |
| PUT    | `/mobile-services/order/v1/items`                      | Add/update items in order |

### Shopping List

| Method | Endpoint                                | Description                    |
| ------ | --------------------------------------- | ------------------------------ |
| GET    | `/mobile-services/lists/v3/lists`       | Get lists (requires productId) |

---

## Required Headers

```go
User-Agent:        Appie/9.28 (iPhone17,3; iPhone; CPU OS 26_1 like Mac OS X)
x-clientname:      ipad
x-clientversion:   9.28
x-application:     AHWEBSHOP
x-accept-language: nl-NL
x-fraud-detection-installation-id: <uuid>
x-correlation-id:  <uuid>
Accept:            application/json
Content-Type:      application/json
Authorization:     Bearer <access_token>  // when authenticated
```

---

## Library Structure

```
appie/
├── appie.go            # Package declaration
├── types.go            # All shared types/structs
├── client.go           # Main client, HTTP setup, config
├── auth.go             # Login, token refresh, logout
├── products.go         # Product search and details (REST)
├── order.go            # Order management
├── shoppinglist.go     # Shopping list operations
├── member.go           # Member profile, bonus card
├── appie_test.go       # Integration tests
│
├── cmd/
│   └── login/
│       └── main.go     # CLI tool for browser-based login
│
└── doc/
    └── albertheijn_api.md  # API documentation
```

---

## Public API

```go
// client.go
func New(opts ...Option) *Client
func NewWithConfig(configPath string) (*Client, error)
func (c *Client) LoadConfig(path string) error
func (c *Client) SaveConfig() error
func (c *Client) IsAuthenticated() bool

// auth.go
func (c *Client) LoginURL() string
func (c *Client) ExchangeCode(ctx context.Context, code string) error
func (c *Client) RefreshToken(ctx context.Context) error
func (c *Client) GetAnonymousToken(ctx context.Context) error

// products.go
func (c *Client) SearchProducts(ctx context.Context, query string, limit int) ([]Product, error)
func (c *Client) GetProduct(ctx context.Context, productID int) (*Product, error)
func (c *Client) GetProductsByIDs(ctx context.Context, productIDs []int) ([]Product, error)
func (c *Client) GetBonusProducts(ctx context.Context, category string, limit int) ([]Product, error)
func (c *Client) GetSpotlightBonusProducts(ctx context.Context) ([]Product, error)

// order.go
func (c *Client) GetOrder(ctx context.Context) (*Order, error)
func (c *Client) AddToOrder(ctx context.Context, items []OrderItem) error
func (c *Client) RemoveFromOrder(ctx context.Context, productID int) error
func (c *Client) UpdateOrderItem(ctx context.Context, productID, quantity int) error
func (c *Client) ClearOrder(ctx context.Context) error
func (c *Client) GetOrderSummary(ctx context.Context) (*OrderSummary, error)

// shoppinglist.go
func (c *Client) GetShoppingLists(ctx context.Context) ([]ShoppingList, error)
func (c *Client) GetShoppingList(ctx context.Context) (*ShoppingList, error)
func (c *Client) AddToShoppingList(ctx context.Context, items []ListItem) error
func (c *Client) RemoveFromShoppingList(ctx context.Context, itemID string) error
func (c *Client) ShoppingListToOrder(ctx context.Context) error

// member.go
func (c *Client) GetMember(ctx context.Context) (*Member, error)
func (c *Client) GetBonusCard(ctx context.Context) (*BonusCard, error)
```

---

## Usage Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    appie "github.com/gwillem/appie-go"
)

func main() {
    // Load client from config file
    client, err := appie.NewWithConfig(".appie.json")
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Refresh token if needed
    if err := client.RefreshToken(ctx); err != nil {
        log.Fatal(err)
    }

    // Search for products
    products, err := client.SearchProducts(ctx, "melk", 10)
    if err != nil {
        log.Fatal(err)
    }

    for _, p := range products {
        fmt.Printf("%s - €%.2f\n", p.Title, p.Price.Now)
    }

    // Add to order
    err = client.AddToOrder(ctx, []appie.OrderItem{
        {ProductID: products[0].ID, Quantity: 1},
    })
    if err != nil {
        log.Fatal(err)
    }

    // Save updated tokens
    client.SaveConfig()
}
```

---

## Implementation Status

### Phase 1: Core Client & Auth ✅

- [x] Anonymous token
- [x] Browser-based login flow (CLI tool)
- [x] Token refresh
- [x] Config file storage

### Phase 2: Products ✅

- [x] Product search (REST API)
- [x] Product details
- [x] Products by IDs
- [x] Bonus products by category
- [x] Spotlight bonus products

### Phase 3: Order Management ✅

- [x] Get active order
- [x] Add/update items
- [x] Remove items
- [x] Get order summary

### Phase 4: Shopping List ✅

- [x] Get shopping lists
- [ ] Add/remove items (endpoint needs verification)

### Phase 5: Member ✅

- [x] Get member (GraphQL FetchMember)
- [x] Get member full profile
- [x] Get bonus card

### Phase 6: Polish

- [x] Error handling (basic)
- [ ] Rate limiting
- [ ] Retry logic
- [ ] Documentation
- [ ] Examples

---

## Files Created

| File | Purpose | Status |
|------|---------|--------|
| `appie.go` | Package declaration | ✅ |
| `types.go` | All shared types/structs | ✅ |
| `client.go` | Main client, HTTP setup, config | ✅ |
| `auth.go` | Login, token refresh, logout | ✅ |
| `products.go` | Product search and details | ✅ |
| `order.go` | Order management | ✅ |
| `shoppinglist.go` | Shopping list | ✅ |
| `member.go` | Member profile (GraphQL) | ✅ |
| `appie_test.go` | Integration tests | ✅ |
| `cmd/login/main.go` | CLI login tool | ✅ |
| `doc/albertheijn_api.md` | API documentation | ✅ |

---

## Tested Functions

| Function | Status | Notes |
|----------|--------|-------|
| `New()` | ✅ | |
| `NewWithConfig()` | ✅ | |
| `LoginURL()` | ✅ | |
| `GetAnonymousToken()` | ✅ | |
| `RefreshToken()` | ✅ | |
| `SearchProducts()` | ✅ | |
| `GetProduct()` | ✅ | |
| `GetBonusProducts()` | ✅ | Returns 0 products when category has no bonus |
| `GetOrder()` | ✅ | |
| `GetShoppingLists()` | ✅ | API requires productId param (pass 0 for default) |
| `GetMember()` | ✅ | Uses GraphQL FetchMember |
| `GetMemberFull()` | ✅ | Full profile with address, cards, audiences |
| `GetBonusCard()` | ✅ | Returns bonus card number from member profile |

---

## Next Steps

1. **Shopping List Mutations**: Discover endpoints for adding/removing items from lists.

2. **Add More Tests**: Add tests for AddToOrder, RemoveFromOrder, etc.

3. **Examples**: Create example scripts for common use cases.

4. **Documentation**: Add godoc comments and README.
