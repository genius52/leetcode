package string_issue

import "strings"

func isCircularSentence(sentence string) bool {
	var s []string = strings.Split(sentence," ")
	var l int = len(s)
	for i := 0;i < l;i++{
		if s[i][len(s[i]) - 1] != s[(i + 1)%l][0]{
			return false
		}
	}
	return true
}