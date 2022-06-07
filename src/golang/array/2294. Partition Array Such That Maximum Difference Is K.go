package array

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	var res int = 0
	var l int = len(nums)
	var left int = 0
	for left < l{
		var right int = left + 1
		for right < l && nums[right] - nums[left] <= k{
			right++
		}
		left = right
		res++
	}
	return res
}