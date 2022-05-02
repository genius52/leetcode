package array

import "math"

//Input: [8,1,5,2,6]
//Output: 11
//Explanation: i = 0, j = 2, A[i] + A[j] + i - j = 8 + 5 + 0 - 2 = 11
func MaxScoreSightseeingPair2(values []int) int{
	var l int = len(values)
	var res int = 0
	var pre_max int = values[0]
	for i := 1;i < l;i++{
		cur := pre_max + values[i] - i
		if cur > res{
			res = cur
		}
		if pre_max < (values[i] + i){
			pre_max = values[i] + i
		}
	}
	return res
}

func MaxScoreSightseeingPair(A []int) int {
	var l int = len(A)
	var add_index []int = make([]int,l) //add_index[i]:如果左边选择位置i的得分
	var minus_index []int = make([]int,l)//minus_index[i]:如果右边选择位置i的得分
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