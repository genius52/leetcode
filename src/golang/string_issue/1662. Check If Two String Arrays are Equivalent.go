package string_issue

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var s1 string = strings.Join(word1,"")
	var s2 string = strings.Join(word2,"")
	return s1 == s2
}