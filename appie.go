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
// For full access (orders, shopping lists, member profile), use Login which
// handles the full browser-based OAuth flow automatically:
//
//	client := appie.New(appie.WithConfigPath(".appie.json"))
//	if err := client.Login(ctx); err != nil {
//		log.Fatal(err)
//	}
//	// Tokens are auto-saved when configPath is set
//
// # Configuration
//
// Use NewWithConfig to load persisted tokens:
//
//	client, err := appie.NewWithConfig(".appie.json")
//	// Expired tokens are refreshed automatically
package appie
