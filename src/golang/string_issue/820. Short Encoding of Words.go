package string_issue

import (
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	sort.Sort(sortString(words))
	var l int = len(words)
	var encode string = words[l - 1] + "#"
	for i := l - 2;i >= 0;i--{
		var res = strings.Contains(encode,words[i] + "#")
		if !res{
			encode += words[i] + "#"
		}
	}
	return len(encode)
}