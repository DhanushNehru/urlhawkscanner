package scanner

// API_ScanURL is a synchronous version of the scan tailored for returning data to the web UI.
// It uses the dynamic plugin registry to run all OSINT checks concurrently.
func API_ScanURL(url string) map[string]interface{} {
	url = normalizeURL(url)

	// Run all registered plugins via the registry
	results := RunAllChecks(url)

	// Inject the target URL directly into the root so the frontend knows what was scanned
	results["url"] = url

	return results
}
