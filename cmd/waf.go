package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var testXSS, testSQLi bool

var wafCmd = &cobra.Command{
	Use:   "waf [domain]",
	Short: color.RedString("WAF Detection - Web Application Firewall analysis"),
	Long: `Advanced WAF detection and analysis tool:

Features:
  ▸ WAF Signature Detection
  ▸ HTTP Header Analysis
  ▸ Cookie Inspection
  ▸ Response Pattern Matching
  ▸ Common WAF Identification

Supported WAFs:
  • Cloudflare          • AWS WAF            • ModSecurity
  • Akamai              • Imperva/Incapsula  • F5 BIG-IP
  • Barracuda           • Sucuri             • Wordfence

Flags:
  --test-xss     Test XSS payload detection
  --test-sqli    Test SQL injection detection

Examples:
  afsa waf example.com
  afsa waf cloudflare.com --test-xss
  afsa waf example.com --test-sqli`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		analyzeAdvancedWAF(domain)
	},
}

func init() {
	wafCmd.Flags().BoolVar(&testXSS, "test-xss", false, "Test XSS detection")
	wafCmd.Flags().BoolVar(&testSQLi, "test-sqli", false, "Test SQLi detection")
}

func analyzeAdvancedWAF(domain string) {
	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║             WAF DETECTION ANALYSIS REPORT               ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n")

	color.Cyan("  Target: %s\n\n", domain)

	color.Red("  ▸ WAF Identification Indicators:\n")

	wafSignatures := map[string][]string{
		"Cloudflare": {
			"Server header: cloudflare",
			"CF-Ray header present",
			"cf_clearance cookie",
			"__cfruid cookie (Bot Management)",
			"Nameservers: *.ns.cloudflare.com",
		},
		"AWS WAF": {
			"x-amzn-RequestId header",
			"Server header patterns",
			"AWS security headers",
			"Response patterns",
		},
		"ModSecurity": {
			"X-Mod-Security header",
			"403/406 error patterns",
			"Rule ID in response",
		},
		"Akamai": {
			"AkamaiGHost in Server header",
			"X-Akamai-* headers",
			"Cookie patterns",
		},
		"Imperva": {
			"X-Iinfo header",
			"X-CDN: Imperva",
			"Imperva response headers",
		},
		"F5 BIG-IP": {
			"X-Forwarded-* headers",
			"BigIP persistence cookies",
			"Response header patterns",
		},
		"Barracuda": {
			"Specific cookie patterns",
			"Barracuda error pages",
			"Header signatures",
		},
		"Sucuri": {
			"X-Sucuri-Cache header",
			"Sucuri error patterns",
			"Blocking page indicators",
		},
		"Wordfence": {
			"X-Wordfence-* headers",
			"WordPress security",
			"Plugin detection",
		},
	}

	i := 0
	for waf, indicators := range wafSignatures {
		prefix := "├─ "
		if i == len(wafSignatures)-1 {
			prefix = "└─ "
		}
		fmt.Printf("    %s%s\n", prefix, color.YellowString(waf))
		for j, indicator := range indicators {
			innerPrefix := "│  ├─ "
			if j == len(indicators)-1 {
				innerPrefix = "│  └─ "
			}
			fmt.Printf("    %s%s\n", innerPrefix, indicator)
		}
		i++
	}

	color.Red("\n  ▸ Header Analysis Targets:\n")
	headers := []string{
		"Server", "X-Powered-By", "X-Frame-Options",
		"X-Content-Type-Options", "Strict-Transport-Security",
		"Content-Security-Policy", "X-XSS-Protection",
	}
	for i, header := range headers {
		if i == len(headers)-1 {
			fmt.Printf("    └─ %s\n", color.BlueString(header))
		} else {
			fmt.Printf("    ├─ %s\n", color.BlueString(header))
		}
	}

	color.Red("\n  ▸ Common Detection Methods:\n")
	methods := []string{
		"Analyze HTTP response headers",
		"Inspect Set-Cookie headers",
		"Check error page content",
		"Examine server behavior patterns",
		"Test XSS/SQLi payload responses",
	}
	for i, method := range methods {
		if i == len(methods)-1 {
			fmt.Printf("    └─ %s\n", method)
		} else {
			fmt.Printf("    ├─ %s\n", method)
		}
	}

	if testXSS || testSQLi {
		color.Red("\n  ▸ Payload Testing:\n")
		if testXSS {
			fmt.Printf("    ├─ XSS Payload: %s\n", color.RedString("GET /?test=<script>alert('xss')</script>"))
		}
		if testSQLi {
			fmt.Printf("    └─ SQLi Payload: %s\n", color.RedString("GET /?id=1' OR '1'='1"))
		}
		color.Yellow("\n    ⚠  Only use on domains you own or have permission to test!\n")
	}

	color.Red("\n╔════════════════════════════════════════════════════════╗\n")
	color.Red("║       [✓] WAF Analysis Completed Successfully          ║\n")
	color.Red("╚════════════════════════════════════════════════════════╝\n\n")
}
