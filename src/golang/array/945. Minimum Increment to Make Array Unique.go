package array

import (
	"math"
	"sort"
)

//Input: [3,2,1,2,1,7]
//Output: 6
//Explanation:  After 6 moves, the array could be [3, 4, 1, 2, 5, 7].
//It can be shown with 5 or less moves that it is impossible for the array to have all unique values.

//Input: nums = [1,2,2]
//Output: 1
func MinIncrementForUnique2(A []int) int{
	var l int = len(A)
	if l < 2{
		return 0
	}
	sort.Ints(A)
	var res int = 0
	var pre int = A[0]
	var need int = A[0] + 1
	for i := 1;i < l;i++{
		if need > A[i]{
			res += (need - A[i])
		}

		if A[i] == pre{
			need++
		}else{
			if A[i] > need{
				need = A[i] + 1
			}else{
				need++
			}
		}
	}
	return res
}

func MinIncrementForUnique(A []int) int {
	var data [40000]int
	var min_num int = math.MaxInt32
	var max_num int = math.MinInt32
	for _,a := range A{
		data[a]++
		if a > max_num{
			max_num = a
		}
		if a < min_num{
			min_num = a
		}
	}
	var res int = 0
	var dup_cnt int = 0
	for i := min_num;i <= max_num;i++{
		if data[i] == 0{
			if dup_cnt > 0{
				dup_cnt--
				res += dup_cnt
			}
		}else if data[i] == 1{
			if dup_cnt > 0{
				res += dup_cnt
			}
		}else if data[i] > 1{
			dup_cnt = dup_cnt + data[i] - 1
			res += dup_cnt
		}
	}
	for dup_cnt > 0{
		dup_cnt--
		res += dup_cnt
	}
	return res
}