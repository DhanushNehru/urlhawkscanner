package scanner

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

var (
	printMu sync.Mutex
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
		},
	}
	sensitivePaths = []string{
		"/.env",
		"/.git/config",
		"/docker-compose.yml",
		"/backup.sql",
	}
)

// RunScan is the entry point for the scanning engine
func RunScan(urls []string, threads int) {
	color.Cyan("[*] Engine initialized. Scanners warming up...")

	urlChan := make(chan string, len(urls))
	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(&wg, urlChan)
	}

	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan)

	wg.Wait()

	fmt.Println()
	color.Green("[+] Scan complete. Hawk is returning to nest.")
}

func worker(wg *sync.WaitGroup, urlChan <-chan string) {
	defer wg.Done()
	for unparsedURL := range urlChan {
		url := normalizeURL(unparsedURL)

		printMu.Lock()
		color.Blue("[~] Scanning %s", url)
		printMu.Unlock()

		checkHeaders(url)
		checkSensitiveFiles(url)
	}
}

func normalizeURL(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url // Default to http, redirect might happen
	}
	return strings.TrimRight(url, "/")
}

func checkHeaders(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "URLHawkScan-Scanner/1.0")

	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var missing []string
	if resp.Header.Get("X-Frame-Options") == "" {
		missing = append(missing, "X-Frame-Options")
	}
	if resp.Header.Get("Content-Security-Policy") == "" {
		missing = append(missing, "Content-Security-Policy")
	}
	if resp.Header.Get("Strict-Transport-Security") == "" {
		missing = append(missing, "Strict-Transport-Security")
	}

	if len(missing) > 0 {
		printMu.Lock()
		color.Yellow("    [!] Missing Headers: %s", strings.Join(missing, ", "))
		printMu.Unlock()
	}
}

func checkSensitiveFiles(baseURL string) {
	for _, path := range sensitivePaths {
		target := baseURL + path
		req, err := http.NewRequest("GET", target, nil)
		if err != nil {
			continue
		}
		req.Header.Set("User-Agent", "URLHawkScan-Scanner/1.0")

		resp, err := httpClient.Do(req)
		if err != nil {
			continue
		}

		if resp.StatusCode == 200 {
			body, err := io.ReadAll(io.LimitReader(resp.Body, 512))
			resp.Body.Close()
			if err != nil {
				continue
			}

			content := string(body)
			isFalsePositive := strings.Contains(strings.ToLower(content), "<html") || strings.Contains(strings.ToLower(content), "<body")

			if !isFalsePositive {
				printMu.Lock()
				color.Red("    [CRITICAL] Exposed file found: %s", target)
				printMu.Unlock()
			}
		} else {
			resp.Body.Close()
		}
	}
}
