package string_issue

import "strings"

func countPrefixSuffixPairs(words []string) int {
	var l int = len(words)
	var cnt int = 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if strings.HasPrefix(words[j], words[i]) && strings.HasSuffix(words[j], words[i]) {
				cnt++
			}
		}
	}
	return cnt
}
