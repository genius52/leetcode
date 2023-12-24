package array

import "sort"

func DivideArray2966(nums []int, k int) [][]int {
	var res [][]int
	var l int = len(nums)
	if l%3 != 0 {
		return res
	}
	sort.Ints(nums)
	for i := 2; i < l; i += 3 {
		if nums[i]-nums[i-2] > k {
			return [][]int{}
		}
	}
	for i := 0; i < l; i += 3 {
		res = append(res, nums[i:i+3])
	}
	return res
}
