# appie-go

Go client library for the Albert Heijn mobile API.

[![Go Reference](https://pkg.go.dev/badge/github.com/gwillem/appie-go.svg)](https://pkg.go.dev/github.com/gwillem/appie-go)

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

For orders, shopping lists, and member data, you need to authenticate.
Use the login tool to obtain and save tokens:

```bash
go run ./cmd/login
```

Then load from saved config:

```go
client, err := appie.NewWithConfig(".appie.json")
if err != nil {
    log.Fatal(err)
}

if !client.IsAuthenticated() {
    // Need to login first
}
```

Expired access tokens are automatically refreshed using the stored refresh token.

### Receipts (Kassabonnen)

Retrieve in-store purchase receipts:

```go
// Get all receipts
receipts, err := client.GetReceipts(ctx)
for _, r := range receipts {
    fmt.Printf("%s: %s - €%.2f\n", r.Date, r.StoreName, r.TotalAmount)
}

// Get receipt details with items
receipt, err := client.GetReceipt(ctx, receipts[0].TransactionID)
for _, item := range receipt.Items {
    fmt.Printf("  %s x%d - €%.2f\n", item.Description, item.Quantity, item.Amount)
}
```

## API Reference

See [pkg.go.dev/github.com/gwillem/appie-go](https://pkg.go.dev/github.com/gwillem/appie-go) for full documentation.

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
  "expires_at": "2025-02-09T12:00:00Z"
}
```

## Command-Line Tools

The repository includes several CLI tools in `cmd/`:

```bash
# Login and save credentials
go run ./cmd/login

# Dump member profile data
go run ./cmd/dump-member

# Dump GraphQL schema
go run ./cmd/dump-graphql
```

## Notes

- Tokens expire after ~7 days. Expired tokens are automatically refreshed when using `NewWithConfig()`.
- The API uses both REST and GraphQL endpoints internally.
- Anonymous tokens work for product browsing but not for orders or member data.
- Rate limiting may apply; implement backoff for production use.

## License

AGPLv3
