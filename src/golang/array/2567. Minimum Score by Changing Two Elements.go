package array

import "sort"

func minimizeSum(nums []int) int {
	var l int = len(nums)
	sort.Ints(nums)
	return min_int_number(nums[l-2]-nums[1], nums[l-1]-nums[2], nums[l-3]-nums[0])
}
