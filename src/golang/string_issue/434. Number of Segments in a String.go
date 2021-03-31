package string_issue

import "strings"

func countSegments(s string) int {
	s = strings.Trim(s," ")
	var l int = len(s)
	if l == 0{
		return 0
	}
	var cnt int = 1
	var left int = 0
	for left < l{
		var right int = left
		for right < l && s[right] != ' '{
			right++
		}
		if right < l{
			cnt++
		}
		for right < l && s[right] == ' '{
			right++
		}
		left = right
	}
	return cnt
}