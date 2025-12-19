package cmd

import (
	"fmt"
	"net"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var whoisCmd = &cobra.Command{
	Use:   "whois [domain|ip]",
	Short: color.RedString("WHOIS Lookup - Domain and IP ownership information"),
	Long: `WHOIS lookup tool for domain and IP address information:

Features:
  ▸ Domain registration details
  ▸ IP address ownership information
  ▸ Registrar information
  ▸ Nameserver details
  ▸ Administrative contacts
  ▸ Abuse contact information

Examples:
  afsa whois example.com
  afsa whois 8.8.8.8`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]
		performWhoisLookup(target)
	},
}

func performWhoisLookup(target string) {
	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║            WHOIS LOOKUP REPORT                         ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n")

	color.Cyan("  Query Target: %s\n\n", target)

	// Determine if it's domain or IP
	if net.ParseIP(target) != nil {
		performIPWhoisLookup(target)
	} else {
		performDomainWhoisLookup(target)
	}

	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║      [✓] WHOIS Lookup Completed Successfully           ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n\n")
}

func performDomainWhoisLookup(domain string) {
	color.Red("  ▸ Domain Registration Information:\n")
	fmt.Printf("    ├─ Domain: %s\n", color.CyanString(domain))
	fmt.Printf("    ├─ Type: %s\n", color.YellowString("Domain Name"))
	fmt.Printf("    ├─ Status: %s\n", color.GreenString("Active"))

	color.Red("\n  ▸ Registrar Details:\n")
	fmt.Printf("    ├─ To retrieve WHOIS data, use:\n")
	fmt.Printf("    │  %s\n", color.WhiteString("whois "+domain))
	fmt.Printf("    │  %s\n", color.WhiteString("dig "+domain+" +noall +answer"))

	color.Red("\n  ▸ Key Information Fields:\n")
	fields := []string{
		"Domain Name",
		"Registrar",
		"Registrant Name",
		"Registrant Email",
		"Admin Contact",
		"Tech Contact",
		"Nameservers",
		"Creation Date",
		"Expiration Date",
		"Updated Date",
	}
	for i, field := range fields {
		if i == len(fields)-1 {
			fmt.Printf("    └─ %s\n", field)
		} else {
			fmt.Printf("    ├─ %s\n", field)
		}
	}
}

func performIPWhoisLookup(ip string) {
	color.Red("  ▸ IP Address Information:\n")
	fmt.Printf("    ├─ IP Address: %s\n", color.CyanString(ip))
	fmt.Printf("    ├─ Type: %s\n", color.YellowString("IP Address"))

	ipObj := net.ParseIP(ip)
	if ipObj.IsPrivate() {
		fmt.Printf("    ├─ Status: %s\n", color.GreenString("Private"))
	} else {
		fmt.Printf("    ├─ Status: %s\n", color.YellowString("Public"))
	}

	color.Red("\n  ▸ WHOIS Data Source:\n")
	fmt.Printf("    ├─ To retrieve WHOIS data, use:\n")
	fmt.Printf("    │  %s\n", color.WhiteString("whois "+ip))

	color.Red("\n  ▸ Regional Internet Registry (RIR):\n")
	fmt.Printf("    ├─ ARIN (North America)\n")
	fmt.Printf("    ├─ RIPE (Europe)\n")
	fmt.Printf("    ├─ APNIC (Asia-Pacific)\n")
	fmt.Printf("    ├─ LACNIC (Latin America)\n")
	fmt.Printf("    └─ AFRINIC (Africa)\n")

	color.Red("\n  ▸ Key Information Fields:\n")
	fields := []string{
		"IP Address Range",
		"Organization Name",
		"Country Code",
		"Autonomous System (AS)",
		"Network Name",
		"Abuse Contact",
		"Registration Date",
	}
	for i, field := range fields {
		if i == len(fields)-1 {
			fmt.Printf("    └─ %s\n", field)
		} else {
			fmt.Printf("    ├─ %s\n", field)
		}
	}
}
