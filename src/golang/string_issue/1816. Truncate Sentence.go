package string_issue

import "strings"

func truncateSentence2(s string, k int) string {
	var res strings.Builder
	for i := 0; i < len(s) && k > 0; i++ {
		if s[i] == ' ' {
			k--
			if k == 0 {
				break
			}
		}
		res.WriteByte(s[i])
	}
	return res.String()
}

func truncateSentence(s string, k int) string {
	var data []string = strings.Split(s, " ")
	data = data[0:k]
	var res string = strings.Join(data, " ")
	return res
}
