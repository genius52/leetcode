package string_issue

func longestSubsequenceRepeatedK(s string, k int) string {
	var freq [26]int
	var l int = len(s)
	for i := 0;i < l;i++{
		freq[s[i] - 'a']++
	}
	return ""
}