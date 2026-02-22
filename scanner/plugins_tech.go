package scanner

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func init() {
	RegisterCheck("tech_stack", "Guesses the backend technology stack from HTTP headers and meta tags", checkTechPlugin)
}

func checkTechPlugin(ctx context.Context, url string) interface{} {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return map[string]string{"error": "Failed to create request"}
	}
	req.Header.Set("User-Agent", "URLHawkScanner-Scanner/1.0")

	resp, err := pluginClient.Do(req)
	if err != nil {
		return map[string]string{"error": "Failed to reach host"}
	}
	defer resp.Body.Close()

	stack := make(map[string]string)

	// Fingerprint headers
	if server := resp.Header.Get("Server"); server != "" {
		stack["Server"] = server
	}
	if poweredBy := resp.Header.Get("X-Powered-By"); poweredBy != "" {
		stack["Powered-By"] = poweredBy
	}
	if asp := resp.Header.Get("X-AspNet-Version"); asp != "" {
		stack["ASP.NET"] = asp
	}

	// Fingerprint basic body HTML tags (first 2kb is enough)
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
	content := strings.ToLower(string(body))

	if strings.Contains(content, "wp-content") || strings.Contains(content, "wordpress") {
		stack["CMS"] = "WordPress"
	} else if strings.Contains(content, "shopify") {
		stack["CMS"] = "Shopify"
	}

	if strings.Contains(content, "react") || strings.Contains(content, "data-reactroot") {
		stack["Frontend"] = "React"
	} else if strings.Contains(content, "ng-app") || strings.Contains(content, "angular") {
		stack["Frontend"] = "Angular"
	}

	// Format uniquely so it matches the other []string arrays for simple rendering
	var results []string
	for k, v := range stack {
		results = append(results, fmt.Sprintf("%s: %s", k, v))
	}

	return results
}
