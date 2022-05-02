package string_issue

import "strings"

func countPrefixes(words []string, s string) int {
	var cnt int = 0
	for _,w := range words{
		if strings.HasPrefix(s,w){
			cnt++
		}
	}
	return cnt
}