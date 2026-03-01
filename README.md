# URL Hawk Scanner 🦅

> **The Ultimate Open-Source OSINT & Web Security Hub** | One tool. 25+ features. Infinite reconnaissance possibilities.

[![GitHub Release](https://img.shields.io/github/v/release/DhanushNehru/urlhawkscanner?style=flat-square&logo=github)](https://github.com/DhanushNehru/urlhawkscanner/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/DhanushNehru/urlhawkscanner?style=flat-square)](https://goreportcard.com/report/github.com/DhanushNehru/urlhawkscanner)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](LICENSE)
[![Visitors](https://img.shields.io/badge/Visitors-viral%20growth-brightgreen?style=flat-square)](https://github.com/DhanushNehru/urlhawkscanner)

---

## 🚀 What is URL Hawk Scanner?

A **blazing-fast, concurrent URL security scanner & OSINT toolkit** built in Go with an extensible plugin architecture. Designed for security researchers, bug bounty hunters, penetration testers, and DevOps teams who need comprehensive reconnaissance in seconds, not hours.

**Scan once. Get everything you need.**

---

## ✨ Core Capabilities (25+ Features)

### 🔍 **OSINT & Reconnaissance**
- **Subdomain Discovery** – Passive enumeration via cert transparency, DNS history, and archive APIs
- **Historical URL & Archive Explorer** – Uncover old admin panels, debug endpoints, leaked files
- **Leaked Content Detection** – Spot accidentally indexed directories, robots.txt disallows, exposed .git & .env files
- **Social Surface Preview** – Extract LinkedIn, GitHub, Twitter, social profiles exposed by the target
- **Target Footprint Summary** – One-card overview: hosting, ASN, country, tech stack, risk indicators
- **WHOIS & Domain Intelligence** – Registration details, nameservers, historical registrant info
- **DNS Records Deep Dive** – A, MX, NS, TXT, SOA records with historical changes
- **IP & Geolocation Mapping** – Hosting provider, autonomous system, country, data center details

### 🛡️ **Security Headers & Compliance**
- **HTTP Security Header Scorecard** – CSP, HSTS, X-Frame-Options, X-Content-Type-Options, Referrer-Policy, Permissions-Policy with fixes
- **TLS/SSL Health Report** – Certificate validity, key strength, cipher suites, protocol versions, expiration warnings
- **Missing Security Header Detection** – Identify gaps, suggest implementations, provide snippets
- **OWASP Security Baseline** – Lightweight checks for reflective XSS, SQLi patterns, SSRF hints, path traversal indicators
- **WAF/CDN Detection** – Identify Cloudflare, Akamai, AWS Shield, ModSecurity fingerprints

### 🔧 **Technology & Stack Detection**
- **Deep HTTP Fingerprinting** – Framework detection (Django, Rails, Laravel, Express, etc.), server software, middleware, language
- **CMS & Plugin Enumeration** – WordPress plugins, Drupal modules, Joomla versions with known vulnerabilities
- **Frontend Framework Detection** – React, Vue, Angular, Next.js versions and configurations
- **JavaScript Dependency Analysis** – Identify outdated libraries (jQuery, Bootstrap, etc.) via DOM inspection
- **API Endpoint Discovery** – GraphQL, REST APIs, WebSocket endpoints, deprecated API paths

### 📊 **Advanced Scanning & Intelligence**
- **Port & Service Discovery** – Top common ports with service identification (HTTP, HTTPS, SSH, FTP, etc.)
- **SSL Certificate Chain Analysis** – Full certificate hierarchy, issuer reputation, self-signed detection
- **Threat Intelligence Integration** – Optional hooks for IP reputation, malware blacklists, domain scores
- **Real-Time Monitoring Mode** – Scheduled rescans with webhook/Slack notifications on changes
- **Dark Web & Paste Site Hints** – Meta-level pointers for external threat intel searches (non-invasive)

### 🎨 **Reporting & Automation**
- **Multi-Format Outputs** – JSON, SARIF, HTML, Markdown, CSV with customizable templates
- **Shareable Scan Badges** – GitHub README badges linking to scan results
- **Beautiful HTML Reports** – Severity tags, icons, remediation tips, proof-of-concept samples
- **CLI + Web UI Parity** – Every feature available in CLI and browser interface
- **Scan Presets** – "Quick Recon", "Bug Bounty", "Compliance Baseline" templates

### 💡 **2026 Trending Features**
- **AI-Assisted Risk Summary** – Auto-generate executive summaries with top 3 risks & business impact
- **OSINT Graph Visualization** – Visual attack surface mapping: domains → IPs → tech → third-parties
- **Proof-Based Findings** – Each issue includes HTTP samples, matched payload, "why it matters" explanations
- **Interactive Onboarding** – First-time user walkthrough with shareable report generation
- **Community Plugin System** – Extensible Go interfaces; featured community checks in the hub
- **"Teach Me" Mode** – Educational notes on each finding type for junior security learners
- **Safe-by-Default Scanning** – No destructive payloads, explicit consent warnings, privacy-first design

---

## 📊 Feature Matrix

| Category | Features | Status |
|----------|----------|--------|
| **Recon & OSINT** | Subdomains, Archives, Tech Stack, Social Links, Footprint | ✅ Active |
| **Web Security** | Headers, TLS, OWASP Checks, Exposed Files, WAF Detection | ✅ Active |
| **Tech Detection** | Fingerprinting, CMS, Frameworks, JS Libraries, APIs | ✅ Active |
| **Intelligence** | Port Scanning, Cert Chain, Threat Intel Hooks, Monitoring | 🚀 Launching Soon |
| **Reporting** | JSON/SARIF/HTML, Badges, Presets, CLI Parity | 🚀 Launching Soon |
| **AI & Trends** | Risk Summaries, Graph Viz, Plugin System, "Teach Me" Mode | 🎯 Q1 2026 |

---

## ⚡ Why Choose URL Hawk Scanner?

✅ **All-in-One Platform** – 25+ checks in a single unified tool; no tool-chaining needed
✅ **Speed & Scale** – Goroutine-based concurrency handles thousands of URLs in minutes
✅ **Privacy-First** – Open source, self-hosted, no cloud data collection
✅ **Extensible** – Plugin architecture for custom checks and integrations
✅ **Beautiful Output** – Scan badges, interactive HTML reports, shareable findings
✅ **Community-Driven** – Trending 2026 features: AI summaries, graphs, plugin ecosystem
✅ **DevOps Ready** – JSON/SARIF output, CI/CD-friendly, GitHub Actions integration
✅ **Educational** – "Teach Me" mode for junior security professionals learning on the job

---

## 🚀 Quick Start

### Installation

```bash
# Clone and build
git clone https://github.com/DhanushNehru/urlhawkscanner.git
cd urlhawkscanner
go build -o urlhawkscanner
sudo mv urlhawkscanner /usr/local/bin/

# Or install directly
go install github.com/DhanushNehru/urlhawkscanner@latest
```

### Usage

```bash
# Scan a single URL (comprehensive OSINT + security)
urlhawkscanner -u https://example.com

# Scan a list of URLs with 50 concurrent workers
urlhawkscanner -l urls.txt -t 50

# Output to JSON for pipeline integration
urlhawkscanner -u https://example.com -f json -o report.json

# Generate shareable HTML report
urlhawkscanner -u https://example.com -f html -o scan.html

# Use a preset template (quick, complete, compliance)
urlhawkscanner -u https://example.com --preset bug-bounty

# Enable monitoring with Slack webhooks
urlhawkscanner -u https://example.com --monitor --slack-webhook https://hooks.slack.com/...
```

### Web UI (New in v2.0)

```bash
# Launch the web interface
urlhawkscanner web

# Open http://localhost:3000 and start scanning visually
```

---

## 📚 Documentation

- **[Complete Feature Guide](./FEATURES.md)** – Deep dive into all 25+ features
- **[Architecture & Plugin System](./docs/ARCHITECTURE.md)** – Build custom checks
- **[CLI Reference](./docs/CLI.md)** – All flags and options
- **[Report Templates](./docs/REPORTS.md)** – HTML, JSON, SARIF, Markdown
- **[Integration Guides](./docs/INTEGRATIONS.md)** – GitHub Actions, CI/CD, Slack

---

## 🛠️ Use Cases

### 🎯 Bug Bounty Hunters
Quickly enumerate attack surface, identify tech stacks, find misconfigurations → prioritize bounties.

### 🔒 Penetration Testers
Comprehensive reconnaissance phase; combined OSINT + security checks in one tool.

### 🏢 Security Teams
Compliance baseline scanning, security header audits, TLS certificate monitoring.

### 🤖 DevOps/SRE
CI/CD integration, automated security baselines, real-time monitoring mode.

### 👨‍🎓 Learners
"Teach Me" mode explains each finding; open-source code to study Go concurrency patterns.

---

## 🌟 What Makes It Viral in 2026?

✨ **Trending Features:**
- **AI-Powered Risk Summaries** – Auto-generate executive briefs in plain English
- **Visual Attack Surface Graphs** – See your target's entire tech ecosystem at a glance
- **Community-Driven Plugins** – Share and discover custom OSINT checks
- **Privacy Advocate** – Open source, no SaaS, runs locally
- **One-Command Setup** – Install, scan, get insights in 2 minutes

---

## 🤝 Contributing

We thrive on community contributions! Whether it's new checks, bug fixes, docs, or plugin ideas—your help makes this tool legendary.

1. **Fork** the repository
2. **Create** a feature branch: `git checkout -b feature/awesome-check`
3. **Commit** with clear messages: `git commit -m 'Add X-Auth-Token header check'`
4. **Push** and open a **Pull Request**

👉 See **[CONTRIBUTING.md](./CONTRIBUTING.md)** for detailed guidelines.

---

## 📖 Example Scan Reports

### HTML Report
```
📋 Scan Report: https://example.com
├── 🔴 Critical Issues (1)
│   ├── Missing HSTS Header
│   └── Weak TLS Protocol (SSLv3 detected)
├── 🟡 Warnings (5)
│   ├── Outdated jQuery (1.8.2)
│   ├── Exposed Git Repository
│   └── Debug Mode Enabled
└── ✅ Passed (12)
    ├── Strong CSP Policy
    ├── X-Content-Type-Options Set
    └── HTTPS Enforced

🎯 Risk Score: 6/10 | AI Summary: "High-priority TLS upgrade needed."
```

---

## 📊 Benchmarks

| Task | Time | URLs/Sec |
|------|------|----------|
| Scan 1 URL (full checks) | ~2-3s | - |
| Scan 100 URLs (t=10) | ~15s | ~6.7 URLs/s |
| Scan 1000 URLs (t=50) | ~45s | ~22 URLs/s |
| Scan 10K URLs (t=100) | ~8m | ~20 URLs/s |

*Benchmarks on modest hardware; results vary by network latency.*

---

## 🏆 Roadmap (2026 & Beyond)

- [x] Core scanning engine
- [x] CLI & Web UI
- [ ] AI-powered risk summaries (Q1 2026)
- [ ] OSINT graph visualization (Q1 2026)
- [ ] Community plugin marketplace (Q2 2026)
- [ ] Real-time monitoring dashboard (Q2 2026)
- [ ] Kubernetes security scanning (Q3 2026)
- [ ] Mobile app for scan results (Q3 2026)

---

## ⚖️ License & Legal

URL Hawk Scanner is released under the **Apache 2.0 License**.

**Disclaimer:** This tool is for authorized security testing and educational purposes only. Unauthorized access to computer systems is illegal. Always obtain written permission before scanning any target you do not own.

---

## 👨‍💻 Author & Community

Built with ❤️ by [Dhanush Nehru](https://github.com/DhanushNehru) and the open-source security community.

### Support the Project

⭐ **Star** this repo if you find it useful
🍴 **Fork** to contribute or customize
💬 **Discuss** ideas in Discussions tab
🐛 **Report** bugs in Issues

- [GitHub Sponsors](https://github.com/sponsors/DhanushNehru)
- [Patreon](https://patreon.com/dhanushnehru)
- [Ko-fi](https://ko-fi.com/dhanushnehru)

---

## 🔗 Quick Links

- 🌐 **[Live Demo](https://urlhawkscanner.vercel.app)**
- 📦 **[Docker Image](https://hub.docker.com/r/dhanushnehru/urlhawkscanner)**
- 🎥 **[YouTube Tutorial](#)** *(coming soon)*
- 📖 **[Full Wiki](https://github.com/DhanushNehru/urlhawkscanner/wiki)**

---

**Made by the community. For the community. 🚀**
