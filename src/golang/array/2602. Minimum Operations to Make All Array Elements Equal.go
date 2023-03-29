package array

import "sort"

func MinOperations2602(nums []int, queries []int) []int64 {
	sort.Ints(nums)
	var l1 int = len(nums)
	var l2 int = len(queries)
	var queries_idx [][2]int = make([][2]int, l2)
	for i := 0; i < l2; i++ {
		queries_idx[i] = [2]int{queries[i], i}
	}
	sort.Slice(queries_idx, func(i, j int) bool {
		return queries_idx[i][0] < queries_idx[j][0]
	})
	var res []int64 = make([]int64, l2)
	var right_sum int64 = 0
	for i := 0; i < l1; i++ {
		right_sum += int64(nums[i])
	}
	var j int = 0
	var left_diff int64 = 0
	for i := 0; i < l2; i++ {
		target_num := queries_idx[i][0]
		pre_idx := queries_idx[i][1]
		if i > 0 {
			left_diff += int64((target_num - queries_idx[i-1][0]) * j)
		}
		for j < l1 && nums[j] <= target_num {
			left_diff += int64(abs_int(nums[j] - target_num))
			right_sum -= int64(nums[j])
			j++
		}
		right_diff := right_sum - int64(target_num*(l1-j))
		res[pre_idx] = left_diff + right_diff
	}
	return res
}
