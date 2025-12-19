package cmd

import (
	"fmt"
	"net"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ipVerbose bool

var ipCmd = &cobra.Command{
	Use:   "ip [address]",
	Short: color.RedString("IP Intelligence - Comprehensive IP analysis"),
	Long: `Advanced IP address lookup and analysis tool:

Features:
  ▸ IP Version Detection (IPv4/IPv6)
  ▸ IP Type Classification (Public, Private, Loopback, etc.)
  ▸ RFC Classifications (RFC1918, RFC5771, etc.)
  ▸ Reverse DNS Lookup
  ▸ Special Address Detection
  ▸ CIDR Range Information

Flags:
  -v, --verbose    Show detailed analysis

Examples:
  afsa ip 8.8.8.8
  afsa ip 192.168.1.1 -v
  afsa ip 2001:4860:4860::8888`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ipAddr := args[0]
		performAdvancedIPLookup(ipAddr)
	},
}

func init() {
	ipCmd.Flags().BoolVarP(&ipVerbose, "verbose", "v", false, "Verbose output")
}

func performAdvancedIPLookup(ipAddr string) {
	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║            IP INTELLIGENCE REPORT                      ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n")

	ip := net.ParseIP(ipAddr)
	if ip == nil {
		color.Red("  [✗] Invalid IP address format: %s\n\n", ipAddr)
		return
	}

	// Basic Info
	color.Cyan("  Target IP: %s\n\n", ipAddr)

	// IP Version
	color.Red("  ▸ IP Version:\n")
	if ip.To4() != nil {
		fmt.Printf("    └─ %s\n", color.GreenString("IPv4"))
	} else if ip.To16() != nil {
		fmt.Printf("    └─ %s\n", color.CyanString("IPv6"))
	}

	// IP Classification
	color.Red("  ▸ IP Classification:\n")
	classifications := getIPClassifications(ip)
	for i, class := range classifications {
		if i == len(classifications)-1 {
			fmt.Printf("    └─ %s\n", color.YellowString(class))
		} else {
			fmt.Printf("    ├─ %s\n", color.YellowString(class))
		}
	}

	// Reverse DNS
	color.Red("  ▸ Reverse DNS Lookup:\n")
	hostnames, err := net.LookupAddr(ipAddr)
	if err == nil && len(hostnames) > 0 {
		for i, hostname := range hostnames {
			if i == len(hostnames)-1 {
				fmt.Printf("    └─ %s\n", color.BlueString(hostname))
			} else {
				fmt.Printf("    ├─ %s\n", color.BlueString(hostname))
			}
		}
	} else {
		fmt.Printf("    └─ %s\n", color.WhiteString("No reverse DNS records"))
	}

	// Special Characteristics
	color.Red("  ▸ Special Characteristics:\n")
	characteristics := getSpecialCharacteristics(ip)
	for i, char := range characteristics {
		if i == len(characteristics)-1 {
			fmt.Printf("    └─ %s\n", color.MagentaString(char))
		} else {
			fmt.Printf("    ├─ %s\n", color.MagentaString(char))
		}
	}

	// Security Analysis
	if ipVerbose {
		color.Red("  ▸ Security Analysis:\n")
		if isPrivateIP(ip) {
			fmt.Printf("    ├─ %s\n", color.GreenString("✓ Private - Safe for internal use"))
		} else {
			fmt.Printf("    ├─ %s\n", color.YellowString("⚠ Public - Exposed to internet"))
		}

		if ip.IsLoopback() {
			fmt.Printf("    ├─ %s\n", color.GreenString("✓ Loopback - Local machine only"))
		}

		if ip.IsMulticast() {
			fmt.Printf("    └─ %s\n", color.YellowString("⚠ Multicast - Group communication"))
		}
	}

	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║      [✓] IP Intelligence Analysis Completed           ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n\n")
}

func getIPClassifications(ip net.IP) []string {
	var classifications []string

	if ip.IsLoopback() {
		classifications = append(classifications, "Loopback Address (RFC5735)")
	} else if ip.IsPrivate() {
		classifications = append(classifications, "Private Address (RFC1918)")
	} else if ip.IsMulticast() {
		classifications = append(classifications, "Multicast Address (RFC5771)")
	} else if ip.IsLinkLocalUnicast() {
		classifications = append(classifications, "Link-Local Unicast (RFC3927)")
	} else if ip.IsLinkLocalMulticast() {
		classifications = append(classifications, "Link-Local Multicast (RFC5771)")
	} else {
		classifications = append(classifications, "Public Address (Routable)")
	}

	return classifications
}

func getSpecialCharacteristics(ip net.IP) []string {
	var chars []string

	if ip.IsUnspecified() {
		chars = append(chars, "Unspecified Address (0.0.0.0 or ::)")
	}
	if ip.IsLoopback() {
		chars = append(chars, "Loopback (127.0.0.1 or ::1)")
	}
	if ip.IsPrivate() {
		chars = append(chars, "RFC1918 Private Range")
	}
	if ip.IsMulticast() {
		chars = append(chars, "Multicast Group")
	}
	if ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		chars = append(chars, "Link-Local Address")
	}
	if ip.IsGlobalUnicast() {
		chars = append(chars, "Global Unicast (Routable)")
	}
	if ip.IsInterfaceLocalMulticast() {
		chars = append(chars, "Interface-Local Multicast")
	}

	if len(chars) == 0 {
		chars = append(chars, "Standard Routable Public Address")
	}

	return chars
}

func isPrivateIP(ip net.IP) bool {
	return ip.IsPrivate() || ip.IsLoopback()
}
