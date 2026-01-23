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
  "member_id": "12345678"
}
```

## Command-Line Tools

The repository includes several CLI tools in `cmd/`:

```bash
# Login and save credentials
go run ./cmd/login

# Dump feature flags
go run ./cmd/dump-feature-flags

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
