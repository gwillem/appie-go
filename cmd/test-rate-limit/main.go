package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	appie "github.com/gwillem/appie-go"
)

func main() {
	client, err := appie.NewWithConfig(".appie.json")
	if err != nil {
		log.Fatal(err)
	}

	if !client.IsAuthenticated() {
		log.Fatal("Not authenticated. Run login first.")
	}

	ctx := context.Background()

	endpoints := []struct {
		name string
		url  string
	}{
		{"Product Search", "https://api.ah.nl/mobile-services/product/search/v2?query=melk&size=1"},
		{"GraphQL", "https://api.ah.nl/graphql"},
		{"Feature Flags", "https://api.ah.nl/mobile-services/config/v1/features/ios?version=9.28"},
		{"Order Summary", "https://api.ah.nl/mobile-services/order/v1/summaries/active?sortBy=DEFAULT"},
	}

	for _, ep := range endpoints {
		fmt.Printf("\n=== %s ===\n", ep.name)
		fmt.Printf("URL: %s\n\n", ep.url)

		method := "GET"
		var body string
		if strings.Contains(ep.url, "graphql") {
			method = "POST"
			body = `{"query":"{ member { id } }"}`
		}

		headers := fetchHeaders(ctx, client, method, ep.url, body)
		printRateLimitHeaders(headers)
	}
}

func fetchHeaders(ctx context.Context, client *appie.Client, method, url, body string) http.Header {
	var bodyReader *strings.Reader
	if body != "" {
		bodyReader = strings.NewReader(body)
	}

	var req *http.Request
	var err error
	if bodyReader != nil {
		req, err = http.NewRequestWithContext(ctx, method, url, bodyReader)
	} else {
		req, err = http.NewRequestWithContext(ctx, method, url, nil)
	}
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return nil
	}

	req.Header.Set("User-Agent", "Appie/9.28 (iPhone17,3; iPhone; CPU OS 26_1 like Mac OS X)")
	req.Header.Set("x-client-name", "appie-ios")
	req.Header.Set("x-client-version", "9.28")
	req.Header.Set("x-application", "AHWEBSHOP")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.AccessToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Request failed: %v", err)
		return nil
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %d %s\n\n", resp.StatusCode, resp.Status)
	return resp.Header
}

func printRateLimitHeaders(headers http.Header) {
	if headers == nil {
		return
	}

	// Common rate limit header patterns
	rateLimitPatterns := []string{
		"rate", "limit", "quota", "throttle", "retry", "backoff",
		"x-ratelimit", "x-rate-limit", "ratelimit",
		"x-quota", "x-throttle",
	}

	fmt.Println("Rate Limit Related Headers:")
	found := false

	// Sort headers for consistent output
	keys := make([]string, 0, len(headers))
	for k := range headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		lower := strings.ToLower(k)
		for _, pattern := range rateLimitPatterns {
			if strings.Contains(lower, pattern) {
				fmt.Printf("  %s: %v\n", k, headers[k])
				found = true
				break
			}
		}
	}

	if !found {
		fmt.Println("  (none found)")
	}

	fmt.Println("\nAll Response Headers:")
	for _, k := range keys {
		fmt.Printf("  %s: %v\n", k, headers[k])
	}
}
