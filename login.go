package appie

import (
	"bytes"
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

	fhttp "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

const loginSuccessPage = `<!DOCTYPE html>
<html><head><title>Login Successful</title></head>
<body style="font-family:system-ui;max-width:500px;margin:80px auto;text-align:center">
<h1>Login successful!</h1>
<p>You can close this tab.</p>
<script>setTimeout(function(){window.close()},500)</script>
</body></html>`

const akamaiReloadScript = `<script>if(!sessionStorage.getItem("ah_akamai_reload")){sessionStorage.setItem("ah_akamai_reload","1");location.reload();}</script>`

// Login performs the full browser-based login flow. It starts a local reverse
// proxy to the AH login page, opens the user's browser, waits for the
// authorization code callback, and exchanges it for tokens.
//
// The proxy uses a Chrome TLS fingerprint to bypass AH's WAF, rewrites
// appie:// redirect URLs in responses, and intercepts the OAuth callback
// server-side so the browser never navigates to the custom scheme.
//
// Cancel the context to abort the login flow.
func (c *Client) Login(ctx context.Context) error {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return fmt.Errorf("failed to start login server: %w", err)
	}

	localOrigin := fmt.Sprintf("http://%s", listener.Addr())
	codeCh := make(chan string, 1)
	doneCh := make(chan error, 1)

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		c.logger.Printf("callback received, code length=%d", len(code))
		codeCh <- code
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, loginSuccessPage)
	})

	target, err := url.Parse(c.loginBaseURL)
	if err != nil {
		listener.Close()
		return fmt.Errorf("invalid login URL: %w", err)
	}

	transport, err := newLoginTransport(target.Scheme)
	if err != nil {
		listener.Close()
		return fmt.Errorf("failed to create login transport: %w", err)
	}

	proxy := &httputil.ReverseProxy{
		Transport: transport,
		Director: func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.Host = target.Host
			req.Header.Del("Accept-Encoding")
			c.logger.Printf("proxy >> %s %s", req.Method, req.URL.Path)
		},
		ModifyResponse: func(resp *http.Response) error {
			path := ""
			if resp.Request != nil && resp.Request.URL != nil {
				path = resp.Request.URL.Path
			}
			c.logger.Printf("proxy << %d %s (%s)", resp.StatusCode, path, resp.Header.Get("Content-Type"))
			if loc := resp.Header.Get("Location"); loc != "" {
				c.logger.Printf("proxy << Location: %s", loc)
			}

			if code, ok := appieRedirectCode(resp); ok {
				if err := c.exchangeCode(ctx, code); err != nil {
					c.logger.Printf("token exchange failed: %v", err)
					replaceWithLoginPage(resp, loginErrorPage(err))
					doneCh <- err
					return nil
				}
				replaceWithLoginPage(resp, loginSuccessPage)
				doneCh <- nil
				return nil
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

	loginURL := fmt.Sprintf("%s/login?client_id=%s&response_type=code&redirect_uri=appie://login-exit",
		localOrigin, c.clientID)

	if c.openBrowser != nil {
		c.openBrowser(loginURL)
	} else {
		openDefaultBrowser(loginURL)
	}

	select {
	case err := <-doneCh:
		return err
	case code := <-codeCh:
		return c.exchangeCode(ctx, code)
	case <-ctx.Done():
		return ctx.Err()
	}
}

func loginErrorPage(err error) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html><head><title>Login Failed</title></head>
<body style="font-family:system-ui;max-width:500px;margin:80px auto;text-align:center">
<h1>Login failed</h1>
<p>%s</p>
</body></html>`, htmlEscape(err.Error()))
}

func htmlEscape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

func appieRedirectCode(resp *http.Response) (string, bool) {
	loc := resp.Header.Get("Location")
	if !strings.HasPrefix(loc, "appie://") {
		return "", false
	}
	u, err := url.Parse(loc)
	if err != nil {
		return "", false
	}
	code := u.Query().Get("code")
	return code, code != ""
}

func replaceWithLoginPage(resp *http.Response, html string) {
	resp.StatusCode = http.StatusOK
	resp.Status = "200 OK"
	resp.Header.Del("Location")
	resp.Header.Set("Content-Type", "text/html; charset=utf-8")
	resp.Body = io.NopCloser(strings.NewReader(html))
	resp.ContentLength = int64(len(html))
	resp.Header.Set("Content-Length", fmt.Sprintf("%d", len(html)))
	resp.Header.Del("Content-Encoding")
	resp.Header.Del("Transfer-Encoding")
}

func newLoginTransport(scheme string) (http.RoundTripper, error) {
	if scheme == "http" {
		return http.DefaultTransport, nil
	}

	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(60),
		tls_client.WithClientProfile(profiles.Chrome_144),
		tls_client.WithNotFollowRedirects(),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		return nil, err
	}
	return &tlsRoundTripper{client: client}, nil
}

type tlsRoundTripper struct {
	client tls_client.HttpClient
}

func (t *tlsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fReq, err := toFhttpRequest(req)
	if err != nil {
		return nil, err
	}
	fResp, err := t.client.Do(fReq)
	if err != nil {
		return nil, err
	}
	return toNetHTTPResponse(fResp), nil
}

func toFhttpRequest(req *http.Request) (*fhttp.Request, error) {
	fReq, err := fhttp.NewRequestWithContext(req.Context(), req.Method, req.URL.String(), req.Body)
	if err != nil {
		return nil, err
	}
	fReq.ContentLength = req.ContentLength
	fReq.Host = req.Host
	fReq.Header = make(fhttp.Header, len(req.Header))
	for k, vv := range req.Header {
		for _, v := range vv {
			fReq.Header.Add(k, v)
		}
	}
	if req.Header.Get("Host") != "" {
		fReq.Header.Set("Host", req.Header.Get("Host"))
	}
	return fReq, nil
}

func toNetHTTPResponse(fResp *fhttp.Response) *http.Response {
	resp := &http.Response{
		Status:        fResp.Status,
		StatusCode:    fResp.StatusCode,
		Proto:         fResp.Proto,
		ProtoMajor:    fResp.ProtoMajor,
		ProtoMinor:    fResp.ProtoMinor,
		Header:        make(http.Header, len(fResp.Header)),
		ContentLength: fResp.ContentLength,
	}
	for k, vv := range fResp.Header {
		if isFhttpInternalHeader(k) {
			continue
		}
		for _, v := range vv {
			resp.Header.Add(k, v)
		}
	}
	resp.Body = fResp.Body
	return resp
}

func isFhttpInternalHeader(k string) bool {
	return k == fhttp.HeaderOrderKey || k == fhttp.PHeaderOrderKey
}

func rewriteLoginResponse(resp *http.Response, localOrigin, targetHost string) error {
	// Rewrite Location headers pointing to the login host
	loc := resp.Header.Get("Location")
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

	if strings.Contains(ct, "text/html") {
		body = injectAkamaiReload(body)
	}

	resp.Body = io.NopCloser(bytes.NewReader(body))
	resp.ContentLength = int64(len(body))
	resp.Header.Set("Content-Length", fmt.Sprintf("%d", len(body)))
	resp.Header.Del("Content-Encoding")
	resp.Header.Del("Transfer-Encoding")

	return nil
}

func injectAkamaiReload(body []byte) []byte {
	s := string(body)
	if strings.Contains(s, "ah_akamai_reload") {
		return body
	}
	lower := strings.ToLower(s)
	for _, tag := range []string{"</head>", "</body>"} {
		if idx := strings.Index(lower, tag); idx >= 0 {
			return []byte(s[:idx] + akamaiReloadScript + s[idx:])
		}
	}
	return append(body, []byte(akamaiReloadScript)...)
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
	data, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.Header.Get("Content-Encoding") != "gzip" {
		return data, nil
	}

	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		// Upstream may have already decompressed while keeping the header.
		return data, nil
	}
	defer gz.Close()

	decompressed, err := io.ReadAll(gz)
	if err != nil {
		return data, nil
	}
	return decompressed, nil
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
