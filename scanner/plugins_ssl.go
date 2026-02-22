package scanner

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func init() {
	RegisterCheck("ssl_certificate", "Analyzes SSL/TLS certificate details", checkSSLPlugin)
}

func checkSSLPlugin(ctx context.Context, url string) interface{} {
	domain := extractDomain(url)
	if domain == "" {
		return map[string]string{"error": "Invalid domain"}
	}

	dialer := &net.Dialer{}

	// Create a context-aware dial
	deadline, ok := ctx.Deadline()
	if ok {
		dialer.Deadline = deadline
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", domain+":443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return map[string]string{"error": "No SSL/TLS on port 443 (or timed out)"}
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return map[string]string{"error": "No certificates found"}
	}

	cert := certs[0]

	return map[string]string{
		"Subject": cert.Subject.CommonName,
		"Issuer": cert.Issuer.CommonName,
		"Expires": cert.NotAfter.Format(time.RFC822),
		"Valid Now": fmt.Sprintf("%t", time.Now().Before(cert.NotAfter)),
		"Algorithm": cert.SignatureAlgorithm.String(),
	}
}
