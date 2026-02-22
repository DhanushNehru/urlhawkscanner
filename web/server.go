package web

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/DhanushNehru/urlhawkscan/scanner"
	"github.com/fatih/color"
)

//go:embed static/*
var staticFiles embed.FS

func StartServer(port int) {
	// Serve static files from the embedded filesystem
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		color.Red("[-] Error loading embedded static files: %v", err)
		return
	}

	http.Handle("/", http.FileServer(http.FS(staticFS)))

	// API Endpoint for scanning
	http.HandleFunc("/api/scan", handleScan)

	addr := fmt.Sprintf(":%d", port)
	color.Cyan(`
  _   _  ____   _      _   _               _        _    _         _
 | | | ||  _ \ | |    | | | |  __ _ __   _| |      | |  | |  ___  | |__
 | | | || |_) || |    | |_| | / _' |\ \ / / |___   | |  | | / _ \ | '_ \
 | |_| ||  _ < | |___ |  _  || (_| | \ V /| |___|  | |__| ||  __/ | |_) |
  \___/ |_| \_\|_____||_| |_| \__,_|  \_/ |_|       \____/  \___| |_.__/

	`)
	color.Green("[+] URLHawkScan Web UI Started")
	color.Green("[+] Open your browser to http://localhost%s", addr)
	color.Yellow("[!] Press Ctrl+C to stop")

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		color.Red("[-] Server failed to start: %v", err)
	}
}

func handleScan(w http.ResponseWriter, r *http.Request) {
	// Enable CORS for potential separate frontend development
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	urlParam := r.URL.Query().Get("url")
	if urlParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing 'url' parameter"})
		return
	}

	color.Blue("[~] API Request received for: %s", urlParam)

	// Add an artificial small delay to make the UI look cool while scanning
	time.Sleep(800 * time.Millisecond)

	result := scanner.API_ScanURL(urlParam)

	json.NewEncoder(w).Encode(result)
}
