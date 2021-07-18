package string_issue

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	data := strings.Split(text, " ")
	var bad [26]bool
	for _,c := range brokenLetters {
		bad[c - 'a'] = true
	}
	var res int = 0
	var l int = len(data)
	for i := 0;i < l;i++{
		var l2 int = len(data[i])
		for j := 0;j < l2;j++{
			if bad[data[i][j] - 'a']{
				res++
				break
			}
		}
	}
	return l - res
}