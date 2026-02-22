package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DhanushNehru/urlhawkscanner/scanner"
)

// Handler is the entrypoint for Vercel Serverless Functions
func Handler(w http.ResponseWriter, r *http.Request) {
	// CORS for external clients (if any)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	urlParam := r.URL.Query().Get("url")
	if urlParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing 'url' parameter"})
		return
	}

	// In a serverless environment, we execute the scan synchronously
	// Vercel free tier limits execution to 10 seconds, which should be fine for URL Hawk Scan
	result := scanner.API_ScanURL(urlParam)

	json.NewEncoder(w).Encode(result)
}
