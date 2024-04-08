package number

import "sort"

func MinOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)
	var l int = len(nums)
	var decrease_cnt int64 = 0
	var idx1 int = l/2 - 1
	for ; idx1 >= 0; idx1-- {
		if nums[idx1] <= k {
			break
		}
		decrease_cnt += int64(nums[idx1] - k)
	}
	var increase_cnt int64 = 0
	var idx2 int = l/2 + 1
	for ; idx2 < l; idx2++ {
		if nums[idx2] > k {
			break
		}
		increase_cnt += int64(k - nums[idx2])
	}
	var mid_idx int = l / 2
	return abs_int64(decrease_cnt-increase_cnt) + abs_int64(int64(nums[mid_idx]-k))
}
