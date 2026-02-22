package scanner

import (
	"context"
	"net"
	"strings"
	"sync"
)

func init() {
	RegisterCheck("dns_records", "Retrieves A, AAAA, MX, NS, and TXT records", checkDNSPlugin)
}

func checkDNSPlugin(ctx context.Context, url string) interface{} {
	domain := extractDomain(url)
	if domain == "" {
		return map[string]string{"error": "Invalid domain"}
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make(map[string]interface{})

	// Custom resolver to respect context timeouts on DNS dial
	var resolver net.Resolver

	// A Records
	wg.Add(1)
	go func() {
		defer wg.Done()
		ips, err := resolver.LookupIPAddr(ctx, domain)
		if err == nil {
			var ipStrs []string
			for _, ip := range ips {
				ipStrs = append(ipStrs, ip.String())
			}
			mu.Lock()
			results["A/AAAA"] = ipStrs
			mu.Unlock()
		}
	}()

	// MX Records
	wg.Add(1)
	go func() {
		defer wg.Done()
		mxs, err := resolver.LookupMX(ctx, domain)
		if err == nil {
			var mxStrs []string
			for _, mx := range mxs {
				mxStrs = append(mxStrs, mx.Host)
			}
			mu.Lock()
			results["MX"] = mxStrs
			mu.Unlock()
		}
	}()

	// NS Records
	wg.Add(1)
	go func() {
		defer wg.Done()
		nss, err := resolver.LookupNS(ctx, domain)
		if err == nil {
			var nsStrs []string
			for _, ns := range nss {
				nsStrs = append(nsStrs, ns.Host)
			}
			mu.Lock()
			results["NS"] = nsStrs
			mu.Unlock()
		}
	}()

	// TXT Records
	wg.Add(1)
	go func() {
		defer wg.Done()
		txts, err := resolver.LookupTXT(ctx, domain)
		if err == nil {
			mu.Lock()
			results["TXT"] = txts
			mu.Unlock()
		}
	}()

	wg.Wait()
	return results
}

func extractDomain(url string) string {
	domain := strings.TrimPrefix(url, "http://")
	domain = strings.TrimPrefix(domain, "https://")
	parts := strings.Split(domain, "/")
	if len(parts) > 0 {
		// remove port if it exists just for safety on DNS lookups
		return strings.Split(parts[0], ":")[0]
	}
	return ""
}
