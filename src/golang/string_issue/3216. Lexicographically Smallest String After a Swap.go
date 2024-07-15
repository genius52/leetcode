package string_issue

import "strings"

func getSmallestString3216(s string) string {
	var l int = len(s)
	var ss strings.Builder
	var swap bool = false
	i := 0
	for i < l-1 {
		cur := s[i] - '0'
		next := s[i+1] - '0'
		if swap || cur <= next {
			ss.WriteByte(s[i])
			i++
			continue
		}
		if (cur%2 == 1 && next%2 == 1) || (cur%2 == 0 && next%2 == 0) {
			ss.WriteByte(s[i+1])
			ss.WriteByte(s[i])
			i += 2
			swap = true
		} else {
			ss.WriteByte(s[i])
			i++
		}
	}
	if i < l {
		ss.WriteByte(s[i])
	}
	return ss.String()
}
