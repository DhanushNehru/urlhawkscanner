package scanner

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Define JSON structures for the web API
type ScanResult struct {
	URL            string   `json:"url"`
	MissingHeaders []string `json:"missing_headers"`
	ExposedFiles   []string `json:"exposed_files"`
}

var (
	apiClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
		},
	}
	apiSensitivePaths = []string{
		"/.env",
		"/.git/config",
		"/docker-compose.yml",
		"/backup.sql",
	}
)

// API_ScanURL is a synchronous version of the scan tailored for returning data to the web UI
func API_ScanURL(url string) ScanResult {
	url = normalizeURL(url)

	result := ScanResult{
		URL:            url,
		MissingHeaders: []string{},
		ExposedFiles:   []string{},
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		result.MissingHeaders = apiCheckHeaders(url)
	}()

	go func() {
		defer wg.Done()
		result.ExposedFiles = apiCheckSensitiveFiles(url)
	}()

	wg.Wait()

	return result
}

func apiCheckHeaders(url string) []string {
	var missing []string
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return missing
	}
	req.Header.Set("User-Agent", "URLHawkScanner-Scanner/1.0")

	resp, err := apiClient.Do(req)
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

func apiCheckSensitiveFiles(baseURL string) []string {
	var exposed []string

	// Check concurrently for speed
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, path := range apiSensitivePaths {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			target := baseURL + p
			req, err := http.NewRequest("GET", target, nil)
			if err != nil {
				return
			}
			req.Header.Set("User-Agent", "URLHawkScanner-Scanner/1.0")

			resp, err := apiClient.Do(req)
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
