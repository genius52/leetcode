package number

import "sort"

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	return nums[l-1]*nums[l-2] - nums[0]*nums[1]
}

func maxProductDifference2(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var small1 smal
	return nums[l-1]*nums[l-2] - nums[0]*nums[1]
}
