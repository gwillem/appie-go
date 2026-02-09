# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

Go client library for the Albert Heijn (AH) mobile API (`github.com/gwillem/appie-go`). Wraps both REST (`/mobile-services/...`) and GraphQL (`/graphql`) endpoints, spoofing the official Appie iOS app user-agent.

## Commands

```bash
make test                # Unit tests (5s timeout, no network)
make integration-test    # Integration tests (requires .appie.json with valid tokens)

# Run a single test
go test -run TestSearchProducts -v ./...
go test -tags integration -run TestGetBonusProductsBatch -v -count=1 ./...
```

**Never use `go build` to verify compilation** — always use `go test ./...` instead.

## Test conventions

- Tests hitting AH servers go in `*_integration_test.go` with `//go:build integration` build tag.
- Unit tests (no network, use `httptest`) go in regular `*_test.go` files.
- Integration tests use `testClient(t)` helper (defined in `appie_integration_test.go`) which loads `.appie.json` and skips if unauthenticated.

## Architecture

**Client** (`client.go`): Thread-safe HTTP client with mutex-protected token state. Functional options pattern (`WithTokens`, `WithBaseURL`, etc.). Auto-refreshes expired tokens before requests via `ensureFreshToken()`. Core request methods: `doRequest()` (REST) and `doGraphQL()` (GraphQL).

**Feature modules** — each file is a self-contained domain:
- `auth.go` — OAuth login flow, token exchange/refresh, anonymous tokens
- `products.go` — Search, detail, batch fetch, nutritional info (via GraphQL)
- `bonus.go` — Bonus/promotion products, metadata (category discovery), spotlight
- `order.go` — Cart management, fulfillments (scheduled deliveries via GraphQL)
- `shoppinglist.go` — Shopping/favorite lists, item CRUD, list-to-order conversion
- `member.go` — Profile, bonus card (via GraphQL)
- `receipts.go` — In-store receipt history

**Types** (`types.go`): All public types. Internal API response types (e.g., `productResponse`) live in their respective feature files with `toProduct()`-style converters to public types.

**CLI tools** (`cmd/`): `login` for OAuth flow, `dump-member` and `dump-graphql` for exploration. Dev tools in `cmd/dev/` are gitignored.
