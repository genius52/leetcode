package string_issue

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var l1 int = len(word1)
	var l2 int = len(word2)
	var i int = 0
	var j int = 0
	var data []string = make([]string, l1+l2)
	var idx int = 0
	for i < l1 && j < l2 {
		data[idx] = string(word1[i])
		idx++
		data[idx] = string(word2[j])
		idx++
	}
	for i < l1 {
		data[idx] = string(word1[i])
		i++
		idx++
	}
	for j < l2 {
		data[idx] = string(word2[j])
		j++
		idx++
	}
	return strings.Join(data, "")
}
