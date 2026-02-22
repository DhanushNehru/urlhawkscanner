package scanner

import (
	"context"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func init() {
	RegisterCheck("whois_info", "Fetches domain registration data like Registrar and Expiry Dates", checkWhoisPlugin)
}

func checkWhoisPlugin(ctx context.Context, url string) interface{} {
	domain := extractDomain(url)
	if domain == "" {
		return map[string]string{"error": "Invalid domain"}
	}

	// Because likexian/whois doesn't natively take a context, we run it in a goroutine
	// and use select to enforce the context timeout.
	type whoisResult struct {
		data interface{}
		err  error
	}

	resultChan := make(chan whoisResult, 1)

	go func() {
		raw, err := whois.Whois(domain)
		if err != nil {
			resultChan <- whoisResult{nil, err}
			return
		}

		parsed, err := whoisparser.Parse(raw)
		if err != nil {
			resultChan <- whoisResult{nil, err}
			return
		}

		results := make(map[string]interface{})
		if parsed.Registrar != nil {
			results["Registrar"] = parsed.Registrar.Name
		}
		if parsed.Domain != nil {
			results["Created"] = parsed.Domain.CreatedDate
			results["Expires"] = parsed.Domain.ExpirationDate
			results["Status"] = parsed.Domain.Status
		}

		resultChan <- whoisResult{results, nil}
	}()

	select {
	case <-ctx.Done():
		return map[string]string{"error": "Lookup timed out"}
	case res := <-resultChan:
		if res.err != nil {
			return map[string]string{"error": "Failed to parse whois data"}
		}
		return res.data
	}
}
