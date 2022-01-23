package array

import "sort"

func countElements(nums []int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var min_val int = 2147483647
	var max_val int = -2147483648
	for i := 0;i < l;i++{
		if nums[i] < min_val{
			min_val = nums[i]
		}
		if nums[i] > max_val{
			max_val = nums[i]
		}
	}
	var res int = 0
	for i := 0;i < l;i++{
		if nums[i] > min_val && nums[i] < max_val{
			res++
		}
	}
	return res
}