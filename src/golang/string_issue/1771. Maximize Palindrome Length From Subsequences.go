package string_issue

//输入：word1 = "cacb", word2 = "cbba"
//输出：5
//解释：从 word1 中选出 "ab" ，从 word2 中选出 "cba" ，得到回文串 "abcba" 。
func LongestPalindrome1771(word1 string, word2 string) int {
	var l1 int = len(word1)
	var l2 int = len(word2)
	var l int = l1 + l2
	s := word1 + word2
	var dp [][]int = make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
		dp[i][i] = 1
	}
	var res int = 0
	for i := 0; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if j+1 == i {
					dp[j][i] = 2
				} else {
					dp[j][i] = 2 + dp[j+1][i-1]
				}
				if i >= l1 && j < l1 {
					res = max_int(res, dp[j][i])
				}
			} else {
				dp[j][i] = max_int(dp[j+1][i], dp[j][i-1])
			}
		}
	}
	return res
}
