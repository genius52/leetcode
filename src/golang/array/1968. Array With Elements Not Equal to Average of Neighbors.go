package array

import "sort"

func rearrangeArray1968(nums []int) []int {
	var length int = len(nums)
	if length <= 2{
		return nums
	}
	sort.Ints(nums)
	var res []int = make([]int,length)
	var left int = 0
	var right int = length - 1
	for i := 0;i < length;i += 2{
		res[i] = nums[left]
		if (i + 1) < length{
			res[i + 1] = nums[right]
		}
		left++
		right--
	}
	return res
}