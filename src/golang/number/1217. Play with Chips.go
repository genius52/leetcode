package number

import "math"

//Input: chips = [2,2,2,3,3]
//Output: 2
//Explanation: Both fourth and fifth chip will be moved to position two with cost 1. Total minimum cost will be 2.
func MinCostToMoveChips(chips []int) int {
	var one_cnt int = 0
	var two_cnt int = 0
	for _, n := range chips{
		if(n % 2 == 1){
			one_cnt++
		}else{
			two_cnt++
		}
	}
	return int(math.Min(float64(one_cnt),float64(two_cnt)))
}
