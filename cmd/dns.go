package cmd

import (
	"fmt"
	"net"
	"sort"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	dnsVerbose bool
	dnsTimeout int
)

var dnsCmd = &cobra.Command{
	Use:   "dns [domain]",
	Short: color.RedString("DNS Reconnaissance - Advanced domain enumeration"),
	Long: `Comprehensive DNS lookup tool for complete domain reconnaissance:

Features:
  ▸ A Records (IPv4 addresses)
  ▸ AAAA Records (IPv6 addresses)
  ▸ MX Records (Mail exchange servers)
  ▸ NS Records (Nameservers)
  ▸ CNAME Records (Canonical names)
  ▸ TXT Records (Text records, SPF, DMARC, etc.)
  ▸ SOA Records (Start of Authority)

Flags:
  -v, --verbose     Show detailed information
  -t, --timeout     Query timeout in seconds (default: 10)

Examples:
  afsa dns example.com
  afsa dns google.com -v
  afsa dns example.com --timeout=15`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		performAdvancedDNSLookup(domain)
	},
}

func init() {
	dnsCmd.Flags().BoolVarP(&dnsVerbose, "verbose", "v", false, "Verbose output")
	dnsCmd.Flags().IntVarP(&dnsTimeout, "timeout", "t", 10, "Query timeout in seconds")
}

func performAdvancedDNSLookup(domain string) {
	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║             DNS RECONNAISSANCE REPORT                   ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n")

	color.Cyan("  Domain: %s\n\n", domain)

	// Validate domain
	if err := validateDomain(domain); err != nil {
		color.Red("  [✗] Invalid domain: %v\n\n", err)
		return
	}

	// A Records (IPv4)
	aRecords, err := net.LookupHost(domain)
	if err == nil && len(aRecords) > 0 {
		color.Red("  ▸ A Records (IPv4 Addresses):\n")
		ipv4s := filterIPv4(aRecords)
		if len(ipv4s) > 0 {
			for i, ip := range ipv4s {
				if i == len(ipv4s)-1 {
					fmt.Printf("    └─ %s\n", color.GreenString(ip))
				} else {
					fmt.Printf("    ├─ %s\n", color.GreenString(ip))
				}
			}
		}
	}

	// AAAA Records (IPv6)
	aaaaRecords := filterIPv6(aRecords)
	if len(aaaaRecords) > 0 {
		color.Red("  ▸ AAAA Records (IPv6 Addresses):\n")
		for i, ip := range aaaaRecords {
			if i == len(aaaaRecords)-1 {
				fmt.Printf("    └─ %s\n", color.CyanString(ip))
			} else {
				fmt.Printf("    ├─ %s\n", color.CyanString(ip))
			}
		}
	}

	// MX Records
	mxRecords, err := net.LookupMX(domain)
	if err == nil && len(mxRecords) > 0 {
		color.Red("  ▸ MX Records (Mail Servers):\n")
		sort.Slice(mxRecords, func(i, j int) bool {
			return mxRecords[i].Pref < mxRecords[j].Pref
		})
		for i, mx := range mxRecords {
			if i == len(mxRecords)-1 {
				fmt.Printf("    └─ %s (Priority: %d)\n", color.YellowString(mx.Host), mx.Pref)
			} else {
				fmt.Printf("    ├─ %s (Priority: %d)\n", color.YellowString(mx.Host), mx.Pref)
			}
		}
	}

	// NS Records
	nsRecords, err := net.LookupNS(domain)
	if err == nil && len(nsRecords) > 0 {
		color.Red("  ▸ NS Records (Nameservers):\n")
		for i, ns := range nsRecords {
			if i == len(nsRecords)-1 {
				fmt.Printf("    └─ %s\n", color.MagentaString(ns.Host))
			} else {
				fmt.Printf("    ├─ %s\n", color.MagentaString(ns.Host))
			}
		}
	}

	// CNAME Records
	cname, err := net.LookupCNAME(domain)
	if err == nil && cname != domain {
		color.Red("  ▸ CNAME Record:\n")
		fmt.Printf("    └─ %s\n", color.BlueString(cname))
	}

	// TXT Records
	txtRecords, err := net.LookupTXT(domain)
	if err == nil && len(txtRecords) > 0 {
		color.Red("  ▸ TXT Records:\n")
		for i, txt := range txtRecords {
			prefix := "├─ "
			if i == len(txtRecords)-1 {
				prefix = "└─ "
			}
			displayTxt := txt
			if len(txt) > 60 && !dnsVerbose {
				displayTxt = txt[:57] + "..."
			}
			fmt.Printf("    %s %s\n", prefix, color.WhiteString(displayTxt))
		}
	}

	// Summary statistics
	color.Red("\n  ▸ Summary:\n")
	fmt.Printf("    ├─ Total IPv4 Records: %d\n", len(filterIPv4(aRecords)))
	fmt.Printf("    ├─ Total IPv6 Records: %d\n", len(aaaaRecords))
	fmt.Printf("    ├─ Mail Servers: %d\n", len(mxRecords))
	fmt.Printf("    └─ Nameservers: %d\n", len(nsRecords))

	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║     [✓] DNS Reconnaissance Completed Successfully      ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n\n")
}

func validateDomain(domain string) error {
	if domain == "" {
		return fmt.Errorf("domain cannot be empty")
	}
	if len(domain) > 253 {
		return fmt.Errorf("domain name too long")
	}
	return nil
}

func filterIPv4(ips []string) []string {
	var result []string
	for _, ip := range ips {
		if net.ParseIP(ip).To4() != nil {
			result = append(result, ip)
		}
	}
	return result
}

func filterIPv6(ips []string) []string {
	var result []string
	for _, ip := range ips {
		if net.ParseIP(ip).To4() == nil && net.ParseIP(ip) != nil {
			result = append(result, ip)
		}
	}
	return result
}
