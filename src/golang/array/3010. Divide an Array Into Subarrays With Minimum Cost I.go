package array

import "sort"

func minimumCost3010(nums []int) int {
	sort.Ints(nums[1:])
	return nums[0] + nums[1] + nums[2]
}
