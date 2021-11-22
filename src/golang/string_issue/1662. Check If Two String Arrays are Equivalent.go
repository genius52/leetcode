package string_issue

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var s1 string = strings.Join(word1,"")
	var s2 string = strings.Join(word2,"")
	return s1 == s2
}

func arrayStringsAreEqual2(word1 []string, word2 []string) bool{
	var i int = 0
	var j int = 0
	var visit1 int = 0
	var visit2 int = 0
	var l1 int = len(word1)
	var l2 int = len(word2)
	for i < l1 && j < l2{
		for visit1 < len(word1[i]) && visit2 < len(word2[j]){
			if word1[i][visit1] != word2[j][visit2]{
				return false
			}
			visit1++
			visit2++
		}
		if visit1 >= len(word1[i]){
			i++
			visit1 = 0
		}
		if visit2 >= len(word2[j]){
			j++
			visit2 = 0
		}
	}
	return i == l1 && j == l2
}