package array

import "sort"

func minimumOperations2357(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var res int = 0
	var cur int = 0
	for i := 0;i < l;i++{
		if nums[i] == cur{
			continue
		}
		res++
		cur = nums[i]
	}
	return res
}