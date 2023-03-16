package array

import "sort"

func MaxScore(nums []int) int {
	var l int = len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var sum int = 0
	for i := 0; i < l; i++ {
		sum += nums[i]
		if sum <= 0 {
			return i
		}
	}
	return l
}
