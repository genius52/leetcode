package string_issue

import "math"

//Input: "sea", "eat"
//Output: 2
//Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
func max_int_number(nums ...int)int{
	var max int = math.MinInt32
	for _,n := range nums{
		if n > max{
			max = n
		}
	}
	return max
}

func MinDistance(word1 string, word2 string) int {
	var l1 int = len(word1)
	var l2 int = len(word2)
	var dp [][]int = make([][]int,l1 + 1)//dp[i][j] 从i到j的最大相同子序列
	for i := 0;i <= l1;i++{
		dp[i] = make([]int,l2 + 1)
	}
	for i := 0;i < l1;i++ {
		for j := 0; j < l2; j++ {
			if word1[i] == word2[j]{
				dp[i + 1][j + 1] = 1 + dp[i][j]
			}else{
				dp[i + 1][j + 1] = max_int_number(dp[i + 1][j],dp[i][j + 1],dp[i + 1][j + 1])
			}
		}
	}
	return l1 - dp[l1][l2] + l2 - dp[l1][l2]
}