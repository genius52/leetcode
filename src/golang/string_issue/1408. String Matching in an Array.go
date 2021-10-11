package string_issue

import (
	"sort"
	"strings"
)

func StringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	var res []string
	var record map[string]bool = make(map[string]bool)
	var l int = len(words)
	for i := 0;i < len(words[l - 1]);i++{
		record[words[l - 1][i:]] = true
	}
	for i := l - 2;i >= 0;i--{
		for s,_ := range record{
			if strings.Contains(s,words[i]){
				res = append(res,words[i])
				break
			}
		}
		for j := 0;j < len(words[i]);j++{
			record[words[i][j:]] = true
		}
	}
	return res
}