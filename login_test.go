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
	"sync/atomic"
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

func TestSanitizeCookie(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{
			"sid=abc123; Path=/; Secure; HttpOnly; SameSite=None",
			"sid=abc123; Path=/; HttpOnly",
		},
		{
			"token=xyz; Domain=.ah.nl; Secure; Path=/; SameSite=Lax",
			"token=xyz; Path=/",
		},
		{
			"simple=val",
			"simple=val",
		},
		{
			"a=b; secure; DOMAIN=login.ah.nl; samesite=strict; HttpOnly",
			"a=b; HttpOnly",
		},
	}
	for _, tt := range tests {
		got := sanitizeCookie(tt.in)
		if got != tt.want {
			t.Errorf("sanitizeCookie(%q)\n got %q\nwant %q", tt.in, got, tt.want)
		}
	}
}

func TestLoginRewritesCookies(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": {"text/html"},
			"Set-Cookie": {
				"sid=abc; Path=/; Secure; HttpOnly; SameSite=None; Domain=.ah.nl",
				"token=xyz; Secure; Path=/",
			},
		},
		Body: io.NopCloser(strings.NewReader("<html>ok</html>")),
	}

	err := rewriteLoginResponse(resp, "http://127.0.0.1:9999", "login.ah.nl")
	if err != nil {
		t.Fatal(err)
	}

	cookies := resp.Header.Values("Set-Cookie")
	for _, c := range cookies {
		lower := strings.ToLower(c)
		if strings.Contains(lower, "secure") {
			t.Errorf("cookie still contains Secure: %s", c)
		}
		if strings.Contains(lower, "samesite") {
			t.Errorf("cookie still contains SameSite: %s", c)
		}
		if strings.Contains(lower, "domain") {
			t.Errorf("cookie still contains Domain: %s", c)
		}
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

func TestTakeNextRedirectCode(t *testing.T) {
	codeCh := make(chan string, 1)
	var taken atomic.Bool
	doneURL := "http://127.0.0.1:9999/login-done"
	newResp := func(body string) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     http.Header{"Content-Type": {"application/json"}},
			Body:       io.NopCloser(strings.NewReader(body)),
		}
	}

	first := newResp(`{"pageProps":{"__N_REDIRECT":"appie://login-exit?code=abc123","__N_REDIRECT_STATUS":302},"__N_SSP":true}`)
	if err := takeNextRedirectCode(first, codeCh, &taken, doneURL); err != nil {
		t.Fatal(err)
	}
	if code := <-codeCh; code != "abc123" {
		t.Errorf("expected code 'abc123', got '%s'", code)
	}
	body, _ := io.ReadAll(first.Body)
	if strings.Contains(string(body), "appie://") || !strings.Contains(string(body), doneURL) {
		t.Errorf("first response not redirected to the success page: %s", body)
	}

	second := newResp(`{"pageProps":{},"__N_SSP":true}`)
	if err := takeNextRedirectCode(second, codeCh, &taken, doneURL); err != nil {
		t.Fatal(err)
	}
	body, _ = io.ReadAll(second.Body)
	if !strings.Contains(string(body), doneURL) {
		t.Errorf("second response not redirected to the success page: %s", body)
	}
}

func TestLoginReadsCodeFromNextRedirect(t *testing.T) {
	installationIDCh := make(chan string, 1)
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/secure/oauth/authorize":
			installationIDCh <- r.Header.Get("x-fraud-detection-installation-id")
			w.Write([]byte("ok"))
		case strings.HasSuffix(r.URL.Path, "/ingelogd.json"):
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"pageProps":{"__N_REDIRECT":"appie://login-exit?code=next-code","__N_REDIRECT_STATUS":302},"__N_SSP":true}`))
		case r.URL.Path == "/mobile-auth/v1/auth/token":
			json.NewEncoder(w).Encode(token{AccessToken: "next-access", RefreshToken: "next-refresh", ExpiresIn: 86400})
		}
	}))
	defer mock.Close()

	client := New(WithBaseURL(mock.URL))
	client.loginBaseURL = mock.URL
	proxyURLCh := make(chan string, 1)
	client.openBrowser = func(u string) { proxyURLCh <- u }

	errCh := make(chan error, 1)
	go func() { errCh <- client.Login(context.Background()) }()

	loginURL := <-proxyURLCh
	u, err := url.Parse(loginURL)
	if err != nil {
		t.Fatal(err)
	}
	if u.Path != "/secure/oauth/authorize" {
		t.Errorf("expected login to start at /secure/oauth/authorize, got %s", u.Path)
	}
	for _, path := range []string{u.RequestURI(), "/login/_next/data/build/nl/ingelogd.json", loginDonePath} {
		resp, err := http.Get(fmt.Sprintf("http://%s%s", u.Host, path))
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
	}

	if err := <-errCh; err != nil {
		t.Fatal(err)
	}
	if id := <-installationIDCh; id == "" {
		t.Error("authorize request did not carry x-fraud-detection-installation-id")
	}
	if client.accessToken != "next-access" {
		t.Errorf("expected 'next-access', got '%s'", client.accessToken)
	}
}
