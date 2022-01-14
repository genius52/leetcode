package number

func dp_rearrangeSticks(n int, k int, memo [][]int) int {
	if n <= 0 || k <= 0 || k > n {
		return 0
	}
	if n == k {
		return 1
	}
	if memo[n][k] != -1 {
		return memo[n][k]
	}
	var res int = 0
	var MOD int = 1e9 + 7
	//选最短的木棍放在最前面
	res += dp_rearrangeSticks(n-1, k-1, memo)
	res += (n - 1) * dp_rearrangeSticks(n-1, k, memo)
	memo[n][k] = res % MOD
	return memo[n][k]
}

func RearrangeSticks(n int, k int) int {
	if n == k {
		return 1
	}
	var memo [][]int = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			memo[i][j] = -1
		}
	}
	return dp_rearrangeSticks(n, k, memo)
}
