package string_issue

import "strings"

func isSubstringPresent(s string) bool {
	var l int = len(s)
	for i := 0; i < l-1; i++ {
		sub := string(s[i+1]) + string(s[i])
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}
