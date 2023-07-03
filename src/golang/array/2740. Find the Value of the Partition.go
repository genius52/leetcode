package array

import "sort"

func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var res int = nums[l-1] - nums[0]
	for i := 1; i < l; i++ {
		res = min_int(res, nums[i]-nums[i-1])
	}
	return res
}
