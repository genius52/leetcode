package string_issue

import "strings"

func makeFancyString2(s string) string{
	var l int = len(s)
	var sb strings.Builder
	var left int = 0
	for left < l{
		right := left + 1
		for right < l && s[left] == s[right]{
			right++
		}
		cur_len := right - left
		if cur_len <= 2{
			sb.WriteString(s[left:right])
		}else{
			sb.WriteString(s[left:left + 2])
		}
		left = right
	}
	return sb.String()
}

func makeFancyString(s string) string {
	var l int = len(s)
	var left int = 0
	var res string
	for left < l{
		var right int = left + 1
		for right < l && s[left] == s[right]{
			right++
		}
		cnt := right - left
		if cnt >= 3{
			cnt = 2
		}
		for cnt > 0{
			res += string(s[left])
			cnt--
		}
		left = right
	}
	return res
}