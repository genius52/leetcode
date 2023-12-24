package number

import "sort"

func MinimumAddedCoins(coins []int, target int) int {
	sort.Ints(coins)
	var l int = len(coins)
	//var record []bool = make([]bool, target+1)
	var cur_sum int = 0
	var idx int = 0
	var res int = 0
	for cur_sum < target && idx < l {
		if cur_sum+1 < coins[idx] {
			res++
			cur_sum += cur_sum + 1
		} else {
			cur_sum += coins[idx]
			idx++
		}
	}
	for cur_sum < target {
		cur_sum += cur_sum + 1
		res++
	}
	return res
}
