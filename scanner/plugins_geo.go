package scanner

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

func init() {
	RegisterCheck("geolocation", "Locates the server IP geographically using ip-api.com", checkGeoPlugin)
}

func checkGeoPlugin(ctx context.Context, url string) interface{} {
	domain := extractDomain(url)
	if domain == "" {
		return map[string]string{"error": "Invalid domain"}
	}

	var resolver net.Resolver
	ips, err := resolver.LookupIPAddr(ctx, domain)
	if err != nil || len(ips) == 0 {
		return map[string]string{"error": "Could not resolve IP"}
	}

	ip := ips[0].String()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://ip-api.com/json/%s", ip), nil)
	if err != nil {
		return map[string]string{"error": "Failed to create request"}
	}

	resp, err := pluginClient.Do(req)
	if err != nil {
		return map[string]string{"error": "Geo API unavailable"}
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return map[string]string{"error": "Failed to parse Geo data"}
	}

	if status, ok := result["status"].(string); ok && status == "success" {
		return map[string]string{
			"IP Address": ip,
			"Country": fmt.Sprintf("%v", result["country"]),
			"City": fmt.Sprintf("%v", result["city"]),
			"ISP": fmt.Sprintf("%v", result["isp"]),
		}
	}

	return map[string]string{"IP Address": ip, "Info": "Geo-location skipped or failed"}
}
