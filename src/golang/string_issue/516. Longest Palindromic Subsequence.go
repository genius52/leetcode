package string_issue

import "math"

//Input:
//"bbbab"
//Output:
//4

//DP from bottom to top
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

//DP from top to bottom
func dp_longestPalindromeSubseq(s string,l int,left int,right int,memo [][]int)int{
	if left == right{
		return 1
	}
	if left > right{
		return 0
	}
	if memo[left][right] != -1{
		return memo[left][right]
	}
	var res int = 0
	if s[left] == s[right]{
		res = 2 + dp_longestPalindromeSubseq(s,l,left + 1,right - 1,memo)
	}else{
		res = max_int(dp_longestPalindromeSubseq(s,l,left + 1,right,memo),dp_longestPalindromeSubseq(s,l,left,right - 1,memo))
	}
	memo[left][right] = res
	return res
}

func longestPalindromeSubseq(s string) int{
	var l int = len(s)
	var memo [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		memo[i] = make([]int,l)
		for j := 0;j < l;j++{
			memo[i][j] = -1
		}
	}
	return dp_longestPalindromeSubseq(s,l,0,l - 1,memo)
}