package string_issue

import "strings"

func numOfStrings(patterns []string, word string) int {
	var l int = len(patterns)
	var res int = 0
	for i := 0;i < l;i++{
		ret := strings.Contains(word,patterns[i])
		if ret{
			res++
		}
	}
	return res
}