package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/DhanushNehru/urlhawkscanner/scanner"
	"github.com/DhanushNehru/urlhawkscanner/web"
	"github.com/fatih/color"
)

var banner = `
 _   _  ____   _      _   _               _
| | | ||  _ \ | |    | | | |  __ _ __   _| | __
| | | || |_) || |    | |_| | / _' |\ \ / / |/ /
| |_| ||  _ < | |___ |  _  || (_| | \ V /|   <
 \___/ |_| \_\|_____||_| |_| \__,_|  \_/ |_|\_\

      A blazing fast URL open-source scanner
`

func printBanner() {
	color.Cyan(banner)
}

func main() {
	urlFlag := flag.String("u", "", "Single URL to scan")
	listFlag := flag.String("l", "", "File containing list of URLs to scan")
	threadsFlag := flag.Int("t", 10, "Number of concurrent threads")
	webFlag := flag.Bool("web", false, "Start the URLHawk web interface")
	portFlag := flag.Int("p", 8080, "Port for the web server (default 8080)")

	flag.Parse()

	if *webFlag {
		web.StartServer(*portFlag)
		return
	}

	printBanner()

	var urls []string

	if *urlFlag != "" {
		urls = append(urls, *urlFlag)
	} else if *listFlag != "" {
		file, err := os.Open(*listFlag)
		if err != nil {
			color.Red("[-] Error opening file: %v", err)
			os.Exit(1)
		}
		defer file.Close()
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			url := strings.TrimSpace(fileScanner.Text())
			if url != "" {
				// Ensure simple URLs have http:// if missing, basic normalization
				if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
					url = "http://" + url
				}
				urls = append(urls, url)
			}
		}
	} else {
		// Provide a default simple example
		color.Yellow("[-] No URLs provided. Provide either -u, -l, or -web")
		fmt.Println("Example CLI: ./urlhawkscanner -u example.com")
		fmt.Println("Example CLI: ./urlhawkscanner -l urls.txt -t 50")
		fmt.Println("Example Web: ./urlhawkscanner -web -p 8080")
		os.Exit(1)
	}

	color.Green("[+] Loaded %d URLs to scan", len(urls))
	color.Green("[+] Starting scan with %d threads...\n\n", *threadsFlag)

	scanner.RunScan(urls, *threadsFlag)
}
