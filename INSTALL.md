## AFSA - Installation & Setup Guide

**Advanced Forensic Security Analyzer v2.0.0**  
Created by: Tanvir Ahmed CS  
Licensed under: MIT License

---

### Prerequisites

- **Go 1.21+** ([Download](https://golang.org/dl))
- **macOS 10.14+** or **Linux** (Ubuntu 18+, Debian 9+, CentOS 7+)
- **Internet connection** for DNS/WHOIS queries

---

### Option 1: Install from GitHub (Recommended)

```bash
# Install directly from GitHub
go install github.com/tanvircs/afsa@latest

# Verify installation
afsa --help

# Start using
afsa dns example.com
```

**Note**: This will install to `$HOME/go/bin/`. Ensure this is in your PATH.

---

### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/tanvircs/afsa.git
cd afsa

# Download dependencies
go mod download

# Build
go build -o afsa .

# Run locally
./afsa dns example.com

# Or install globally
sudo cp afsa /usr/local/bin/
afsa dns example.com
```

---

### Option 3: Build Script

```bash
# Use the included build script
chmod +x build.sh
./build.sh

# The script will:
# - Check Go installation
# - Download dependencies
# - Build the binary
# - Run a quick test
```

---

### Add to PATH (Optional)

```bash
# Copy to system bin
sudo cp afsa /usr/local/bin/

# Or add current directory to PATH
export PATH="$PATH:$(pwd)"

# Make permanent (add to ~/.bashrc or ~/.zshrc)
echo 'export PATH="$PATH:/path/to/afsa"' >> ~/.zshrc
```

---

### Verify Installation

```bash
# Check version and help
afsa --help

# Test DNS command
afsa dns google.com

# Test IP command
afsa ip 8.8.8.8

# Test firewall command
afsa firewall status
```

---

### Troubleshooting

**"afsa: command not found"**
- Add `/usr/local/bin` to your PATH
- Or use full path: `/usr/local/bin/afsa dns example.com`

**"permission denied"**
```bash
chmod +x afsa
sudo cp afsa /usr/local/bin/
```

**Go installation issues**
- Check: `go version`
- Download from: https://golang.org/dl
- Follow: https://golang.org/doc/install

**Need sudo for some commands**
```bash
sudo afsa firewall status
sudo afsa firewall rules
```

---

### Project Structure

```
afsa/
├── main.go                  # Entry point
├── go.mod                   # Module definition
├── go.sum                   # Dependencies
├── LICENSE                  # MIT License
├── README.md               # Full documentation
├── INSTALL.md              # This file
├── build.sh                # Build helper
├── afsa                    # Binary (after build)
└── cmd/
    ├── root.go             # CLI framework
    ├── dns.go              # DNS module
    ├── ip.go               # IP module
    ├── firewall.go         # Firewall module
    ├── waf.go              # WAF module
    ├── whois.go            # WHOIS module
    ├── scan.go             # Port scanning
    └── geo.go              # Geolocation
```

---

### Getting Started

```bash
# View all commands
afsa --help

# Get command-specific help
afsa dns --help
afsa ip --help
afsa firewall --help
afsa waf --help
afsa whois --help
afsa scan --help
afsa geo --help

# Basic examples
afsa dns example.com
afsa ip 8.8.8.8
afsa firewall test example.com
afsa waf github.com
afsa whois example.com
afsa scan example.com
afsa geo 8.8.8.8
```

---

### Support

- **GitHub**: https://github.com/tanvircs/afsa
- **Author**: Tanvir Ahmed CS
- **License**: MIT

For issues or questions, open an issue on GitHub.

---

## Important: Security & Legal

⚠️ **AFSA is for authorized testing only!**

- Only test systems you own or have permission to test
- Unauthorized network access is illegal
- Respect privacy laws (GDPR, CCPA, etc.)
- Use responsibly and ethically

---

**Thank you for using AFSA!**
