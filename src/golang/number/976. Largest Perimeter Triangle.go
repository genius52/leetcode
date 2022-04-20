package number

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	for high := l - 1;high > 1;high--{
		if (nums[high - 2] + nums[high - 1]) > nums[high]{
			return nums[high - 2] + nums[high - 1] + nums[high]
		}
	}
	return 0
}