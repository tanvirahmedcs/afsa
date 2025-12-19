# 🔴 AFSA - Advanced Forensic Security Analyzer

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue)](https://golang.org)
[![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux-blue)](https://github.com/tanvircs/afsa)

**A professional-grade, all-in-one command-line tool for advanced network and security reconnaissance.**

> Created by **Tanvir Ahmed CS** | Licensed under MIT | Version 2.0.0

![AFSA](https://img.shields.io/badge/AFSA-Advanced%20Forensic%20Security%20Analyzer-red?style=for-the-badge)

---

## 🌟 Features

AFSA is a comprehensive security analysis platform combining multiple reconnaissance techniques in one powerful CLI tool:

| Feature | Description | Status |
|---------|-------------|--------|
| **DNS Reconnaissance** | Complete DNS record enumeration (A, AAAA, MX, NS, CNAME, TXT) | ✅ |
| **IP Intelligence** | Advanced IP analysis with RFC classification & reverse DNS | ✅ |
| **Firewall Analysis** | Firewall status checking, rule enumeration, & port scanning | ✅ |
| **WAF Detection** | Web Application Firewall identification & analysis | ✅ |
| **WHOIS Lookup** | Domain & IP ownership information retrieval | ✅ |
| **Port Scanning** | Advanced TCP port scanning with service detection | ✅ |
| **Geolocation** | IP geographical analysis & ISP information | ✅ |

### 🎯 Supported WAFs
- ☁️ Cloudflare
- 🟠 AWS WAF
- 🔒 ModSecurity
- 🌐 Akamai
- 🛡️ Imperva/Incapsula
- 🟦 F5 BIG-IP
- 🟨 Barracuda
- 🟥 Sucuri
- 📝 Wordfence

---

## 📦 Installation

### Method 1: Direct Installation from GitHub (Recommended)

```bash
go install github.com/tanvircs/afsa@latest
```

Then use it globally:
```bash
afsa dns example.com
```

### Method 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/tanvircs/afsa.git
cd afsa

# Download dependencies
go mod download

# Build the application
go build -o afsa .

# Run it
./afsa dns example.com
```

### Method 3: Install Globally

```bash
# Build and install globally (macOS/Linux)
go install github.com/tanvircs/afsa@latest

# Verify installation
which afsa
afsa --help
```

---

## 🚀 Quick Start

### Basic Usage

```bash
# Show help
afsa --help

# DNS Reconnaissance
afsa dns example.com
afsa dns google.com -v

# IP Address Analysis
afsa ip 8.8.8.8
afsa ip 192.168.1.1 --verbose

# Firewall Analysis
afsa firewall status
afsa firewall test example.com
afsa firewall test example.com -p 22,80,443,3306

# WAF Detection
afsa waf github.com
afsa waf cloudflare.com --test-xss

# WHOIS Lookup
afsa whois example.com
afsa whois 8.8.8.8

# Port Scanning
afsa scan example.com
afsa scan example.com --deep
afsa scan example.com -r 1-1000

# Geolocation
afsa geo 8.8.8.8
afsa geo 1.1.1.1
```

---

## 📋 Command Reference

### DNS Reconnaissance
```bash
afsa dns [domain] [flags]

Flags:
  -v, --verbose     Show detailed information
  -t, --timeout     Query timeout in seconds (default: 10)

Examples:
  afsa dns example.com
  afsa dns google.com -v
  afsa dns example.com --timeout=15
```

### IP Intelligence
```bash
afsa ip [address] [flags]

Flags:
  -v, --verbose    Show detailed analysis

Examples:
  afsa ip 8.8.8.8
  afsa ip 192.168.1.1 -v
  afsa ip 2001:4860:4860::8888
```

### Firewall Analysis
```bash
afsa firewall [status|rules|test] [flags]

Subcommands:
  status     - Check firewall status
  rules      - List firewall rules (requires sudo)
  test       - Test port connectivity

Flags:
  -p, --ports      Comma-separated ports (default: 80,443,22)
  -d, --detailed   Show detailed information

Examples:
  afsa firewall status
  afsa firewall test example.com
  afsa firewall test example.com -p 22,80,443,3306 -d
```

### WAF Detection
```bash
afsa waf [domain] [flags]

Flags:
  --test-xss     Test XSS payload detection
  --test-sqli    Test SQL injection detection

Examples:
  afsa waf example.com
  afsa waf cloudflare.com --test-xss
  afsa waf example.com --test-sqli
```

### WHOIS Lookup
```bash
afsa whois [domain|ip]

Examples:
  afsa whois example.com
  afsa whois 8.8.8.8
```

### Port Scanning
```bash
afsa scan [hostname] [flags]

Flags:
  -r, --range        Port range (e.g., 1-1000)
  -d, --deep         Deep scan (all ports)
  -c, --common-only  Scan only common ports

Examples:
  afsa scan example.com
  afsa scan example.com -r 1-1000
  afsa scan example.com --deep
```

### Geolocation
```bash
afsa geo [ip]

Examples:
  afsa geo 8.8.8.8
  afsa geo 1.1.1.1
```

---

## 💡 Usage Examples

### Complete Domain Reconnaissance
```bash
#!/bin/bash
# Comprehensive domain security audit

DOMAIN="example.com"

echo "=== Starting Full Domain Reconnaissance ==="

# DNS Records
afsa dns $DOMAIN -v

# Firewall Check
afsa firewall test $DOMAIN -p 80,443,22

# WAF Detection
afsa waf $DOMAIN

# WHOIS Information
afsa whois $DOMAIN

# Port Scan
afsa scan $DOMAIN
```

### IP Address Analysis
```bash
# Analyze an IP address
IP="8.8.8.8"

afsa ip $IP -v
afsa whois $IP
afsa geo $IP
```

### Port Scanning with Service Detection
```bash
# Scan common ports
afsa scan example.com

# Scan specific range
afsa scan example.com -r 1-5000

# Deep scan (slow)
afsa scan example.com --deep
```

---

## 🏗️ Architecture

```
afsa/
├── main.go                 # Entry point
├── go.mod                  # Module definition
├── go.sum                  # Dependency checksums
├── LICENSE                 # MIT License
├── README.md               # This file
├── build.sh                # Build helper script
├── afsa                    # Compiled binary
└── cmd/
    ├── root.go             # CLI framework & banner
    ├── dns.go              # DNS reconnaissance
    ├── ip.go               # IP intelligence
    ├── firewall.go         # Firewall analysis
    ├── waf.go              # WAF detection
    ├── whois.go            # WHOIS lookup
    ├── scan.go             # Port scanning
    └── geo.go              # Geolocation analysis
```

---

## 🔧 System Requirements

| Requirement | Version | Notes |
|-------------|---------|-------|
| Go | 1.21+ | [Download](https://golang.org/dl) |
| macOS | 10.14+ | M1/M2 compatible |
| Linux | Ubuntu 18+, Debian 9+, CentOS 7+ | Any Linux distro |
| Network | Internet connection | For DNS/WHOIS queries |

### Elevated Privileges
Some features require `sudo`:
- `afsa firewall rules` - View firewall rules
- `afsa firewall status` - Detailed firewall info

---

## 🔐 Security & Disclaimer

### ⚠️ Important
- **AFSA is for authorized security testing only**
- Only test systems you own or have explicit permission to test
- Unauthorized access to computer networks is illegal
- Respect privacy laws and regulations (GDPR, CCPA, etc.)
- Use responsibly and ethically

### Best Practices
1. ✅ Always get written authorization before security testing
2. ✅ Respect rate limits and don't overload target systems
3. ✅ Follow responsible disclosure practices
4. ✅ Document and log all security tests
5. ✅ Keep AFSA updated for latest security features

---

## 🤝 Contributing

Contributions are welcome! To contribute:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Areas for Contribution
- Additional WAF detection signatures
- GeoIP database integration
- HTTP client for header analysis
- Additional port service definitions
- Performance improvements
- Documentation enhancements

---

## 📝 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

```
Copyright (c) 2024 Tanvir Ahmed CS

Licensed under the MIT License. See LICENSE file for full details.
```

---

## 🙋 Support & Contact

- **Author**: Tanvir Ahmed CS
- **Repository**: [github.com/tanvircs/afsa](https://github.com/tanvircs/afsa)
- **Issues**: [GitHub Issues](https://github.com/tanvircs/afsa/issues)

### Getting Help
```bash
# Show general help
afsa --help

# Show command-specific help
afsa dns --help
afsa ip --help
afsa firewall --help
afsa waf --help
afsa whois --help
afsa scan --help
afsa geo --help
```

---

## 🎓 Learning Resources

### Security Concepts
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [DNS Security](https://en.wikipedia.org/wiki/DNS_security_extension)
- [RFC 1918 - Private Networks](https://tools.ietf.org/html/rfc1918)
- [WAF Bypass Techniques](https://www.acunetix.com/blog/articles/waf-bypass-techniques/)

### CLI Development
- [Cobra Framework](https://github.com/spf13/cobra)
- [Color Package](https://github.com/fatih/color)
- [Go Official Docs](https://golang.org/doc/)

---

## 📊 Roadmap

### Version 2.1 (Planned)
- [ ] HTTPS header analysis
- [ ] DNSSEC validation
- [ ] Certificate transparency logs
- [ ] Subdomain enumeration
- [ ] HTTP request client

### Version 3.0 (Planned)
- [ ] Database integration (MaxMind GeoIP2)
- [ ] Web UI dashboard
- [ ] Automated vulnerability scanning
- [ ] Report generation (PDF/HTML)
- [ ] Integration with VirusTotal

---

## 🙏 Credits

**Created by**: Tanvir Ahmed CS

**Built with**:
- [Go Programming Language](https://golang.org)
- [Cobra CLI Framework](https://github.com/spf13/cobra)
- [Color Package](https://github.com/fatih/color)

**Inspired by**: Security researchers and penetration testers worldwide

---

## 📜 Version History

### v2.0.0 (Current)
- ✅ Complete rewrite with red branding
- ✅ Added WHOIS lookup
- ✅ Added port scanning
- ✅ Added geolocation module
- ✅ Enhanced WAF detection
- ✅ Professional documentation
- ✅ GitHub installation support

### v1.0.0
- Basic DNS, IP, Firewall, WAF modules

---

<div align="center">

### Made with ❤️ by Tanvir Ahmed CS

**[⬆ back to top](#-afsa---advanced-forensic-security-analyzer)**

**[GitHub](https://github.com/tanvircs/afsa) | [License](LICENSE) | [Version 2.0.0](https://github.com/tanvircs/afsa/releases)**

</div>
