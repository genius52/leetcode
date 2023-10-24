package string_issue

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	var res []string = []string{words[0]}
	for i := 1; i < n; i++ {
		if groups[i] != groups[i-1] {
			res = append(res, words[i])
		}
	}
	return res
}
