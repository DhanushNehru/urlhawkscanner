package scanner

import (
	"fmt"
	"strings"
	"sync"

	"github.com/fatih/color"
)

var (
	printMu sync.Mutex
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
		color.Blue("\n[~] Scanning %s", url)
		printMu.Unlock()

		// Run all dynamic plugins
		results := RunAllChecks(url)

		printMu.Lock()
		for key, checkFunc := range registry {
			res := results[key]

			// Generic print logic based on data type returned
			switch v := res.(type) {
			case []string:
				if len(v) > 0 {
					color.Yellow("    [!] %s: %s", checkFunc.Name, strings.Join(v, ", "))
				}
			case string:
				if v != "" {
					color.Cyan("    [i] %s: %s", checkFunc.Name, v)
				}
			case map[string]string:
				if len(v) > 0 {
					color.Red("    [x] %s: Data Blocked or Errored", checkFunc.Name)
				}
			default:
				// other types ignored in simple CLI printout for now, or print generic
				if v != nil {
					color.White("    [-] %s: Data Found", checkFunc.Name)
				}
			}
		}
		printMu.Unlock()
	}
}

func normalizeURL(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url // Default to http, redirect might happen
	}
	return strings.TrimRight(url, "/")
}
