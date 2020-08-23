package string_issue

import "math"

//Input:
//"bbbab"
//Output:
//4
func LongestPalindromeSubseq(s string) int {
	var l int = len(s)
	var dp [][]int = make([][]int,l)//dp[i][j]: 从i到j最大的回文长度
	for i := 0;i < l;i++{
		dp[i] = make([]int,l)
		dp[i][i] = 1
	}
	//if s[j] == s[i] dp[i][j] = 1 + max(dp[i][j-1],dp[i][j-2]...dp[i][i])
	for left := l - 1;left >= 0;left-- {
		for right := left + 1;right < l;right++{
			if s[left] == s[right]{
				dp[left][right] = 2 + dp[left + 1][right - 1]
			}else{
				dp[left][right] = int(math.Max(float64(dp[left][right - 1]),float64(dp[left + 1][right])))
			}
		}
	}
	return dp[0][l - 1]
}