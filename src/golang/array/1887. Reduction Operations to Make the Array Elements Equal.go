package array

import "sort"

func ReductionOperations(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var res int = 0
	for i := l - 1; i > 0; i-- {
		if nums[i] == nums[0] {
			break
		}
		if nums[i] != nums[i-1] {
			res += l - i
		}
	}
	return res
}
