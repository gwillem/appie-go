package appie

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// The Albert Heijn market is selected by the OAuth client_id, not by the host:
// nl uses appie-ios, be uses appie-be-ios. The API/login hosts stay on the
// default backend.
func TestWithCountrySetsClientID(t *testing.T) {
	tests := []struct {
		country, want string
	}{
		{"be", "appie-be-ios"},
		{"nl", "appie-ios"},
		{"BE", "appie-be-ios"},
	}
	for _, tt := range tests {
		client := New(WithCountry(tt.country))
		if client.clientID != tt.want {
			t.Errorf("WithCountry(%q): expected clientID %q, got %q", tt.country, tt.want, client.clientID)
		}
		if client.baseURL != defaultBaseURL {
			t.Errorf("WithCountry(%q): expected default baseURL %q, got %q", tt.country, defaultBaseURL, client.baseURL)
		}
	}
}

// The assortment/pricing market is selected by the x-application header:
// nl uses AHWEBSHOP, be uses AHBEWEBSHOP.
func TestWithCountrySetsApplication(t *testing.T) {
	tests := []struct {
		country, want string
	}{
		{"be", "AHBEWEBSHOP"},
		{"nl", "AHWEBSHOP"},
	}
	for _, tt := range tests {
		client := New(WithCountry(tt.country))
		if client.application != tt.want {
			t.Errorf("WithCountry(%q): expected application %q, got %q", tt.country, tt.want, client.application)
		}
	}
}

func TestRequestCarriesApplicationHeader(t *testing.T) {
	tests := []struct {
		name, country, want string
	}{
		{"belgium", "be", "AHBEWEBSHOP"},
		{"default stays dutch", "", "AHWEBSHOP"},
	}
	for _, tt := range tests {
		var got string
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			got = r.Header.Get("x-application")
			w.Write([]byte(`{}`))
		}))

		opts := []Option{WithBaseURL(srv.URL), WithTokens("test", "test")}
		if tt.country != "" {
			opts = append(opts, WithCountry(tt.country))
		}
		client := New(opts...)

		var result map[string]string
		if err := client.DoRequest(context.Background(), http.MethodGet, "/test-endpoint", nil, &result); err != nil {
			t.Fatalf("%s: unexpected error: %v", tt.name, err)
		}
		if got != tt.want {
			t.Errorf("%s: expected x-application %q, got %q", tt.name, tt.want, got)
		}
		srv.Close()
	}
}

func TestLoginURLCarriesCountryClientID(t *testing.T) {
	client := New(WithCountry("be"))
	url := client.loginURL()
	if !strings.Contains(url, "client_id=appie-be-ios") {
		t.Errorf("expected login URL to contain client_id=appie-be-ios, got %s", url)
	}
}

// A saved country is restored on load and re-selects the client_id, so callers
// don't have to re-specify it every run.
func TestConfigStoresAndRestoresCountry(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "appie-test-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	client := New(WithConfigPath(tmpFile.Name()), WithCountry("be"), WithTokens("access", "refresh"))
	if err := client.saveConfig(); err != nil {
		t.Fatal(err)
	}

	client2 := New(WithConfigPath(tmpFile.Name()))
	if err := client2.loadConfig(); err != nil {
		t.Fatal(err)
	}
	if client2.clientID != "appie-be-ios" {
		t.Errorf("expected restored clientID 'appie-be-ios', got %q", client2.clientID)
	}
}

// An explicitly configured country wins over whatever is stored in the config.
func TestExplicitCountryBeatsStored(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "appie-test-*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	client := New(WithConfigPath(tmpFile.Name()), WithCountry("be"), WithTokens("access", "refresh"))
	if err := client.saveConfig(); err != nil {
		t.Fatal(err)
	}

	client2 := New(WithConfigPath(tmpFile.Name()), WithCountry("nl"))
	if err := client2.loadConfig(); err != nil {
		t.Fatal(err)
	}
	if client2.clientID != "appie-ios" {
		t.Errorf("expected explicit nl to beat stored be: expected clientID 'appie-ios', got %q", client2.clientID)
	}
}
