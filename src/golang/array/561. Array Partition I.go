package array

import "sort"

func arrayPairSum(nums []int) int {
	var l int = len(nums)
	sort.Ints(nums)
	var res int = 0
	for i := 0;i < l;{
		res += nums[i]
		i += 2
	}
	return res
}