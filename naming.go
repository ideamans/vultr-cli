package main

import (
	"strings"
	"unicode"
)

// kebab converts a Go identifier like "BareMetalServer" or "CreateIPv4"
// into a CLI-friendly name like "bare-metal-server" or "create-ipv4".
func kebab(s string) string {
	words := splitCamel(s)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, "-")
}

// splitCamel splits a camel/pascal-case identifier into words, keeping
// acronym runs (SSH, VPC2, ISO) and version markers (IPv4, IPv6) intact.
func splitCamel(s string) []string {
	r := []rune(s)
	n := len(r)
	splitAt := map[int]bool{}
	for i := 1; i < n; i++ {
		prev := r[i-1]
		cur := r[i]
		if unicode.IsUpper(cur) && (unicode.IsLower(prev) || unicode.IsDigit(prev)) {
			// lower/digit -> Upper boundary: "createIPv4" -> split before I
			splitAt[i] = true
		} else if unicode.IsLower(cur) && i >= 2 && unicode.IsUpper(prev) && unicode.IsUpper(r[i-2]) {
			// end of an acronym run: "VPCInfo" -> split before "Info"
			// exception: 'v' followed by a digit keeps "IPv4"/"IPv6" together
			if !(cur == 'v' && i+1 < n && unicode.IsDigit(r[i+1])) {
				splitAt[i-1] = true
			}
		}
	}
	var words []string
	start := 0
	for i := 1; i <= n; i++ {
		if i == n || splitAt[i] {
			words = append(words, string(r[start:i]))
			start = i
		}
	}
	return words
}

// typeName renders a reflect type name without the govultr package prefix.
func typeName(s string) string {
	return strings.ReplaceAll(s, "govultr.", "")
}
