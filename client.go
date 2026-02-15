package appie

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	defaultBaseURL    = "https://api.ah.nl"
	defaultUserAgent  = "Appie/9.28 (iPhone17,3; iPhone; CPU OS 26_1 like Mac OS X)"
	defaultClientID   = "appie-ios"
	defaultClientVersion = "9.28"
)

// Client is the AH API client. It handles authentication, token management,
// and provides methods to interact with products, orders, shopping lists, and more.
//
// Client is safe for concurrent use. Token state is protected by a mutex.
type Client struct {
	httpClient *http.Client
	baseURL    string
	userAgent  string
	clientID   string
	clientVersion string

	mu           sync.RWMutex
	accessToken  string
	refreshToken string
	memberID     string
	expiresAt    time.Time
	orderID      string
	orderHash    string

	configPath   string
	loginBaseURL string      // overridable for testing; defaults to "https://login.ah.nl"
	openBrowser  func(string) // overridable for testing; nil uses default
}

// Option configures the client. Use With* functions to create options.
type Option func(*Client)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// WithBaseURL sets a custom base URL.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithTokens sets the access and refresh tokens.
func WithTokens(accessToken, refreshToken string) Option {
	return func(c *Client) {
		c.accessToken = accessToken
		c.refreshToken = refreshToken
	}
}

// WithConfigPath sets the path to the config file.
func WithConfigPath(path string) Option {
	return func(c *Client) {
		c.configPath = path
	}
}

// New creates a new AH API client.
func New(opts ...Option) *Client {
	c := &Client{
		httpClient:    http.DefaultClient,
		baseURL:       defaultBaseURL,
		userAgent:     defaultUserAgent,
		clientID:      defaultClientID,
		clientVersion: defaultClientVersion,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// NewWithConfig creates a new client and loads config from the given path.
func NewWithConfig(configPath string) (*Client, error) {
	c := New(WithConfigPath(configPath))

	if err := c.loadConfig(); err != nil {
		if os.IsNotExist(err) {
			return c, nil // Config doesn't exist yet, that's OK
		}
		return nil, err
	}

	return c, nil
}

// loadConfig loads the configuration from the config file.
func (c *Client) loadConfig() error {
	if c.configPath == "" {
		return fmt.Errorf("no config path set")
	}

	data, err := os.ReadFile(c.configPath)
	if err != nil {
		return err
	}

	var cfg config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	c.mu.Lock()
	c.accessToken = cfg.AccessToken
	c.refreshToken = cfg.RefreshToken
	c.memberID = cfg.MemberID
	c.expiresAt = cfg.ExpiresAt
	c.mu.Unlock()

	return nil
}

// saveConfig saves the current configuration to the config file.
func (c *Client) saveConfig() error {
	if c.configPath == "" {
		return fmt.Errorf("no config path set")
	}

	c.mu.RLock()
	cfg := config{
		AccessToken:  c.accessToken,
		RefreshToken: c.refreshToken,
		MemberID:     c.memberID,
		ExpiresAt:    c.expiresAt,
	}
	c.mu.RUnlock()

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(c.configPath, data, 0600)
}

// IsAuthenticated returns true if the client has an access token.
func (c *Client) IsAuthenticated() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.accessToken != ""
}

// setHeaders sets the common headers for API requests.
func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("x-client-name", c.clientID)
	req.Header.Set("x-client-version", c.clientVersion)
	req.Header.Set("x-application", "AHWEBSHOP")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	c.mu.RLock()
	if c.accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.accessToken)
	}
	if c.orderID != "" {
		req.Header.Set("appie-current-order-id", c.orderID)
		if c.orderHash != "" {
			req.Header.Set("appie-current-order-hash", c.orderHash)
		}
	}
	c.mu.RUnlock()
}

// ensureFreshToken refreshes the access token if it has expired and a refresh token is available.
// Auth endpoints are excluded to avoid infinite loops.
func (c *Client) ensureFreshToken(ctx context.Context, path string) {
	// Don't auto-refresh for auth endpoints
	if strings.HasPrefix(path, "/mobile-auth/") {
		return
	}

	c.mu.RLock()
	expired := !c.expiresAt.IsZero() && time.Now().After(c.expiresAt)
	hasRefresh := c.refreshToken != ""
	c.mu.RUnlock()

	if expired && hasRefresh {
		// Best-effort refresh; if it fails, the original request will proceed
		// with the expired token and the API will return an appropriate error.
		if err := c.refreshAccessToken(ctx); err == nil {
			_ = c.saveConfig()
		}
	}
}

// DoRequest performs an HTTP request and decodes the response.
func (c *Client) DoRequest(ctx context.Context, method, path string, body, result any) error {
	c.ensureFreshToken(ctx, path)

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiErr apiError
		if json.Unmarshal(respBody, &apiErr) == nil && (apiErr.Code != "" || apiErr.Message != "") {
			return &apiErr
		}
		return fmt.Errorf("API error: %d %s", resp.StatusCode, string(respBody))
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// DoGraphQL performs a GraphQL request.
func (c *Client) DoGraphQL(ctx context.Context, query string, variables map[string]any, result any) error {
	req := graphQLRequest{
		Query:     query,
		Variables: variables,
	}

	var resp graphQLResponse[json.RawMessage]
	if err := c.DoRequest(ctx, http.MethodPost, "/graphql", req, &resp); err != nil {
		return err
	}

	if len(resp.Errors) > 0 {
		return fmt.Errorf("graphql error: %s", resp.Errors[0].Message)
	}

	if result != nil {
		if err := json.Unmarshal(resp.Data, result); err != nil {
			return fmt.Errorf("failed to decode graphql response: %w", err)
		}
	}

	return nil
}
