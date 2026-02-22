package scanner

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

var commonPorts = []int{
	21,   // FTP
	22,   // SSH
	23,   // Telnet
	25,   // SMTP
	53,   // DNS
	80,   // HTTP
	110,  // POP3
	143,  // IMAP
	443,  // HTTPS
	445,  // SMB
	3306, // MySQL
	5432, // PostgreSQL
	6379, // Redis
	8080, // HTTP Alt
}

func init() {
	RegisterCheck("open_ports", "Scans common ports to see what services are exposed", checkPortsPlugin)
}

func checkPortsPlugin(ctx context.Context, url string) interface{} {
	domain := extractDomain(url)
	if domain == "" {
		return map[string]string{"error": "Invalid domain"}
	}

	var exposed []string
	var wg sync.WaitGroup
	var mu sync.Mutex

	dialer := &net.Dialer{}

	for _, port := range commonPorts {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			target := fmt.Sprintf("%s:%d", domain, p)

			// Vercel Egress Safety: Super fast 1-second timeout per port.
			// If it hangs here, the global context timeout will also catch it.
			timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
			defer cancel()

			conn, err := dialer.DialContext(timeoutCtx, "tcp", target)
			if err == nil {
				conn.Close()
				mu.Lock()
				exposed = append(exposed, fmt.Sprintf("%d", p))
				mu.Unlock()
			}
		}(port)
	}

	wg.Wait()
	return exposed
}
