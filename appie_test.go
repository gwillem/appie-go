package appie

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := New()
	if client == nil {
		t.Fatal("expected client, got nil")
	}
	if client.baseURL != defaultBaseURL {
		t.Errorf("expected baseURL %s, got %s", defaultBaseURL, client.baseURL)
	}
}

func TestNewWithConfig(t *testing.T) {
	// Create a temp config file
	tmpFile, err := os.CreateTemp("", "appie-test-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(`{"access_token": "test-token", "refresh_token": "test-refresh"}`)
	if err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	client, err := NewWithConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if client.accessToken != "test-token" {
		t.Errorf("expected access token 'test-token', got '%s'", client.accessToken)
	}
	if client.refreshToken != "test-refresh" {
		t.Errorf("expected refresh token 'test-refresh', got '%s'", client.refreshToken)
	}
}

func TestNewWithConfigMissing(t *testing.T) {
	client, err := NewWithConfig("/nonexistent/path.json")
	if err != nil {
		t.Fatalf("unexpected error for missing config: %v", err)
	}
	if client == nil {
		t.Fatal("expected client even with missing config")
	}
}

func TestLoginURL(t *testing.T) {
	client := New()
	url := client.loginURL()
	expected := "https://login.ah.nl/login?client_id=appie-ios&response_type=code&redirect_uri=appie://login-exit"
	if url != expected {
		t.Errorf("expected %s, got %s", expected, url)
	}
}

func TestConfigExpiresAtRoundTrip(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "appie-test-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	client := New(WithConfigPath(tmpFile.Name()), WithTokens("access", "refresh"))
	client.mu.Lock()
	client.expiresAt = time.Now().Add(24 * time.Hour).Truncate(time.Second)
	client.mu.Unlock()

	if err := client.saveConfig(); err != nil {
		t.Fatal(err)
	}

	client2 := New(WithConfigPath(tmpFile.Name()))
	if err := client2.loadConfig(); err != nil {
		t.Fatal(err)
	}

	client.mu.RLock()
	expected := client.expiresAt
	client.mu.RUnlock()

	client2.mu.RLock()
	got := client2.expiresAt
	client2.mu.RUnlock()

	if !expected.Equal(got) {
		t.Errorf("expiresAt mismatch: expected %v, got %v", expected, got)
	}
}

func TestAutoRefreshOnExpiredToken(t *testing.T) {
	refreshCalled := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/mobile-auth/v1/auth/token/refresh":
			refreshCalled = true
			json.NewEncoder(w).Encode(token{
				AccessToken:  "new-access",
				RefreshToken: "new-refresh",
				ExpiresIn:    86400,
			})
		case "/test-endpoint":
			json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("expired-access", "my-refresh"))
	client.mu.Lock()
	client.expiresAt = time.Now().Add(-1 * time.Hour) // expired
	client.mu.Unlock()

	ctx := context.Background()
	var result map[string]string
	if err := client.doRequest(ctx, http.MethodGet, "/test-endpoint", nil, &result); err != nil {
		t.Fatal(err)
	}

	if !refreshCalled {
		t.Error("expected refresh to be called for expired token")
	}
	if client.accessToken != "new-access" {
		t.Errorf("expected access token 'new-access', got '%s'", client.accessToken)
	}
}

func TestNoAutoRefreshForAuthEndpoints(t *testing.T) {
	refreshCount := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/mobile-auth/v1/auth/token/refresh" {
			refreshCount++
			json.NewEncoder(w).Encode(token{
				AccessToken:  "new-access",
				RefreshToken: "new-refresh",
				ExpiresIn:    86400,
			})
		}
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("expired-access", "my-refresh"))
	client.mu.Lock()
	client.expiresAt = time.Now().Add(-1 * time.Hour)
	client.mu.Unlock()

	ctx := context.Background()
	// Calling refreshAccessToken directly should not trigger auto-refresh recursion
	if err := client.refreshAccessToken(ctx); err != nil {
		t.Fatal(err)
	}

	if refreshCount != 1 {
		t.Errorf("expected refresh to be called exactly once, got %d", refreshCount)
	}
	if client.accessToken != "new-access" {
		t.Errorf("expected 'new-access', got '%s'", client.accessToken)
	}
}

func TestNoAutoRefreshWhenNotExpired(t *testing.T) {
	refreshCalled := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/mobile-auth/v1/auth/token/refresh" {
			refreshCalled = true
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("valid-access", "my-refresh"))
	client.mu.Lock()
	client.expiresAt = time.Now().Add(24 * time.Hour) // not expired
	client.mu.Unlock()

	ctx := context.Background()
	var result map[string]string
	if err := client.doRequest(ctx, http.MethodGet, "/test-endpoint", nil, &result); err != nil {
		t.Fatal(err)
	}

	if refreshCalled {
		t.Error("refresh should not be called when token is not expired")
	}
}

func TestGetBonusGroupProductsMock(t *testing.T) {
	callCount := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		switch {
		case r.URL.Path == "/mobile-services/bonuspage/v3/metadata":
			json.NewEncoder(w).Encode(bonusMetadataResponse{
				Periods: []struct {
					BonusStartDate string `json:"bonusStartDate"`
					BonusEndDate   string `json:"bonusEndDate"`
					Tabs           []struct {
						Description     string `json:"description"`
						URLMetadataList []struct {
							URL         string `json:"url"`
							Count       int    `json:"count"`
							BonusType   string `json:"bonusType"`
							Description string `json:"description"`
						} `json:"urlMetadataList"`
					} `json:"tabs"`
				}{
					{BonusStartDate: "2026-02-09", BonusEndDate: "2026-02-15"},
				},
			})
		case r.URL.Path == "/graphql":
			json.NewEncoder(w).Encode(map[string]any{
				"data": map[string]any{
					"bonusPromotions": []map[string]any{
						{
							"id":           "764081",
							"title":        "Alle Hak*",
							"productCount": 2,
							"products": []map[string]any{
								{
									"id":            189671,
									"title":         "Hak Appelmoes extra kwaliteit",
									"brand":         "Hak",
									"category":      "Soepen, sauzen/Appelmoes",
									"salesUnitSize": "355 g",
									"icons":         []string{"NUTRISCORE_A"},
									"availability":  map[string]any{"isOrderable": true},
									"priceV2": map[string]any{
										"now": map[string]any{"amount": 2.19},
										"was": map[string]any{"amount": 2.19},
										"promotionLabel": map[string]any{
											"tiers": []map[string]any{
												{"mechanism": "X_PLUS_Y_FREE", "description": "1+1 gratis"},
											},
										},
									},
									"imagePack": []map[string]any{
										{"large": map[string]any{"url": "https://example.com/hak.jpg", "width": 800, "height": 800}},
									},
								},
								{
									"id":            456789,
									"title":         "Hak Bonensalade",
									"brand":         "Hak",
									"category":      "Groente/Bonen",
									"salesUnitSize": "370 ml",
									"icons":         []string{"NUTRISCORE_B", "VEGETARIAN"},
									"availability":  map[string]any{"isOrderable": true},
									"priceV2": map[string]any{
										"now": map[string]any{"amount": 1.99},
										"was": map[string]any{"amount": 2.49},
									},
									"imagePack": []map[string]any{},
								},
							},
						},
					},
				},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))
	ctx := context.Background()

	products, err := client.GetBonusGroupProducts(ctx, "764081")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(products) != 2 {
		t.Fatalf("expected 2 products, got %d", len(products))
	}

	// Verify first product
	p := products[0]
	if p.ID != 189671 {
		t.Errorf("expected ID 189671, got %d", p.ID)
	}
	if p.Title != "Hak Appelmoes extra kwaliteit" {
		t.Errorf("expected title 'Hak Appelmoes extra kwaliteit', got %q", p.Title)
	}
	if p.Brand != "Hak" {
		t.Errorf("expected brand 'Hak', got %q", p.Brand)
	}
	if p.NutriScore != "A" {
		t.Errorf("expected NutriScore 'A', got %q", p.NutriScore)
	}
	if p.BonusMechanism != "1+1 gratis" {
		t.Errorf("expected BonusMechanism '1+1 gratis', got %q", p.BonusMechanism)
	}
	if !p.IsBonus {
		t.Error("expected IsBonus true")
	}
	if !p.IsOrderable {
		t.Error("expected IsOrderable true")
	}
	if p.Price.Now != 2.19 {
		t.Errorf("expected price 2.19, got %.2f", p.Price.Now)
	}
	if len(p.Images) != 1 {
		t.Errorf("expected 1 image, got %d", len(p.Images))
	}

	// Verify second product (no promotion label, no images)
	p2 := products[1]
	if p2.ID != 456789 {
		t.Errorf("expected ID 456789, got %d", p2.ID)
	}
	if p2.BonusMechanism != "" {
		t.Errorf("expected empty BonusMechanism, got %q", p2.BonusMechanism)
	}
	if p2.Price.Was != 2.49 {
		t.Errorf("expected was price 2.49, got %.2f", p2.Price.Was)
	}

	// Should have made 2 calls: metadata + graphql
	if callCount != 2 {
		t.Errorf("expected 2 API calls, got %d", callCount)
	}
}

func TestBonusGroupToProductHasSegmentID(t *testing.T) {
	bg := &bonusGroupResponse{
		ID:                  "764081",
		SegmentDescription:  "Alle Hak*",
		DiscountDescription: "1+1 GRATIS",
		Category:            "Groente, aardappelen",
	}
	p := bg.toProduct()
	if p.BonusSegmentID != "764081" {
		t.Errorf("expected BonusSegmentID '764081', got %q", p.BonusSegmentID)
	}
	if p.ID != 0 {
		t.Errorf("expected ID 0 for bonus group, got %d", p.ID)
	}
}

func TestFetchNutritionalInfoMock(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"product": map[string]any{
					"id": 123,
					"tradeItem": map[string]any{
						"nutritions": []map[string]any{
							{
								"nutrients": []map[string]any{
									{"type": "FAT", "name": "Fat", "value": "15.5 g"},
									{"type": "PROTEIN", "name": "Protein", "value": "8.2 g"},
								},
							},
						},
					},
				},
			},
		})
	}))
	defer srv.Close()

	client := New(WithBaseURL(srv.URL), WithTokens("test", "test"))
	ctx := context.Background()

	info, err := client.fetchNutritionalInfo(ctx, 123)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(info) != 2 {
		t.Fatalf("expected 2 nutrients, got %d", len(info))
	}

	if info[0].Name != "Fat" || info[0].Type != "FAT" || info[0].Value != "15.5 g" {
		t.Errorf("unexpected first nutrient: %+v", info[0])
	}
	if info[1].Name != "Protein" || info[1].Type != "PROTEIN" || info[1].Value != "8.2 g" {
		t.Errorf("unexpected second nutrient: %+v", info[1])
	}
}
