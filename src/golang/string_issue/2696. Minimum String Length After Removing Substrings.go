package string_issue

import "strings"

func MinLength(s string) int {
	var change bool = true
	for change {
		change = false
		for strings.Contains(s, "AB") {
			change = true
			idx := strings.Index(s, "AB")
			s = s[:idx] + s[idx+2:]
		}
		for strings.Contains(s, "CD") {
			change = true
			idx := strings.Index(s, "CD")
			s = s[:idx] + s[idx+2:]
		}
	}
	return len(s)
}
