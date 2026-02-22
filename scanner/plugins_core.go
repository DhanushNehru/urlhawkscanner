package scanner

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"sync"
)

var (
	pluginClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
		},
		// Timeouts are handled by the Context passed to each plugin
	}

	sensitivePaths = []string{
		"/.env",
		"/.git/config",
		"/docker-compose.yml",
		"/backup.sql",
	}
)

func init() {
	RegisterCheck("missing_headers", "Checks for missing critical security headers", checkHeadersPlugin)
	RegisterCheck("exposed_files", "Checks for commonly exposed sensitive files", checkSensitiveFilesPlugin)
}

func checkHeadersPlugin(ctx context.Context, url string) interface{} {
	var missing []string
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return missing
	}
	req.Header.Set("User-Agent", "URLHawkScanner-Scanner/1.0")

	resp, err := pluginClient.Do(req)
	if err != nil {
		return append(missing, "Host Unreachable")
	}
	defer resp.Body.Close()

	if resp.Header.Get("X-Frame-Options") == "" {
		missing = append(missing, "X-Frame-Options")
	}
	if resp.Header.Get("Content-Security-Policy") == "" {
		missing = append(missing, "Content-Security-Policy")
	}
	if resp.Header.Get("Strict-Transport-Security") == "" {
		missing = append(missing, "Strict-Transport-Security")
	}
	return missing
}

func checkSensitiveFilesPlugin(ctx context.Context, baseURL string) interface{} {
	var exposed []string
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, path := range sensitivePaths {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			target := baseURL + p
			req, err := http.NewRequestWithContext(ctx, "GET", target, nil)
			if err != nil {
				return
			}
			req.Header.Set("User-Agent", "URLHawkScanner-Scanner/1.0")

			resp, err := pluginClient.Do(req)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == 200 {
				body, err := io.ReadAll(io.LimitReader(resp.Body, 512))
				if err != nil {
					return
				}

				content := string(body)
				isFalsePositive := strings.Contains(strings.ToLower(content), "<html") || strings.Contains(strings.ToLower(content), "<body")

				if !isFalsePositive {
					mu.Lock()
					exposed = append(exposed, p)
					mu.Unlock()
				}
			}
		}(path)
	}

	wg.Wait()
	return exposed
}
