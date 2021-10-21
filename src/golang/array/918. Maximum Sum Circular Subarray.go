package array

import "math"

//918
//Input: [3,-1,2,-1]
//Output: 4
//Explanation: Subarray [2,-1,3] has maximum sum 2 + (-1) + 3 = 4
//Input: [5,-3,5]
//Output: 10
//Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10
func maxSubarraySumCircular(A []int) int {
	l := len(A)
	var max int = math.MinInt32
	var min int = math.MaxInt32
	var last_sum_max int
	var cur_sum_max int
	var last_sum_min int
	var cur_sum_min int
	var sum int = 0
	for i := 0;i < l;i++{
		sum += A[i]
		if i == 0{
			cur_sum_max = A[i]
			cur_sum_min = A[i]
		}else{
			cur_sum_max = max_int(A[i],last_sum_max + A[i])
			cur_sum_min = min_int(A[i],last_sum_min + A[i])
		}
		if cur_sum_max > max{
			max = cur_sum_max
		}
		if cur_sum_min < min{
			min = cur_sum_min
		}
		last_sum_max = cur_sum_max
		last_sum_min = cur_sum_min
	}
	if sum != min{
		return max_int(max, sum - min)
	}
	return max
}
