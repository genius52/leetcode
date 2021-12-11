package number

import "sort"

func dfs_closestCost(toppingCosts []int, l int, idx int, cur int, target int, res *int) {
	if idx >= l {
		diff := abs_int(target - cur)
		if diff == abs_int(target-*res) {
			if cur < *res {
				*res = cur
			}
		} else if diff < abs_int(target-*res) {
			*res = cur
		}
		return
	}
	//skip current
	dfs_closestCost(toppingCosts, l, idx+1, cur, target, res)
	//choose current one piece
	dfs_closestCost(toppingCosts, l, idx+1, cur+toppingCosts[idx], target, res)
	//choose current two piece
	dfs_closestCost(toppingCosts, l, idx+1, cur+toppingCosts[idx]*2, target, res)
}

func ClosestCost(baseCosts []int, toppingCosts []int, target int) int {
	sort.Ints(baseCosts)
	sort.Ints(toppingCosts)
	var l1 int = len(baseCosts)
	var l2 int = len(toppingCosts)
	//var min_diff int = 2147483647
	var res int = 2147483647
	for i := 0; i < l1; i++ {
		dfs_closestCost(toppingCosts, l2, 0, baseCosts[i], target, &res)
	}
	return res
}
