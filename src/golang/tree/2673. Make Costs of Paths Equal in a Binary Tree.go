package tree

func abs_int(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func recrusive_minIncrements(cur int, n int, cost []int, res *int) int {
	if cur > n {
		return 0
	}
	left := cur * 2
	right := cur*2 + 1
	left_sum := recrusive_minIncrements(left, n, cost, res)
	right_sum := recrusive_minIncrements(right, n, cost, res)
	*res += abs_int(left_sum - right_sum)
	return cost[cur-1] + max_int(left_sum, right_sum)
}

func MinIncrements(n int, cost []int) int {
	var res int = 0
	recrusive_minIncrements(1, n, cost, &res)
	return res
}
