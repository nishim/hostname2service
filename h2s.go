package h2s

import "strings"

// Detect service name by hostname.
func Detect(hostname string) string {
	tailMatch := map[string]string{
		".compute.amazonaws.com": "AWS EC2",
		".cloudfront.net":        "AWS CloudFront",
	}

	for k, v := range tailMatch {
		if strings.HasSuffix(hostname, k) {
			return v
		}
	}

	return "undefined"
}
