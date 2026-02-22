<h1 align="center">
  URLHawkScan 🦅
</h1>

<p align="center">
  A blazing-fast, concurrent URL vulnerability and misconfiguration scanner built in Go.
</p>

<p align="center">
  <a href="https://github.com/DhanushNehru/urlhawk/releases"><img src="https://img.shields.io/github/v/release/DhanushNehru/urlhawk" alt="Release"></a>
  <a href="https://github.com/DhanushNehru/urlhawk/blob/main/LICENSE"><img src="https://img.shields.io/github/license/DhanushNehru/urlhawk" alt="License"></a>
  <a href="https://golang.org"><img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="Made with Go"></a>
</p>

## Overview

**URLHawk** is designed to quickly identify low-hanging fruit and common misconfigurations across large lists of URLs or single targets. It is built for speed and simplicity.

It currently checks for:
- Missing Security Headers (`X-Frame-Options`, `Content-Security-Policy`, `Strict-Transport-Security`).
- Exposed Sensitive Files (`.env`, `.git/config`, `docker-compose.yml`, `backup.sql`).

## ⚡ Features
- **Extremely Fast:** Uses Go concurrency routines (Goroutines) to scan massive lists of URLs asynchronously.
- **Colorized Output:** Easily distinguish between info, warnings, and critical findings with `fatih/color`.
- **Lightweight:** Minimal dependencies, compiles to a single, portable binary.
- **Customizable Threads:** Fine-tune your scanning speed with the `-t` (threads) flag.

---

## 🚀 Installation

Ensure you have [Go](https://go.dev/dl/) installed on your system.

### Build From Source
```bash
git clone https://github.com/DhanushNehru/urlhawkscan.git
cd urlhawkscan
go build -o urlhawkscan
sudo mv urlhawkscan /usr/local/bin/
```

### Or Install via `go install`
```bash
go install github.com/DhanushNehru/urlhawkscan@latest
```

---

## 🛠️ Usage

### Scan a single URL
```bash
urlhawkscan -u https://example.com
```

### Scan a list of URLs
```bash
urlhawkscan -l urls.txt
```

### Adjust Concurrency (Threads)
Speed up your scans over a huge list by increasing the thread count (default is 10).
```bash
urlhawkscan -l urls.txt -t 50
```

---

## 📸 Screenshots
*(Add a cool terminal screenshot of URLHawkScan running here)*

---

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## ⚠️ Disclaimer

URLHawkScan is created for educational and security assessment purposes only. The authors take no responsibility and are not liable for any misuse or damage caused by this tool. Only use URLHawkScan on authorized networks and domains.

---

<p align="center">
  Developed with ❤️ by <a href="https://github.com/DhanushNehru">Dhanush Nehru</a>
</p>
