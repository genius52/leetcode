package string_issue

import "strings"

func findOcurrences(text string, first string, second string) []string {
	var res []string
	var words []string = strings.Split(text," ")
	if len(words) <= 1{
		return res
	}
	for slow := 0; slow < len(words) - 2;slow++{
		if words[slow] == first{
			if words[slow+1] == second{
				res = append(res, words[slow+2])
			}
		}
	}
	return res
}
