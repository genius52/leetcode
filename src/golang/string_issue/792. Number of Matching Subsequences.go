package string_issue

//Input:
//S = "abcde"
//words = ["a", "bb", "acd", "ace"]
//Output: 3
//Explanation: There are three words in words that are a subsequence of S: "a", "acd", "ace".
func NumMatchingSubseq(S string, words []string) int {
	var l int = len(S)
	var res int = 0
	for _,word := range words{
		word_len := len(word)
		if word_len > l{
			continue
		}
		var i int = 0
		var j int = 0
		for i < word_len && j < l{
			if word[i] == S[j] {
				i++
				j++
			}else{
				j++
			}
		}
		if i == word_len{
			res++
		}
	}
	return res
}