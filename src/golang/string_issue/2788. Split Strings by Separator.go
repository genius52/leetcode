package string_issue

import "strings"

func splitWordsBySeparator(words []string, separator byte) []string {
	var res []string
	for _, word := range words {
		cur := strings.Split(word, string(separator))
		for _, w := range cur {
			if len(w) > 0 {
				res = append(res, w)
			}
		}
	}
	return res
}
