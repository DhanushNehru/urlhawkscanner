package scanner

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func init() {
	RegisterCheck("robots_txt", "Extracts disallowed or hidden paths from robots.txt", checkRobotsPlugin)
}

func checkRobotsPlugin(ctx context.Context, baseURL string) interface{} {
	target := baseURL + "/robots.txt"
	req, err := http.NewRequestWithContext(ctx, "GET", target, nil)
	if err != nil {
		return []string{}
	}
	req.Header.Set("User-Agent", "URLHawkScanner-Scanner/1.0")

	resp, err := pluginClient.Do(req)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []string{}
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 10240)) // limit to 10kb
	if err != nil {
		return []string{}
	}

	content := string(body)
	if strings.Contains(strings.ToLower(content), "<html") {
		return []string{} // False positive redirect to home page
	}

	var disallowed []string
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(strings.ToLower(line), "disallow:") {
			path := strings.TrimSpace(strings.TrimPrefix(strings.ToLower(line), "disallow:"))
			if path != "" && path != "/" {
				disallowed = append(disallowed, path)
			}
		}
	}

	// Limit to top 15 so we don't blow up the UI if the file is massive
	if len(disallowed) > 15 {
		disallowed = disallowed[:15]
		disallowed = append(disallowed, "... (more hidden paths found in file)")
	}

	return disallowed
}
