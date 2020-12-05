package array

import "math"

//Input: [8,1,5,2,6]
//Output: 11
//Explanation: i = 0, j = 2, A[i] + A[j] + i - j = 8 + 5 + 0 - 2 = 11
func MaxScoreSightseeingPair(A []int) int {
	var l int = len(A)
	var add_index []int = make([]int,l)
	var minus_index []int = make([]int,l)
	for i := 0;i < l;i++{
		add_index[i] = A[i] + i
		minus_index[i] = A[i] - i
	}
	var left_max []int = make([]int,l)
	left_max[0] = add_index[0]
	var right_max []int = make([]int,l)
	right_max[l - 1] = minus_index[l - 1]
	for i := 1;i < l;i++{
		left_max[i] = int(math.Max(float64(add_index[i]),float64(left_max[i - 1])))
		right_max[l - 1 - i] = int(math.Max(float64(right_max[l - i]),float64(minus_index[l - 1 - i])))
	}
	var max_val int = 0
	for i := 0;i < l - 1;i++{
		cur_val := left_max[i] + right_max[i + 1]
		if cur_val > max_val{
			max_val = cur_val
		}
	}
	return max_val
}