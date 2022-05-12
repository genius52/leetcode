package string_issue

import "strings"

func lengthOfLastWord(s string) int {
	var data []string = strings.Split(strings.TrimSpace(s)," ")
	var l int = len(data)
	if l == 0{
		return 0
	}
	return len(data[l - 1])
}

func lengthOfLastWord2(s string) int{
	var l int = len(s)
	var right int = l - 1
	for right >= 0{
		if s[right] != ' '{
			break
		}
		right--
	}
	var left int = right - 1
	for left >= 0{
		if s[left] == ' '{
			break
		}
		left--
	}
	return right - left
}