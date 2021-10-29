package array

import "sort"

func MinDifference1509(nums []int) int {
	var l int = len(nums)
	if l <= 4{
		return 0
	}
	sort.Ints(nums)
	var res int = 2147483647
	for i := 0;i <= 3;i++{
		res = min_int(res,nums[l - 4 + i] - nums[i])
	}
	return res
}