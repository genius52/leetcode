package string_issue

import "strings"

func removeOccurrences(s string, part string) string {
	var part_len int = len(part)
	for true{
		pos := strings.Index(s,part)
		if pos == -1{
			return s
		}
		s = s[:pos] + s[pos + part_len:]
	}
	return s
}