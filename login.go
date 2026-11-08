package appie

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	appieRedirect  = "appie://login-exit"
	loginDonePath  = "/login-done"
	loginDoneGrace = 15 * time.Second
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
	doneURL := localOrigin + loginDonePath
	codeCh := make(chan string, 1)
	var codeTaken atomic.Bool
	pageServed := make(chan struct{})
	var pageServedOnce sync.Once
	serveSuccessPage := func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, loginSuccessPage)
		pageServedOnce.Do(func() { close(pageServed) })
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		c.logger.Printf("callback received, code length=%d", len(code))
		deliverCode(codeCh, code)
		serveSuccessPage(w)
	})
	mux.HandleFunc(loginDonePath, func(w http.ResponseWriter, r *http.Request) {
		serveSuccessPage(w)
	})

	installationID, err := newInstallationID()
	if err != nil {
		listener.Close()
		return fmt.Errorf("failed to create installation id: %w", err)
	}

	target, err := url.Parse(c.loginBaseURL)
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
			// AH only issues the app's authorization code when the authorize
			// request carries the app's installation id; without it the login
			// ends on the "already logged in" page.
			if req.URL.Path == "/secure/oauth/authorize" {
				req.Header.Set("x-fraud-detection-installation-id", installationID)
			}
			c.logger.Printf("proxy >> %s %s", req.Method, req.URL.Path)
		},
		ModifyResponse: func(resp *http.Response) error {
			c.logger.Printf("proxy << %d %s (%s)", resp.StatusCode, resp.Request.URL.Path, resp.Header.Get("Content-Type"))
			if loc := resp.Header.Get("Location"); loc != "" {
				c.logger.Printf("proxy << Location: %s", loc)
			}
			if codeTaken.Load() && resp.Request.URL.Path == "/login/ingelogd" {
				redirectTo(resp, doneURL)
				return nil
			}
			if strings.HasSuffix(resp.Request.URL.Path, "/ingelogd.json") {
				if err := takeNextRedirectCode(resp, codeCh, &codeTaken, doneURL); err != nil {
					return err
				}
			}
			return rewriteLoginResponse(resp, localOrigin, target.Host)
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			c.logger.Printf("proxy error: %s %s: %v", r.Method, r.URL.Path, err)
			http.Error(w, "proxy error", http.StatusBadGateway)
		},
	}
	mux.Handle("/", proxy)

	srv := &http.Server{Handler: mux}
	go srv.Serve(listener)
	defer srv.Shutdown(ctx)

	loginURL := fmt.Sprintf("%s/secure/oauth/authorize?client_id=%s&response_type=code&redirect_uri=%s",
		localOrigin, c.clientID, appieRedirect)

	if c.openBrowser != nil {
		c.openBrowser(loginURL)
	} else {
		openDefaultBrowser(loginURL)
	}

	select {
	case code := <-codeCh:
		if err := c.exchangeCode(ctx, code); err != nil {
			return err
		}
		// The code can be taken before the browser reaches the success page, so
		// keep the server up until that page has been served.
		select {
		case <-pageServed:
		case <-time.After(loginDoneGrace):
		case <-ctx.Done():
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// takeNextRedirectCode takes the authorization code from the Next.js
// __N_REDIRECT in ingelogd.json and sends the browser to doneURL. The login
// page ignores the redirect in the first ingelogd.json response, the only one
// that carries the code, and fetches the page again; so the code is taken from
// the response itself, and once it is taken every ingelogd.json response
// redirects to doneURL.
func takeNextRedirectCode(resp *http.Response, codeCh chan<- string, taken *atomic.Bool, doneURL string) error {
	body, err := readResponseBody(resp)
	if err != nil {
		return err
	}
	var page map[string]any
	if json.Unmarshal(body, &page) != nil {
		setResponseBody(resp, body)
		return nil
	}
	props, _ := page["pageProps"].(map[string]any)
	redirect, _ := props["__N_REDIRECT"].(string)
	if code := codeFromRedirect(redirect); code != "" {
		deliverCode(codeCh, code)
		taken.Store(true)
	}
	if !taken.Load() {
		setResponseBody(resp, body)
		return nil
	}
	page["pageProps"] = map[string]any{"__N_REDIRECT": doneURL, "__N_REDIRECT_STATUS": http.StatusFound}
	rewritten, err := json.Marshal(page)
	if err != nil {
		return err
	}
	setResponseBody(resp, rewritten)
	return nil
}

// codeFromRedirect returns the code query parameter of an appie:// redirect, or "".
func codeFromRedirect(redirect string) string {
	if !strings.HasPrefix(redirect, appieRedirect) {
		return ""
	}
	u, err := url.Parse(redirect)
	if err != nil {
		return ""
	}
	return u.Query().Get("code")
}

// deliverCode passes the first code on; later ones are dropped, because AH can
// expose the same login more than once and only one exchange may run.
func deliverCode(codeCh chan<- string, code string) {
	if code == "" {
		return
	}
	select {
	case codeCh <- code:
	default:
	}
}

// redirectTo turns resp into an empty 302 to location.
func redirectTo(resp *http.Response, location string) {
	resp.Body.Close()
	resp.StatusCode = http.StatusFound
	resp.Status = "302 Found"
	resp.Header = http.Header{"Location": {location}}
	setResponseBody(resp, nil)
}

// setResponseBody replaces the (already decoded) body of resp.
func setResponseBody(resp *http.Response, body []byte) {
	resp.Header.Del("Content-Encoding")
	resp.Body = io.NopCloser(bytes.NewReader(body))
	resp.ContentLength = int64(len(body))
}

// newInstallationID returns a random uppercase UUID, the format of the app's
// installation id.
func newInstallationID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	b[6] = b[6]&0x0f | 0x40
	b[8] = b[8]&0x3f | 0x80
	return strings.ToUpper(fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])), nil
}

func rewriteLoginResponse(resp *http.Response, localOrigin, targetHost string) error {
	// Intercept server-side redirects to appie://
	loc := resp.Header.Get("Location")
	if strings.HasPrefix(loc, "appie://") {
		u, err := url.Parse(loc)
		if err != nil {
			return fmt.Errorf("failed to parse appie URL %q: %w", loc, err)
		}
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

	// Rewrite cookies: strip Secure/SameSite/Domain so they work over plain
	// HTTP on 127.0.0.1. Safari (unlike Chrome) does not treat localhost as a
	// secure context, so Secure cookies are silently dropped, breaking the
	// login session.
	if cookies := resp.Header.Values("Set-Cookie"); len(cookies) > 0 {
		resp.Header.Del("Set-Cookie")
		for _, c := range cookies {
			resp.Header.Add("Set-Cookie", sanitizeCookie(c))
		}
	}

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

// sanitizeCookie strips Secure, SameSite, and Domain attributes from a
// Set-Cookie header value so the cookie works over plain HTTP on localhost.
func sanitizeCookie(cookie string) string {
	parts := strings.Split(cookie, ";")
	out := parts[:1] // always keep the name=value part
	for _, p := range parts[1:] {
		attr := strings.TrimSpace(p)
		lower := strings.ToLower(attr)
		if lower == "secure" ||
			strings.HasPrefix(lower, "samesite") ||
			strings.HasPrefix(lower, "domain") {
			continue
		}
		out = append(out, p)
	}
	return strings.Join(out, ";")
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
	fmt.Println(url)
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return
	}
	_ = cmd.Start()
}
