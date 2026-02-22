package scanner

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func init() {
	RegisterCheck("social_links", "Finds potential social media profiles linked on the homepage", checkSocialsPlugin)
}

func checkSocialsPlugin(ctx context.Context, url string) interface{} {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return []string{}
	}
	req.Header.Set("User-Agent", "URLHawkScanner/1.0")

	resp, err := pluginClient.Do(req)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 50000)) // limit to 50kb
	if err != nil {
		return []string{}
	}

	content := string(body)
	var links []string

	socialDomains := []string{"twitter.com", "github.com", "linkedin.com", "facebook.com", "instagram.com", "youtube.com"}

	lines := strings.Split(content, "href=\"")
	for _, line := range lines {
		if idx := strings.Index(line, "\""); idx > -1 {
			link := line[:idx]
			for _, domain := range socialDomains {
				if strings.Contains(link, domain) && !containsDomain(links, link) {
					links = append(links, link)
				}
			}
		}
	}

	if len(links) > 10 {
		links = links[:10]
	}

	return links
}

func containsDomain(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
