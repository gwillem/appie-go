package appie

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	loginURLTemplate = "https://login.ah.nl/login?client_id=%s&response_type=code&redirect_uri=appie://login-exit"
)

// LoginURL returns the URL for browser-based login.
// After login, the browser will redirect to appie://login-exit?code=...
// Extract the code and pass it to ExchangeCode.
func (c *Client) LoginURL() string {
	return fmt.Sprintf(loginURLTemplate, c.clientID)
}

// ExchangeCode exchanges an authorization code for tokens.
func (c *Client) ExchangeCode(ctx context.Context, code string) error {
	body := map[string]string{
		"clientId": c.clientID,
		"code":     code,
	}

	var token Token
	if err := c.doRequest(ctx, http.MethodPost, "/mobile-auth/v1/auth/token", body, &token); err != nil {
		return fmt.Errorf("failed to exchange code: %w", err)
	}

	c.mu.Lock()
	c.accessToken = token.AccessToken
	c.refreshToken = token.RefreshToken
	c.memberID = token.MemberID
	if token.ExpiresIn > 0 {
		c.expiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}
	c.mu.Unlock()

	return nil
}

// RefreshToken refreshes the access token using the refresh token.
func (c *Client) RefreshToken(ctx context.Context) error {
	c.mu.RLock()
	rt := c.refreshToken
	c.mu.RUnlock()

	if rt == "" {
		return fmt.Errorf("no refresh token available")
	}

	body := map[string]string{
		"clientId":     c.clientID,
		"refreshToken": rt,
	}

	var token Token
	if err := c.doRequest(ctx, http.MethodPost, "/mobile-auth/v1/auth/token/refresh", body, &token); err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	c.mu.Lock()
	c.accessToken = token.AccessToken
	c.refreshToken = token.RefreshToken
	if token.ExpiresIn > 0 {
		c.expiresAt = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	}
	c.mu.Unlock()

	return nil
}

// GetAnonymousToken gets an anonymous token (no login required).
// This is useful for browsing products without authentication.
func (c *Client) GetAnonymousToken(ctx context.Context) error {
	body := map[string]string{
		"clientId": c.clientID,
	}

	var token Token
	if err := c.doRequest(ctx, http.MethodPost, "/mobile-auth/v1/auth/token/anonymous", body, &token); err != nil {
		return fmt.Errorf("failed to get anonymous token: %w", err)
	}

	c.mu.Lock()
	c.accessToken = token.AccessToken
	c.refreshToken = token.RefreshToken
	c.mu.Unlock()

	return nil
}

// Logout clears the authentication tokens.
func (c *Client) Logout() {
	c.mu.Lock()
	c.accessToken = ""
	c.refreshToken = ""
	c.memberID = ""
	c.mu.Unlock()
}
