package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	ForceColor = true
)

func SetupColor() error {
	color.NoColor = false
	return nil
}

func displayBanner() {
	red := color.New(color.FgRed, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	white := color.New(color.FgWhite)
	cyan := color.New(color.FgCyan)

	banner := `
	    .++++.              
	   +WQQQQQ0         
	   +o&       ╔════════════════════════════════════════════════════╗
	   80        ║                                                    ║
	   wW        ║   ╔═══╗ ╔═══╗ ╔═══╗ ╔═══╗ ╔═══╗                  ║
	   #O        ║   ║   ║ ║   ║ ║   ║ ║   ║ ║                      ║
	   o@+       ║   ║ ╔═╝ ║   ║ ║ ╔═╝ ║   ║ ║                      ║
	   wW        ║   ║ ║   ║   ║ ║ ║   ║   ║ ║                      ║
	   :OW:      ║   ║ ╚═╗ ╚═══╝ ║ ╚═╗ ╚═══╝ ║                      ║
	   +OW08.    ║                                                    ║
	             ║   Advanced Forensic Security Analyzer v2.0.0     ║
	             ║   In-depth Attack Surface Analysis & Detection    ║
	             ║                                                    ║
	             ║   Created by Tanvir Ahmed CS                      ║
	             ║   MIT License | github.com/tanvircs/afsa          ║
	             ║                                                    ║
	             ╚════════════════════════════════════════════════════╝
	`

	red.Println(banner)
	yellow.Println("\n   🔴 AFSA - Advanced Forensic Security Analyzer v2.0.0")
	cyan.Println("   🎯 Professional Security Reconnaissance Tool\n")
	white.Println()
}

var rootCmd = &cobra.Command{
	Use:   "afsa",
	Short: color.RedString("AFSA - Advanced Forensic Security Analyzer"),
	Long: color.RedString(`
╔════════════════════════════════════════════════════════════════╗
║  AFSA - Advanced Forensic Security Analyzer v2.0.0            ║
║  Professional Network & Security Reconnaissance Tool          ║
║  Created by Tanvir Ahmed CS | MIT License                     ║
╚════════════════════════════════════════════════════════════════╝

DESCRIPTION:
  AFSA is a comprehensive all-in-one command-line tool for advanced
  network and security reconnaissance. It combines multiple security
  analysis techniques in a single unified platform with professional
  red-themed styling and detailed reporting capabilities.

FEATURES:
  🔴 DNS Reconnaissance     - Complete DNS record enumeration (A, AAAA, MX, NS, CNAME, TXT)
  🔴 IP Intelligence       - Advanced IP analysis with RFC 1918/5735/5771 classification
  🔴 Firewall Analysis     - Firewall status checking and TCP port connectivity testing
  🔴 WAF Detection         - Web Application Firewall signature detection (9+ WAFs)
  🔴 WHOIS Lookup          - Domain and IP ownership information lookup
  🔴 Port Scanning         - Advanced TCP port scanning with service detection
  🔴 Geolocation Analysis  - IP geographical and ISP information analysis

USAGE:
  afsa [command] [flags] [arguments]

COMMANDS:
  dns       DNS Reconnaissance - Enumerate all DNS records for a domain
  ip        IP Intelligence - Analyze and classify an IP address
  firewall  Firewall Analysis - Check firewall status and test ports
  waf       WAF Detection - Identify Web Application Firewalls
  whois     WHOIS Lookup - Get domain and IP ownership information
  scan      Port Scanning - Scan and identify open ports
  geo       Geolocation - Get IP geographical and ISP information
  help      Show help information for any command

EXAMPLES:
  afsa dns example.com                    # Lookup all DNS records
  afsa dns google.com -v                  # Verbose DNS lookup
  afsa ip 8.8.8.8                         # Analyze IP address
  afsa ip 1.1.1.1 --verbose               # Detailed IP analysis
  afsa firewall status                    # Check firewall status
  afsa firewall test example.com          # Test port connectivity
  afsa firewall test example.com -p 22,80,443  # Test specific ports
  afsa waf github.com                     # Detect WAF signatures
  afsa waf cloudflare.com --test-xss      # WAF detection with XSS test
  afsa whois example.com                  # Domain WHOIS lookup
  afsa whois 8.8.8.8                      # IP WHOIS lookup
  afsa scan example.com                   # Scan common ports
  afsa scan example.com -r 1-1000         # Scan port range
  afsa scan example.com --deep            # Deep scan with all ports
  afsa geo 8.8.8.8                        # Get geolocation info

FLAGS:
  -h, --help              Show this help message
  -v, --verbose           Enable verbose output with additional details
  -d, --detailed          Show detailed output with extended information
  -p, --ports string      Specify ports to test (comma-separated)
  -r, --range string      Specify port range (e.g., 1-1000)
  -t, --timeout int       Set timeout in seconds (default: 5)

GLOBAL FLAGS:
  --help                  Show help for any command
  --version               Show version information

SECURITY NOTES:
  • Always get proper authorization before scanning networks/IPs
  • Respect responsible disclosure guidelines
  • Follow applicable laws and regulations
  • Use for legitimate security testing only
  • See LICENSE file for full terms

AUTHOR:
  Created by Tanvir Ahmed CS
  GitHub: github.com/tanvircs/afsa
  License: MIT

For detailed help on a specific command, use:
  afsa [command] --help

Example:
  afsa dns --help         # Get help for DNS command
  afsa waf --help         # Get help for WAF command
  afsa scan --help        # Get help for Port Scan command`),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Name() != "help" && len(args) > 0 {
			// Only show banner for specific commands
		}
	},
}

func Execute() {
	displayBanner()
	if err := rootCmd.Execute(); err != nil {
		color.Red("[✗] Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(dnsCmd)
	rootCmd.AddCommand(ipCmd)
	rootCmd.AddCommand(firewallCmd)
	rootCmd.AddCommand(wafCmd)
	rootCmd.AddCommand(whoisCmd)
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(geoCmd)
}
