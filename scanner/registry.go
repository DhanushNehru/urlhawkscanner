package scanner

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

// CheckFunc is the signature for all OSINT plugin checks
type CheckFunc func(ctx context.Context, url string) interface{}

// CheckDefinition holds the metadata and execution logic for a check
type CheckDefinition struct {
	Name        string
	Description string
	Execute     CheckFunc
}

// registry holds all the registered plugins
var registry = make(map[string]CheckDefinition)

// RegisterCheck allows a plugin to register itself in its init() func
func RegisterCheck(name, description string, check CheckFunc) {
	registry[name] = CheckDefinition{
		Name:        name,
		Description: description,
		Execute:     check,
	}
}

// RunAllChecks executes all registered plugins concurrently with Vercel safety guarantees
func RunAllChecks(url string) map[string]interface{} {
	results := make(map[string]interface{})
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Vercel Serverless Safety Guarantee 1: Strict Global Timeout
	// The Vercel Free Tier hard limits at 10 seconds. We use 8 seconds to allow JSON serialization time.
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	for key, check := range registry {
		wg.Add(1)
		go func(k string, chk CheckDefinition) {
			defer wg.Done()

			// Vercel Serverless Safety Guarantee 2: Graceful Panic Recovery
			defer func() {
				if r := recover(); r != nil {
					color.Red("[-] Plugin Panic (%s): %v", k, r)
					mu.Lock()
					results[k] = map[string]string{"error": fmt.Sprintf("Plugin execution crashed: %v", r)}
					mu.Unlock()
				}
			}()

			// Execute the check, passing the context down so plugins can abort network calls if time runs out
			res := chk.Execute(ctx, url)

			mu.Lock()
			results[k] = res
			mu.Unlock()
		}(key, check)
	}

	// Wait for all plugins to finish OR the global timeout context to expire inside the plugins
	wg.Wait()

	return results
}
