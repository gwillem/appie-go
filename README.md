# appie-go

<img src="doc/appie.png" align="right" alt="appie gopher">

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

For orders, shopping lists, and member data, use `Login()` which handles the
full browser-based OAuth flow automatically:

```go
client, err := appie.NewWithConfig(".appie.json")

ctx := context.Background()

if err := client.Login(ctx); err != nil {
    log.Fatal(err)
}
// Tokens are auto-saved when configPath is set
```

Expired access tokens are automatically refreshed using the stored refresh token.

### Receipts (Kassabonnen)

Retrieve in-store purchase receipts:

```go
// Get all receipts
receipts, err := client.GetReceipts(ctx)
for _, r := range receipts {
    fmt.Printf("%s: €%.2f\n", r.Date, r.TotalAmount)
}

// Get receipt details with items, discounts, and payments
receipt, err := client.GetReceipt(ctx, receipts[0].TransactionID)
for _, item := range receipt.Items {
    fmt.Printf("  %s x%d - €%.2f\n", item.Description, item.Quantity, item.Amount)
}
```

## API Reference

See [pkg.go.dev/github.com/gwillem/appie-go](https://pkg.go.dev/github.com/gwillem/appie-go) for full documentation.

## CLI

Install the `appie` command:

```bash
go install github.com/gwillem/appie-go/cmd/appie@latest
```

Usage:

```bash
# Login to Albert Heijn (opens browser for OAuth)
appie login

# List recent receipts
appie receipt

# List last 5 receipts
appie receipt -n 5

# Show receipt details (items, discounts, payment)
appie receipt <transaction-id>
```

Config is stored at `~/.config/appie/config.json` (or `$XDG_CONFIG_HOME/appie/config.json`). Override with `-c`.

## Notes

- Rate limiting may apply. AH does not send back-off headers, so the practical rate limit is unknown.
- **Server-side state:** Albert Heijn maintains an "active order" on the server, which determines the delivery date context for bonus promo visibility. When you call `ReopenOrder`, that order becomes the active one. Always call `RevertOrder` when done to avoid the account being stuck in a future order with incorrect bonus promos.

## License

AGPLv3
