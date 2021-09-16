package string_issue

//Input: text = "abcabcabc"
//Output: 3
//Explanation: The 3 substrings are "abcabc", "bcabca" and "cabcab".
func DistinctEchoSubstrings(text string) int {
	var l int = len(text)
	var record map[string]bool = make(map[string]bool)//保存符合的子串 abab
	for i := 0;i < l;i++{
		for j := 2;(i + j) <= l;j += 2{
			s1 := text[i:i + j/2]
			s2 := text[i + j/2:i + j]
			if s1 == s2{
				if _,ok := record[s1 + s2];!ok{
					record[s1 + s2] = true
				}
			}
		}
	}
	return len(record)
}