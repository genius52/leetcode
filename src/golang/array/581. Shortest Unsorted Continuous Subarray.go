package array

import "sort"

func findUnsortedSubarray(nums []int) int {
	var l int = len(nums)
	var sorted []int = make([]int,l)
	copy(sorted,nums)
	sort.Ints(sorted)
	var left int = 0
	var right int = l - 1
	for left < l && nums[left] == sorted[left]{
		left++
	}
	if left == l{
		return 0
	}
	for right >= 0 && nums[right] == sorted[right]{
		right--
	}
	return right - left + 1
}