package array

import "sort"

func isGood(nums []int) bool {
	var l int = len(nums)
	sort.Ints(nums)
	if nums[l-1] != l-1 {
		return false
	}
	for i := 0; i < l-1; i++ {
		if nums[i] != (i + 1) {
			return false
		}
	}
	return true
}
