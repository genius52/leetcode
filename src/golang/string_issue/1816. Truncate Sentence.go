package string_issue

import "strings"

func truncateSentence(s string, k int) string {
	var data []string = strings.Split(s," ")
	data = data[0:k]
	var res string = strings.Join(data," ")
	return res
}