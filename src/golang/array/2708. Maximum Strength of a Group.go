package array

import "sort"

func MaxStrength(nums []int) int64 {
	sort.Ints(nums)
	var l int = len(nums)
	if l == 1 {
		return int64(nums[0])
	}
	if nums[0] == 0 && nums[l-1] == 0 {
		return 0
	}
	var res int64 = 1
	var right int = l - 1
	var find bool = false
	for right >= 0 && nums[right] > 0 {
		res *= int64(nums[right])
		right--
		find = true
	}
	for right >= 0 && nums[right] == 0 {
		right--
	}
	var left int = 0
	if (right | 1) != right {
		right--
	}
	if !find && right < 0 {
		return 0
	}
	for left <= right {
		res *= int64(nums[left])
		left++
	}
	return res
}
