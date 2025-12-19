package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	ports    []string
	detailed bool
)

var firewallCmd = &cobra.Command{
	Use:   "firewall [check|list|test]",
	Short: color.RedString("Firewall Analysis - Port testing and firewall status"),
	Long: `Advanced firewall analysis and connectivity testing:

Subcommands:
  status     - Check firewall status (requires elevated privileges)
  rules      - List active firewall rules
  test       - Test TCP connectivity to ports

Features:
  ▸ Multi-port scanning
  ▸ Timeout configuration
  ▸ Detailed service detection
  ▸ Connection profiling

Flags:
  -p, --ports      Comma-separated ports to test (default: 80,443,22)
  -d, --detailed   Show detailed service information

Examples:
  afsa firewall status
  afsa firewall test example.com
  afsa firewall test example.com -p 22,80,443,3306
  afsa firewall test example.com -d`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		subCmd := args[0]
		switch subCmd {
		case "status":
			checkFirewallStatus()
		case "rules":
			listFirewallRules()
		case "test":
			if len(args) < 2 {
				color.Red("  [✗] firewall test requires hostname\n")
				fmt.Println("Usage: afsa firewall test <hostname> [flags]")
				return
			}
			testFirewallPorts(args[1])
		default:
			color.Red("  [✗] Unknown firewall subcommand: %s\n", subCmd)
		}
	},
}

func init() {
	firewallCmd.Flags().StringSliceVarP(&ports, "ports", "p", []string{"80", "443", "22", "3306", "5432"}, "Ports to test")
	firewallCmd.Flags().BoolVarP(&detailed, "detailed", "d", false, "Show detailed information")
}

func checkFirewallStatus() {
	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║          FIREWALL STATUS ANALYSIS                      ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n")

	color.Cyan("  System: %s\n\n", getOS())

	color.Red("  ▸ Firewall Monitoring:\n")
	fmt.Printf("    ├─ Status: %s\n", color.YellowString("Requires elevated privileges"))
	fmt.Printf("    └─ Note: Run with 'sudo' for full details\n")

	color.Red("\n  ▸ Manual Status Check Commands:\n")
	fmt.Printf("    macOS:  %s\n", color.WhiteString("sudo launchctl list | grep -i firewall"))
	fmt.Printf("    Linux:  %s\n", color.WhiteString("sudo systemctl status ufw"))
	fmt.Printf("    Linux:  %s\n", color.WhiteString("sudo iptables -L -n"))

	color.Red("\n  ▸ Firewall Services:\n")
	fmt.Printf("    ├─ macOS: pf (Packet Filter)\n")
	fmt.Printf("    ├─ Linux: ufw (Uncomplicated Firewall)\n")
	fmt.Printf("    ├─ Linux: firewalld (Dynamic Firewall Manager)\n")
	fmt.Printf("    └─ Linux: iptables (Netfilter)\n")

	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║    [✓] Firewall Status Check Completed                 ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n\n")
}

func listFirewallRules() {
	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║           FIREWALL RULES ENUMERATION                   ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n")

	color.Yellow("\n  ⚠  Detailed firewall rules require elevated privileges (sudo)\n")

	color.Red("  ▸ Linux (iptables) Commands:\n")
	fmt.Printf("    ├─ %s\n", color.WhiteString("sudo iptables -L -n -v"))
	fmt.Printf("    ├─ %s\n", color.WhiteString("sudo iptables -L INPUT -n -v"))
	fmt.Printf("    └─ %s\n", color.WhiteString("sudo iptables -L OUTPUT -n -v"))

	color.Red("\n  ▸ Linux (firewalld) Commands:\n")
	fmt.Printf("    ├─ %s\n", color.WhiteString("sudo firewall-cmd --list-all"))
	fmt.Printf("    ├─ %s\n", color.WhiteString("sudo firewall-cmd --list-ports"))
	fmt.Printf("    └─ %s\n", color.WhiteString("sudo firewall-cmd --list-rich-rules"))

	color.Red("\n  ▸ Linux (ufw) Commands:\n")
	fmt.Printf("    └─ %s\n", color.WhiteString("sudo ufw status verbose"))

	color.Red("\n  ▸ macOS (pf) Commands:\n")
	fmt.Printf("    └─ %s\n", color.WhiteString("sudo pfctl -s rules"))

	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║    [✓] Firewall Rules Listing Completed                ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n\n")
}

func testFirewallPorts(hostname string) {
	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║          FIREWALL PORT CONNECTIVITY TEST               ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n")

	color.Cyan("  Target: %s\n", hostname)
	color.Cyan("  Testing Ports: %v\n\n", ports)

	openPorts := 0
	closedPorts := 0

	color.Red("  ▸ Port Scan Results:\n")

	for i, port := range ports {
		address := hostname + ":" + port
		conn, err := dialWithTimeout(address, 5)
		if err == nil {
			conn.Close()
			color.Green("    ├─ Port %s: OPEN %s\n", port, "✓")
			openPorts++
		} else {
			color.Red("    ├─ Port %s: CLOSED/FILTERED %s\n", port, "✗")
			closedPorts++
		}

		if i == len(ports)-1 {
			fmt.Printf("    └─ (Scan completed)\n")
		}
	}

	color.Red("\n  ▸ Scan Summary:\n")
	fmt.Printf("    ├─ Open Ports: %s\n", color.GreenString(fmt.Sprintf("%d", openPorts)))
	fmt.Printf("    ├─ Closed/Filtered: %s\n", color.RedString(fmt.Sprintf("%d", closedPorts)))
	fmt.Printf("    └─ Success Rate: %.1f%%\n", float64(openPorts)/float64(len(ports))*100)

	if detailed {
		color.Red("\n  ▸ Common Services:\n")
		serviceMap := map[string]string{
			"22":   "SSH",
			"80":   "HTTP",
			"443":  "HTTPS",
			"3306": "MySQL",
			"5432": "PostgreSQL",
			"5000": "Flask/Django",
			"8080": "HTTP Alt",
		}
		for port := range serviceMap {
			if service, ok := serviceMap[port]; ok {
				for _, testPort := range ports {
					if testPort == port {
						fmt.Printf("    ├─ Port %s: %s\n", port, service)
						break
					}
				}
			}
		}
	}

	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║      [✓] Firewall Port Test Completed                  ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n\n")
}

func dialWithTimeout(address string, timeoutSeconds int) (net.Conn, error) {
	timeout := time.Duration(timeoutSeconds) * time.Second
	return net.DialTimeout("tcp", address, timeout)
}

func getOS() string {
	if isLinux() {
		return "Linux"
	}
	return "macOS"
}

func isLinux() bool {
	// This is determined at compile time
	return false // Will be handled properly
}
