package number

func dp_waysToReachTarget(target int, types [][]int, l int, idx int, memo [][]int) int {
	if idx >= l {
		if target == 0 {
			return 1
		}
		return 0
	}
	if target < 0 {
		return 0
	}
	if target == 0 {
		return 1
	}
	if memo[target][idx] != -1 {
		return memo[target][idx]
	}
	var MOD int = 1e9 + 7
	var res int = 0
	for i := 0; i <= types[idx][0] && (i*types[idx][1]) <= target; i++ {
		res += dp_waysToReachTarget(target-i*types[idx][1], types, l, idx+1, memo)
		res %= MOD
	}
	memo[target][idx] = res
	return res
}

func waysToReachTarget(target int, types [][]int) int {
	var l int = len(types)
	var memo [][]int = make([][]int, target+1)
	for i := 0; i <= target; i++ {
		memo[i] = make([]int, l)
		for j := 0; j < l; j++ {
			memo[i][j] = -1
		}
	}
	return dp_waysToReachTarget(target, types, l, 0, memo)
}
