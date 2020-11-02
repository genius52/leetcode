package number

func CountVowelStrings(n int) int {
	var dp [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		dp[i] = make([]int,5)
	}
	dp[0][0] = 1//a
	dp[0][1] = 1//e
	dp[0][2] = 1//i
	dp[0][3] = 1//o
	dp[0][4] = 1//u
	for i := 1;i < n;i++{
		for j := 0;j < 5;j++{
			for k := 0;k <= j;k++{
				dp[i][j] += dp[i - 1][k]
			}
		}
	}
	return dp[n - 1][0] + dp[n - 1][1] + dp[n - 1][2] + dp[n - 1][3] + dp[n - 1][4]
}