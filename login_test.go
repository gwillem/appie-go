package appie

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/mobile-auth/v1/auth/token":
			json.NewEncoder(w).Encode(token{
				AccessToken:  "test-access",
				RefreshToken: "test-refresh",
				MemberID:     "member-123",
				ExpiresIn:    86400,
			})
		default:
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<html>mock login</html>`))
		}
	}))
	defer mock.Close()

	client := New(WithBaseURL(mock.URL))
	client.loginBaseURL = mock.URL

	proxyURLCh := make(chan string, 1)
	client.openBrowser = func(u string) {
		proxyURLCh <- u
	}

	ctx := context.Background()
	errCh := make(chan error, 1)
	go func() {
		errCh <- client.Login(ctx)
	}()

	proxyURL := <-proxyURLCh
	u, err := url.Parse(proxyURL)
	if err != nil {
		t.Fatal(err)
	}

	callbackURL := fmt.Sprintf("http://%s/callback?code=test-auth-code", u.Host)
	resp, err := http.Get(callbackURL)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	if err := <-errCh; err != nil {
		t.Fatal(err)
	}

	if client.accessToken != "test-access" {
		t.Errorf("expected 'test-access', got '%s'", client.accessToken)
	}
	if client.refreshToken != "test-refresh" {
		t.Errorf("expected 'test-refresh', got '%s'", client.refreshToken)
	}
	if client.memberID != "member-123" {
		t.Errorf("expected 'member-123', got '%s'", client.memberID)
	}
}

func TestLoginRewritesBody(t *testing.T) {
	body := `{"pageProps":{"__N_REDIRECT":"appie://login-exit?code=abc123","other":"appie://login-exit"}}`
	resp := &http.Response{
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": {"application/json"},
		},
		Body: io.NopCloser(strings.NewReader(body)),
	}

	err := rewriteLoginResponse(resp, "http://127.0.0.1:9999", "login.ah.nl")
	if err != nil {
		t.Fatal(err)
	}

	result, _ := io.ReadAll(resp.Body)
	resultStr := string(result)

	if strings.Contains(resultStr, "appie://") {
		t.Errorf("body still contains appie:// : %s", resultStr)
	}
	if !strings.Contains(resultStr, "http://127.0.0.1:9999/callback") {
		t.Errorf("body doesn't contain callback URL: %s", resultStr)
	}
}

func TestLoginRewritesLocationHeader(t *testing.T) {
	resp := &http.Response{
		StatusCode: 302,
		Header: http.Header{
			"Location":     {"appie://login-exit?code=xyz789"},
			"Content-Type": {"text/html"},
		},
		Body: io.NopCloser(strings.NewReader("")),
	}

	err := rewriteLoginResponse(resp, "http://127.0.0.1:9999", "login.ah.nl")
	if err != nil {
		t.Fatal(err)
	}

	loc := resp.Header.Get("Location")
	if strings.Contains(loc, "appie://") {
		t.Errorf("Location still contains appie:// : %s", loc)
	}
	if !strings.Contains(loc, "http://127.0.0.1:9999/callback") {
		t.Errorf("Location doesn't contain callback URL: %s", loc)
	}
	if !strings.Contains(loc, "code=xyz789") {
		t.Errorf("Location lost the code param: %s", loc)
	}
}

func TestLoginContextCancel(t *testing.T) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))
	defer mock.Close()

	client := New(WithBaseURL(mock.URL))
	client.loginBaseURL = mock.URL
	client.openBrowser = func(string) {} // no-op

	ctx, cancel := context.WithCancel(context.Background())

	errCh := make(chan error, 1)
	go func() {
		errCh <- client.Login(ctx)
	}()

	cancel()

	err := <-errCh
	if err != context.Canceled {
		t.Errorf("expected context.Canceled, got %v", err)
	}
}
