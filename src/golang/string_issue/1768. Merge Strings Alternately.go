package string_issue

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var l1 int = len(word1)
	var l2 int = len(word2)
	var i int = 0
	var j int = 0
	var data strings.Builder
	for i < l1 || j < l2 {
		if i < l1{
			data.WriteByte(word1[i])
			i++
		}
		if j < l2{
			data.WriteByte(word2[j])
			j++
		}
	}
	return data.String()
}