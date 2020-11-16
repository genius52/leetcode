package string_issue

import (
	"sort"
)

//Input: word1 = "cabbba", word2 = "abbccc"
//Output: true
//Explanation: You can attain word2 from word1 in 3 operations.
//Apply Operation 1: "cabbba" -> "caabbb"
//Apply Operation 2: "caabbb" -> "baaccc"
//Apply Operation 2: "baaccc" -> "abbccc"
func CloseStrings(word1 string, word2 string) bool {
	var len1 int = len(word1)
	var len2 int = len(word2)
	if len1 != len2{
		return false
	}
	var record1 [26]int
	var record2 [26]int
	for _,c := range word1{
		record1[c - 'a']++
	}
	for _,c := range word2{
		record2[c - 'a']++
	}
	var freq1 []int
	var freq2 []int
	for i := 0;i < 26;i++{
		if (record1[i] == 0 && record2[i] != 0) || (record1[i] != 0 && record2[i] == 0){
			return false
		}
		if record1[i] == 0{
			continue
		}
		freq1 = append(freq1,int(record1[i]))
		freq2 = append(freq2,int(record2[i]))
	}
	sort.Ints(freq1)
	sort.Ints(freq2)
	for i := 0;i < len(freq1);i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}
	return true
}