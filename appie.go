// Package appie provides a Go client for the Albert Heijn mobile API.
//
// The client supports both anonymous and authenticated access to the AH API,
// allowing you to search products, manage shopping lists, and place orders.
//
// # Quick Start
//
// Create a client and get an anonymous token to browse products:
//
//	client := appie.New()
//	ctx := context.Background()
//
//	if err := client.GetAnonymousToken(ctx); err != nil {
//		log.Fatal(err)
//	}
//
//	products, err := client.SearchProducts(ctx, "melk", 10)
//
// # Authentication
//
// For full access (orders, shopping lists, member profile), authenticate via
// the browser-based OAuth flow:
//
//	// 1. Open LoginURL() in browser
//	// 2. User logs in and gets redirected to appie://login-exit?code=...
//	// 3. Exchange code for tokens
//	client.ExchangeCode(ctx, code)
//
// # Configuration
//
// Use NewWithConfig to persist tokens across sessions:
//
//	client, err := appie.NewWithConfig(".appie.json")
//	// ... use client ...
//	client.SaveConfig() // Save updated tokens
package appie
