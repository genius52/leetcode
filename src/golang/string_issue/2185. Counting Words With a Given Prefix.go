package string_issue

import "strings"

func prefixCount(words []string, pref string) int {
	var res int = 0
	for _,w := range words{
		if strings.HasPrefix(w,pref){
			res++
		}
	}
	return res
}