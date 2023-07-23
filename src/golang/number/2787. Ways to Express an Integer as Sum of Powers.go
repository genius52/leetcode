package number

import "math"

func dfs_numberOfWays2787(n int, x int, last int, memo [][]int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if memo[n][last] != -1 {
		return memo[n][last]
	}
	var MOD int = 1e9 + 7
	var res int = 0
	for i := last + 1; int(math.Pow(float64(i), float64(x))) <= n; i++ {
		cur_num := int(math.Pow(float64(i), float64(x)))
		res += dfs_numberOfWays2787(n-cur_num, x, i, memo)
		res %= MOD
	}
	memo[n][last] = res
	return res
}

func NumberOfWays2787(n int, x int) int {
	var memo [][]int = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			memo[i][j] = -1
		}
	}
	return dfs_numberOfWays2787(n, x, 0, memo)
}
