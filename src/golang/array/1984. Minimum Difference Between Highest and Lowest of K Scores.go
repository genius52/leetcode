package array

import "sort"

func minimumDifference(nums []int, k int) int {
	if k == 0{
		return 0
	}
	var l int = len(nums)
	sort.Ints(nums)
	var res int = 2147483647
	for i := 0;i < l && (i + k - 1) < l;i++{
		cur := nums[i + k - 1] - nums[i]
		if cur < res{
			res = cur
		}
	}
	return res
}