package appie

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// loginURL returns the URL for browser-based login.
func (c *Client) loginURL() string {
	return fmt.Sprintf("%s/login?client_id=%s&response_type=code&redirect_uri=appie://login-exit", c.loginBaseURL, c.clientID)
}

// exchangeCode exchanges an authorization code for tokens.
func (c *Client) exchangeCode(ctx context.Context, code string) error {
	body := map[string]string{
		"clientId": c.clientID,
		"code":     code,
	}

	var tok token
	if err := c.DoRequest(ctx, http.MethodPost, "/mobile-auth/v1/auth/token", body, &tok); err != nil {
		return fmt.Errorf("failed to exchange code: %w", err)
	}

	c.mu.Lock()
	c.accessToken = tok.AccessToken
	c.refreshToken = tok.RefreshToken
	c.memberID = tok.MemberID
	if tok.ExpiresIn > 0 {
		c.expiresAt = time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second)
	}
	c.mu.Unlock()

	_ = c.saveConfig()
	return nil
}

// refreshAccessToken refreshes the access token using the refresh token.
// Concurrent callers are serialized to avoid burning a rotated refresh token.
func (c *Client) refreshAccessToken(ctx context.Context) error {
	c.refreshMu.Lock()
	defer c.refreshMu.Unlock()
	return c.doRefreshAccessToken(ctx)
}

func (c *Client) doRefreshAccessToken(ctx context.Context) error {
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

	var tok token
	if err := c.DoRequest(ctx, http.MethodPost, "/mobile-auth/v1/auth/token/refresh", body, &tok); err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	c.mu.Lock()
	c.accessToken = tok.AccessToken
	c.refreshToken = tok.RefreshToken
	if tok.ExpiresIn > 0 {
		c.expiresAt = time.Now().Add(time.Duration(tok.ExpiresIn) * time.Second)
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

	var tok token
	if err := c.DoRequest(ctx, http.MethodPost, "/mobile-auth/v1/auth/token/anonymous", body, &tok); err != nil {
		return fmt.Errorf("failed to get anonymous token: %w", err)
	}

	c.mu.Lock()
	c.accessToken = tok.AccessToken
	c.refreshToken = tok.RefreshToken
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
