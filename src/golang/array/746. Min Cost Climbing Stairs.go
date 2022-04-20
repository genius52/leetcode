package array

func dp_minCostClimbingStairs(cost []int,l int,pos int,memo []int)int{
	if pos >= l{
		return 0
	}
	if memo[pos] != -1{
		return memo[pos]
	}
	memo[pos] = min_int(cost[pos] + dp_minCostClimbingStairs(cost,l,pos + 1,memo),cost[pos] + dp_minCostClimbingStairs(cost,l,pos + 2,memo))
	return memo[pos]
}

//Top to bottom
func minCostClimbingStairs(cost []int) int {
	var l int = len(cost)
	var memo []int = make([]int,l)
	for i := 0;i < l;i++{
		memo[i] = -1
	}
	return min_int(dp_minCostClimbingStairs(cost,l,0,memo),dp_minCostClimbingStairs(cost,l,1,memo))
}

//Bottom to top
func MinCostClimbingStairs2(cost []int) int {
	var l int = len(cost)
	if l <= 1{
		return 0
	}
	var dp []int = make([]int,l)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2;i < l;i++{
		dp[i] = min_int(cost[i] + dp[i - 2],cost[i] + dp[i - 1])
	}
	return min_int(dp[l - 1],dp[l - 2])
}

func minCostClimbingStairs2(cost []int) int {
	var l int = len(cost)
	if l <= 1{
		return 0
	}
	var cost_last_last int = cost[0]
	var cost_last int = cost[1]
	for i := 2;i < l;i++{
		cur := min_int(cost_last_last + cost[i],cost_last + cost[i])
		cost_last_last = cost_last
		cost_last = cur
	}
	return min_int(cost_last,cost_last_last)
}