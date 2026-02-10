package appie

import (
	"bytes"
	"cmp"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

const loginSuccessPage = `<!DOCTYPE html>
<html><head><title>Login Successful</title></head>
<body style="font-family:system-ui;max-width:500px;margin:80px auto;text-align:center">
<h1>Login successful!</h1>
<p>You can close this tab.</p>
<script>setTimeout(function(){window.close()},500)</script>
</body></html>`

// Login performs the full browser-based login flow. It starts a local reverse
// proxy to the AH login page, opens the user's browser, waits for the
// authorization code callback, and exchanges it for tokens.
//
// The proxy rewrites appie:// redirect URLs in the login response to a local
// callback endpoint, so the browser never navigates to the custom scheme.
//
// Cancel the context to abort the login flow.
func (c *Client) Login(ctx context.Context) error {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return fmt.Errorf("failed to start login server: %w", err)
	}

	localOrigin := fmt.Sprintf("http://%s", listener.Addr())
	codeCh := make(chan string, 1)

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		codeCh <- r.URL.Query().Get("code")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, loginSuccessPage)
	})

	loginBaseURL := cmp.Or(c.loginBaseURL, "https://login.ah.nl")
	target, err := url.Parse(loginBaseURL)
	if err != nil {
		listener.Close()
		return fmt.Errorf("invalid login URL: %w", err)
	}

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.Host = target.Host
			req.Header.Del("Accept-Encoding")
		},
		ModifyResponse: func(resp *http.Response) error {
			return rewriteLoginResponse(resp, localOrigin, target.Host)
		},
	}
	mux.Handle("/", proxy)

	srv := &http.Server{Handler: mux}
	go srv.Serve(listener)
	defer srv.Shutdown(context.Background())

	loginURL := fmt.Sprintf("%s/login?client_id=%s&response_type=code&redirect_uri=appie://login-exit",
		localOrigin, c.clientID)

	if c.openBrowser != nil {
		c.openBrowser(loginURL)
	} else {
		openDefaultBrowser(loginURL)
	}

	select {
	case code := <-codeCh:
		return c.exchangeCode(ctx, code)
	case <-ctx.Done():
		return ctx.Err()
	}
}

func rewriteLoginResponse(resp *http.Response, localOrigin, targetHost string) error {
	// Intercept server-side redirects to appie://
	loc := resp.Header.Get("Location")
	if strings.HasPrefix(loc, "appie://") {
		u, _ := url.Parse(loc)
		resp.Header.Set("Location", fmt.Sprintf("%s/callback?%s", localOrigin, u.Query().Encode()))
		return nil
	}

	// Rewrite Location headers pointing to the login host
	if strings.Contains(loc, targetHost) {
		resp.Header.Set("Location", strings.ReplaceAll(loc, "https://"+targetHost, localOrigin))
	}

	// Strip security headers that would block the proxy
	resp.Header.Del("Content-Security-Policy")
	resp.Header.Del("Strict-Transport-Security")
	resp.Header.Del("X-Frame-Options")

	// Only rewrite text response bodies
	ct := resp.Header.Get("Content-Type")
	if !strings.Contains(ct, "text/html") && !strings.Contains(ct, "javascript") && !strings.Contains(ct, "json") {
		return nil
	}

	body, err := readResponseBody(resp)
	if err != nil {
		return err
	}

	body = bytes.ReplaceAll(body, []byte("appie://login-exit"), []byte(localOrigin+"/callback"))
	body = bytes.ReplaceAll(body, []byte("https://"+targetHost), []byte(localOrigin))

	resp.Body = io.NopCloser(bytes.NewReader(body))
	resp.ContentLength = int64(len(body))
	resp.Header.Del("Content-Encoding")

	return nil
}

func readResponseBody(resp *http.Response) ([]byte, error) {
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer gz.Close()
		reader = gz
	}
	data, err := io.ReadAll(reader)
	resp.Body.Close()
	return data, err
}

func openDefaultBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		fmt.Printf("Open this URL in your browser:\n\n  %s\n\n", url)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Printf("Open this URL in your browser:\n\n  %s\n\n", url)
	}
}
