package number

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var res int = 0
	for high := l - 1;high > 1;high--{
		if (nums[high - 2] + nums[high - 1]) > nums[high]{
			sum := nums[high - 2] + nums[high - 1] + nums[high]
			if sum > res{
				return sum
			}
		}
	}
	return res
}