package array

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var res int = 0
	for i := 0;i < l / 2;i++{
		cur := nums[i] + nums[l - 1 - i]
		if cur > res{
			res = cur
		}
	}
	return res
}