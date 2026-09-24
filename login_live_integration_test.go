//go:build integration

package appie

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestLoginProxyUpstream(t *testing.T) {
	client := New(WithLogger(log.New(os.Stderr, "", log.Ltime)))

	urlCh := make(chan string, 1)
	client.openBrowser = func(u string) { urlCh <- u }

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	errCh := make(chan error, 1)
	go func() { errCh <- client.Login(ctx) }()

	var proxyURL string
	select {
	case proxyURL = <-urlCh:
	case <-time.After(5 * time.Second):
		t.Fatal("login proxy did not start")
	}

	time.Sleep(500 * time.Millisecond)

	req, err := http.NewRequest(http.MethodGet, proxyURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("proxy fetch: %v", err)
	}
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 8192))
	resp.Body.Close()

	cancel()
	<-errCh

	if resp.StatusCode != http.StatusOK {
		prefix := string(body)
		if len(prefix) > 200 {
			prefix = prefix[:200]
		}
		t.Fatalf("expected 200, got %d body=%q", resp.StatusCode, prefix)
	}
	if strings.Contains(strings.ToLower(string(body)), "access denied") {
		prefix := string(body)
		if len(prefix) > 200 {
			prefix = prefix[:200]
		}
		t.Fatalf("got WAF block page: %q", prefix)
	}
}
