package scanner

import (
	"context"
	"net/http"
)

func init() {
	RegisterCheck("http_methods", "Identifies allowed HTTP methods", checkMethodsPlugin)
}

func checkMethodsPlugin(ctx context.Context, url string) interface{} {
	req, err := http.NewRequestWithContext(ctx, "OPTIONS", url, nil)
	if err != nil {
		return map[string]string{"error": "Request failed"}
	}

	resp, err := pluginClient.Do(req)
	if err != nil {
		return map[string]string{"error": "Host unreachable"}
	}
	defer resp.Body.Close()

	allow := resp.Header.Get("Allow")
	if allow == "" {
		return map[string]string{"Allowed Methods": "Not explicitly defined (No Allow header)"}
	}
	return map[string]string{"Allowed Methods": allow}
}
