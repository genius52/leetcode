package array

func dp_minimumCoins(prices []int, l int, idx int, memo []int) int {
	if idx >= l {
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}
	var res int = 1<<31 - 1
	for i := idx + 1; i <= idx+idx+2; i++ {
		cur := prices[idx] + dp_minimumCoins(prices, l, i, memo)
		if cur < res {
			res = cur
		}
	}
	if res != 1<<31-1 {
		memo[idx] = res
	}
	return res
}

func MinimumCoins(prices []int) int {
	var l int = len(prices)
	if l == 1 {
		return prices[0]
	}
	var memo []int = make([]int, l)
	for i := 0; i < l; i++ {
		memo[i] = -1
	}
	return dp_minimumCoins(prices, l, 0, memo)
}
