package string_issue

import "strings"

func SmallestString(s string) string {
	var l int = len(s)
	var ss strings.Builder
	var left int = 0
	for left < l {
		if s[left] == 'a' {
			ss.WriteString(string(s[left]))
			left++
		} else {
			break
		}
	}
	var replace bool = false
	for left < l && s[left] != 'a' {
		ss.WriteString(string(s[left] - 1))
		left++
		replace = true
	}
	for left < l {
		ss.WriteString(string(s[left]))
		left++
	}
	if !replace {
		return s[:l-1] + "z"
	}
	return ss.String()
}
