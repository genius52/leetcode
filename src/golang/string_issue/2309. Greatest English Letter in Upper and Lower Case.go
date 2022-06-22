package string_issue

import "strings"

func greatestLetter(s string) string {
	var l int = len(s)
	var lower [26]bool
	var upper [26]bool
	var res string
	var biggest int = -1
	for i := 0;i < l;i++{
		if s[i] >= 'a' && s[i] <= 'z'{
			idx := s[i] - 'a'
			lower[idx] = true
			if lower[idx] && upper[idx] && int(idx) > biggest{
				biggest = int(idx)
				res = strings.ToUpper(string(s[i]))
			}
		}
		if s[i] >= 'A' && s[i] <= 'Z'{
			idx := s[i] - 'A'
			upper[idx] = true
			if lower[idx] && upper[idx] && int(idx) > biggest{
				biggest = int(idx)
				res = string(s[i])
			}
		}
	}
	return res
}