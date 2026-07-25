package appie

import (
	"path/filepath"
	"strings"
	"testing"
)

// The Albert Heijn market is selected by the OAuth client_id, not by the host:
// nl uses appie-ios, be uses appie-be-ios. The API/login hosts stay on the
// default backend.
func TestWithCountrySetsClientID(t *testing.T) {
	tests := []struct {
		country      string
		wantClientID string
	}{
		{"be", "appie-be-ios"},
		{"nl", "appie-ios"},
		{"BE", "appie-be-ios"},
	}
	for _, tt := range tests {
		c := New(WithCountry(tt.country))
		if c.clientID != tt.wantClientID {
			t.Errorf("WithCountry(%q): clientID = %q, want %q", tt.country, c.clientID, tt.wantClientID)
		}
		if c.baseURL != defaultBaseURL {
			t.Errorf("WithCountry(%q): baseURL = %q, want default %q (host must not change)", tt.country, c.baseURL, defaultBaseURL)
		}
	}
}

func TestLoginURLCarriesCountryClientID(t *testing.T) {
	c := New(WithCountry("be"))
	got := c.loginURL()
	if !strings.Contains(got, "client_id=appie-be-ios") {
		t.Errorf("loginURL() = %q, want it to contain client_id=appie-be-ios", got)
	}
}

// A saved country is restored on load and re-selects the client_id, so callers
// don't have to re-specify it every run.
func TestConfigStoresAndRestoresCountry(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")

	c := New(WithConfigPath(path), WithCountry("be"), WithTokens("access", "refresh"))
	if err := c.saveConfig(); err != nil {
		t.Fatalf("saveConfig: %v", err)
	}

	c2 := New(WithConfigPath(path))
	if err := c2.loadConfig(); err != nil {
		t.Fatalf("loadConfig: %v", err)
	}
	if c2.clientID != "appie-be-ios" {
		t.Errorf("restored clientID = %q, want appie-be-ios", c2.clientID)
	}
}

// An explicitly configured country wins over whatever is stored in the config.
func TestExplicitCountryBeatsStored(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")

	saver := New(WithConfigPath(path), WithCountry("be"), WithTokens("access", "refresh"))
	if err := saver.saveConfig(); err != nil {
		t.Fatalf("saveConfig: %v", err)
	}

	c := New(WithConfigPath(path), WithCountry("nl"))
	if err := c.loadConfig(); err != nil {
		t.Fatalf("loadConfig: %v", err)
	}
	if c.clientID != "appie-ios" {
		t.Errorf("clientID = %q, want appie-ios (explicit nl should beat stored be)", c.clientID)
	}
}
