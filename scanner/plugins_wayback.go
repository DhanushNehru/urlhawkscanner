package scanner

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	RegisterCheck("wayback_machine", "Checks for historical snapshots on archive.org", checkWaybackPlugin)
}

func checkWaybackPlugin(ctx context.Context, url string) interface{} {
	domain := extractDomain(url)
	if domain == "" {
		return map[string]string{"error": "Invalid domain"}
	}

	target := fmt.Sprintf("http://archive.org/wayback/available?url=%s", domain)
	req, err := http.NewRequestWithContext(ctx, "GET", target, nil)
	if err != nil {
		return map[string]string{"error": "Request failed"}
	}

	resp, err := pluginClient.Do(req)
	if err != nil {
		return map[string]string{"error": "Archive API unreachable"}
	}
	defer resp.Body.Close()

	var result struct {
		ArchivedSnapshots struct {
			Closest struct {
				Available bool   `json:"available"`
				Url       string `json:"url"`
				Timestamp string `json:"timestamp"`
			} `json:"closest"`
		} `json:"archived_snapshots"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return map[string]string{"error": "Failed to parse Archive data"}
	}

	if result.ArchivedSnapshots.Closest.Available {
		return map[string]string{
			"Archived": "Yes",
			"Latest Snapshot": result.ArchivedSnapshots.Closest.Url,
			"Timestamp": result.ArchivedSnapshots.Closest.Timestamp,
		}
	}

	return map[string]string{"Archived": "No snapshots found"}
}
