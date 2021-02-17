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