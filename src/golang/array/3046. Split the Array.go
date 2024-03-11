package array

import "sort"

func isPossibleToSplit(nums []int) bool {
	var l int = len(nums)
	if l|1 == l {
		return false
	}
	sort.Ints(nums)
	for i := 1; i < l; i++ {
		if nums[i] == nums[i-1] {
			if i > 1 && nums[i] == nums[i-2] {
				return false
			}
		}
	}
	return true
}
