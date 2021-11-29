package array

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	var l int = len(nums)
	var res []int
	for i := 0;i < l;i++{
		if nums[i] == target{
			res = append(res,i)
		}
	}
	return res
}