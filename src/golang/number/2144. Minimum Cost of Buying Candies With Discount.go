package number

import "sort"

func MinimumCost(cost []int) int {
	sort.Ints(cost)
	var l int = len(cost)
	var res int = 0
	var i int = l - 1
	for ;i - 1 >= 0;i -= 3{
		res += cost[i] + cost[i - 1]
	}
	for i >= 0{
		res += cost[i]
		i--
	}
	return res
}