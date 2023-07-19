package array

import "sort"

func maximumBeauty2779(nums []int, k int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var res int = 1
	var left int = 0
	var right int = 0
	for left < l && right < l {
		for right < l && (nums[right]-nums[left]) <= k*2 {
			right++
		}
		cur_len := right - left
		if cur_len > res {
			res = cur_len
		}
		left++
	}
	return res
}
