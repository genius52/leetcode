package number

import "math"

//Input: chips = [2,2,2,3,3]
//Output: 2
//Explanation: Both fourth and fifth chip will be moved to position two with cost 1. Total minimum cost will be 2.
func MinCostToMoveChips(chips []int) int {
	var cnt int = 0
	var l int = len(chips)
	for i := 0;i < l;i++{
		if (chips[i] | 1) == chips[i]{
			cnt++
		}
	}
	return int(math.Min(float64(cnt),float64(l - cnt)))
}
