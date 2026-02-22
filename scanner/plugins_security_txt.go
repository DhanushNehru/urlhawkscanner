package scanner

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func init() {
	RegisterCheck("security_txt", "Checks for a standard security.txt policy file", checkSecurityTxtPlugin)
}

func checkSecurityTxtPlugin(ctx context.Context, baseURL string) interface{} {
	target := baseURL + "/.well-known/security.txt"
	req, err := http.NewRequestWithContext(ctx, "GET", target, nil)
	if err != nil {
		return map[string]string{"Policy Found": "No"}
	}
	req.Header.Set("User-Agent", "URLHawkScanner/1.0")

	resp, err := pluginClient.Do(req)
	if err != nil {
		return map[string]string{"Policy Found": "Connection Failed"}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return map[string]string{"Policy Found": "No (404/Denied)"}
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 2048))
	if err != nil {
		return map[string]string{"Policy Found": "Read Error"}
	}
	content := string(body)

	if strings.Contains(strings.ToLower(content), "<html") {
		return map[string]string{"Policy Found": "No (Redirected to HTML)"}
	}

	return map[string]string{
		"Policy Found": "Yes",
		"Path": "/.well-known/security.txt",
		"Snippet": content[:min(len(content), 100)] + "...",
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
