# Contributing to URL Hawk Scanner 🦅

**Thank you for your interest in contributing!** We welcome contributors of all levels—whether you're fixing bugs, adding features, improving docs, or spreading the word.

## 🚀 Quick Start (30 seconds)

1. **Fork** the repo
2. **Create** a feature branch: `git checkout -b feature/your-awesome-feature`
3. **Make changes** and test
4. **Push** and open a **Pull Request**

## 🎯 What Can You Contribute?

### 🐛 **Bug Fixes**
Found a bug? Open an issue with:
- Steps to reproduce
- Expected vs actual behavior
- Your environment (OS, Go version)

### ✨ **New Features**
Check our [ROADMAP.md](./ROADMAP.md) for trending 2026 features:
- **AI-powered risk summaries** (Q1 2026)
- **OSINT graph visualization** (Q1 2026)
- **Community plugin system** (Q2 2026)
- **Real-time monitoring** (Q2 2026)

### 📚 **Documentation**
- Improve READMEs
- Add tutorials
- Fix typos
- Create examples

### 🧪 **Tests**
Write unit tests, integration tests, or E2E tests. We need more coverage!

### 🎨 **Web UI / Frontend**
Improve the Vercel-deployed interface at [urlhawkscanner.vercel.app](https://urlhawkscanner.vercel.app)

### 🌟 **Community**
- Share the project on social media
- Write blog posts about URL Hawk Scanner
- Give talks at security meetups
- Answer questions in Discussions

---

## 🏗️ Development Setup

### Prerequisites
- Go 1.18+
- Node.js 16+ (for web UI)
- Git

### Local Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/urlhawkscanner.git
cd urlhawkscanner

# Install dependencies
go mod download

# Run the scanner
go run main.go -u https://example.com

# Run tests
go test ./...

# Setup web UI (optional)
cd web
npm install
npm run dev  # Open http://localhost:3000
```

---

## 📋 Code Style & Standards

### Go Code
- Follow [Effective Go](https://golang.org/doc/effective_go)
- Run `gofmt` and `golint` before committing
- Write clear function names and comments
- Add unit tests for new functionality

### Commit Messages
Use semantic commit format:
```
feat: Add X-Auth-Token header check
fix: Resolve nil pointer in DNS scanner
docs: Update plugin development guide
test: Improve coverage for TLS module
```

### PR Guidelines
- **Title**: Clear, descriptive (e.g., "feat: Add AI risk summary generator")
- **Description**: What, why, how. Link related issues.
- **Tests**: Include tests for new code
- **Docs**: Update README/docs if needed

---

## 🔌 Plugin Development (Advanced)

URL Hawk Scanner has an **extensible plugin architecture**. Create custom checks easily:

```go
package plugins

import "github.com/DhanushNehru/urlhawkscanner/scanner"

type MyCustomCheck struct{}

func (m *MyCustomCheck) Name() string {
    return "My Custom Check"
}

func (m *MyCustomCheck) Run(target string) (*scanner.Result, error) {
    // Your logic here
    return &scanner.Result{Finding: "..."}, nil
}
```

See [ARCHITECTURE.md](./docs/ARCHITECTURE.md) for full plugin guide.

---

## 💬 Communication

- **GitHub Issues**: Bug reports, feature requests
- **GitHub Discussions**: Questions, ideas, announcements
- **Pull Requests**: Code reviews, feedback
- **Social**: [@DhanushNehru](https://twitter.com/dhanushnehru)

---

## 🏆 Recognition

**Contributors will be:**
- Added to [CONTRIBUTORS.md](./CONTRIBUTORS.md)
- Mentioned in release notes
- Featured in social media shoutouts

---

## ⚠️ Code of Conduct

Be respectful, inclusive, and constructive. No harassment, discrimination, or spam.

---

## 📝 License

By contributing, you agree your code will be licensed under Apache 2.0.

---

## 🎓 Learning Resources

- [Go Concurrency Patterns](https://go.dev/blog/pipelines) (perfect for understanding URL Hawk's goroutines)
- [OWASP Testing Guide](https://owasp.org/www-project-web-security-testing-guide/)
- [Security Research](https://securitytrails.com/)

---

## 🤔 Questions?

Open a **Discussion** or tag us in an **Issue**. We're here to help!

**Happy contributing!** 🚀
