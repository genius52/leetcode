package number

import "sort"

func GetMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	var l int = len(coins)
	var cur int = 0
	for i := 0; i < l; i++ {
		if coins[i] > cur+1 {
			break
		}
		cur += coins[i]
	}
	return cur + 1
}
