package array

import "sort"

func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
	return nums
}
