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

	if client.AccessToken() != "test-token" {
		t.Errorf("expected access token 'test-token', got '%s'", client.AccessToken())
	}
	if client.RefreshTokenValue() != "test-refresh" {
		t.Errorf("expected refresh token 'test-refresh', got '%s'", client.RefreshTokenValue())
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
	url := client.LoginURL()
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

	if err := client.SaveConfig(); err != nil {
		t.Fatal(err)
	}

	client2 := New(WithConfigPath(tmpFile.Name()))
	if err := client2.LoadConfig(); err != nil {
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
			json.NewEncoder(w).Encode(Token{
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
	if client.AccessToken() != "new-access" {
		t.Errorf("expected access token 'new-access', got '%s'", client.AccessToken())
	}
}

func TestNoAutoRefreshForAuthEndpoints(t *testing.T) {
	refreshCount := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/mobile-auth/v1/auth/token/refresh" {
			refreshCount++
			json.NewEncoder(w).Encode(Token{
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
	// Calling RefreshToken directly should not trigger auto-refresh recursion
	if err := client.RefreshToken(ctx); err != nil {
		t.Fatal(err)
	}

	if refreshCount != 1 {
		t.Errorf("expected refresh to be called exactly once, got %d", refreshCount)
	}
	if client.AccessToken() != "new-access" {
		t.Errorf("expected 'new-access', got '%s'", client.AccessToken())
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
