package string_issue

import "math"

//Input: text1 = "abcde", text2 = "ace"
//Output: 3
//Explanation: The longest common subsequence is "ace" and its length is 3.
func LongestCommonSubsequence(text1 string, text2 string) int {
	var l1 int = len(text1)
	var l2 int = len(text2)
	var dp [][]int = make([][]int,l1 + 1)//dp[i][j] = the length of longest from text1[0:i] and text2[0:j]
	for i := 0;i <= l1;i++{
		dp[i] = make([]int,l2 + 1)
	}
	var res int = 0
	//dp[0][0] = 1 dp[0][1] dp[1][0] dp[1][1]
	for i := 1;i <= l1;i++{
		for j := 1;j <= l2;j++{
			if text1[i - 1] == text2[j - 1]{
				dp[i][j] = 1 + dp[i - 1][j - 1]
				if dp[i][j] > res{
					res = dp[i][j]
				}
			}else{
				dp[i][j] = int(math.Max(float64(dp[i - 1][j]),float64(dp[i][j - 1])))
			}
		}
	}
	return res
}