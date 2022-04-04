package number

func climbStairs(n int) int {
	if n == 1{
		return 1
	}
	var dp []int = make([]int,n + 1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3;i <= n;i++{
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}

func climbStairs2(n int) int {
	if n == 1{
		return 1
	}
	if n == 2{
		return 2
	}
	var onestep_before int = 2//start with n == 2
	var twostep_before int = 1//start with n == 1
	var cur int = 0
	for i := 3;i <= n;i++{
		cur = onestep_before + twostep_before
		twostep_before = onestep_before
		onestep_before = cur
	}
	return cur
}