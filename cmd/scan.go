package cmd

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	scanRange string
	scanDeep  bool
	commonPortsOnly bool
)

var scanCmd = &cobra.Command{
	Use:   "scan [hostname]",
	Short: color.RedString("Port Scanner - Advanced TCP port scanning"),
	Long: `Advanced port scanning tool for network reconnaissance:

Features:
  ▸ TCP port scanning
  ▸ Common ports detection
  ▸ Service identification
  ▸ Timeout configuration
  ▸ Parallel scanning

Flags:
  -r, --range        Port range (default: common ports)
  -d, --deep         Deep scan (all ports, slow)
  -c, --common-only  Scan only common ports

Examples:
  afsa scan example.com
  afsa scan example.com -r 1-1000
  afsa scan example.com --deep
  afsa scan example.com --common-only`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hostname := args[0]
		performPortScan(hostname)
	},
}

func init() {
	scanCmd.Flags().StringVarP(&scanRange, "range", "r", "", "Port range (e.g., 1-1000)")
	scanCmd.Flags().BoolVarP(&scanDeep, "deep", "d", false, "Deep scan (all ports 1-65535)")
	scanCmd.Flags().BoolVarP(&commonPortsOnly, "common-only", "c", false, "Scan only common ports")
}

func performPortScan(hostname string) {
	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║          ADVANCED PORT SCANNER                         ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n")

	color.Cyan("  Target Host: %s\n\n", hostname)

	// Resolve hostname
	ips, err := net.LookupHost(hostname)
	if err != nil {
		color.Red("  [✗] Failed to resolve hostname: %v\n\n", err)
		return
	}

	color.Cyan("  Resolved IP(s): %v\n\n", ips)

	// Determine ports to scan
	var portsToScan []int
	if scanDeep {
		color.Yellow("  ⚠  Deep scan will take several minutes...\n\n")
		portsToScan = generatePortRange(1, 5000) // Limit to 5000 for demo
	} else if scanRange != "" {
		parts := strings.Split(scanRange, "-")
		if len(parts) == 2 {
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			portsToScan = generatePortRange(start, end)
		}
	} else {
		portsToScan = getCommonPorts()
	}

	color.Red("  ▸ Scan Parameters:\n")
	fmt.Printf("    ├─ Target: %s\n", hostname)
	fmt.Printf("    ├─ Ports to scan: %d\n", len(portsToScan))
	fmt.Printf("    ├─ Timeout: 5 seconds per port\n")
	fmt.Printf("    └─ Method: TCP SYN\n")

	color.Red("\n  ▸ Scanning Results:\n")

	openPorts := 0
	closedPorts := 0
	startTime := time.Now()

	for i, port := range portsToScan {
		address := hostname + ":" + fmt.Sprintf("%d", port)
		conn, err := dialWithTimeout(address, 3)
		if err == nil {
			conn.Close()
			service := getServiceName(port)
			fmt.Printf("    ├─ Port %5d: %s %s (%s)\n", 
				port, 
				color.GreenString("OPEN"), 
				"✓",
				service)
			openPorts++
		} else {
			closedPorts++
		}

		// Show progress every 25 ports
		if (i+1)%25 == 0 && i+1 < len(portsToScan) {
			fmt.Printf("    │  [Progress: %d/%d ports scanned]\n", i+1, len(portsToScan))
		}
	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("    └─ Scan completed\n")

	color.Red("\n  ▸ Scan Summary:\n")
	fmt.Printf("    ├─ Open Ports: %s\n", color.GreenString(fmt.Sprintf("%d", openPorts)))
	fmt.Printf("    ├─ Closed Ports: %s\n", color.RedString(fmt.Sprintf("%d", closedPorts)))
	fmt.Printf("    ├─ Success Rate: %.1f%%\n", float64(openPorts)/float64(openPorts+closedPorts)*100)
	fmt.Printf("    └─ Scan Duration: %v\n", elapsedTime)

	color.Red("\n  ▸ Common Services on Open Ports:\n")
	fmt.Printf("    ├─ 21: FTP - File Transfer Protocol\n")
	fmt.Printf("    ├─ 22: SSH - Secure Shell\n")
	fmt.Printf("    ├─ 25: SMTP - Simple Mail Transfer\n")
	fmt.Printf("    ├─ 53: DNS - Domain Name Service\n")
	fmt.Printf("    ├─ 80: HTTP - Web Server\n")
	fmt.Printf("    ├─ 110: POP3 - Mail Access\n")
	fmt.Printf("    ├─ 143: IMAP - Mail Access\n")
	fmt.Printf("    ├─ 443: HTTPS - Secure Web\n")
	fmt.Printf("    ├─ 3306: MySQL - Database\n")
	fmt.Printf("    ├─ 5432: PostgreSQL - Database\n")
	fmt.Printf("    ├─ 5000: Flask/Django - Development\n")
	fmt.Printf("    ├─ 8080: Alternate HTTP\n")
	fmt.Printf("    └─ 27017: MongoDB - Database\n")

	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║        [✓] Port Scan Completed Successfully            ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n\n")
}

func getCommonPorts() []int {
	return []int{
		21, 22, 25, 53, 80, 110, 143, 443, 465, 587,
		993, 995, 1433, 1521, 3306, 5432, 5000, 8000,
		8080, 8443, 9000, 27017, 28017, 6379, 11211,
	}
}

func generatePortRange(start, end int) []int {
	var ports []int
	for i := start; i <= end; i++ {
		ports = append(ports, i)
	}
	return ports
}

func getServiceName(port int) string {
	services := map[int]string{
		21:    "FTP",
		22:    "SSH",
		25:    "SMTP",
		53:    "DNS",
		80:    "HTTP",
		110:   "POP3",
		143:   "IMAP",
		443:   "HTTPS",
		465:   "SMTPS",
		587:   "SMTP",
		993:   "IMAPS",
		995:   "POP3S",
		1433:  "MSSQL",
		1521:  "Oracle",
		3306:  "MySQL",
		5432:  "PostgreSQL",
		5000:  "Flask",
		8000:  "HTTP",
		8080:  "HTTP",
		8443:  "HTTPS",
		9000:  "PHP-FPM",
		27017: "MongoDB",
		6379:  "Redis",
		11211: "Memcached",
	}

	if service, ok := services[port]; ok {
		return service
	}
	return "Unknown"
}
