package array

import "sort"

func cal_bit(n int) int {
	var cnt int = 0
	for n > 0 {
		if n|1 == n {
			cnt++
		}
		n >>= 1
	}
	return cnt
}

func CanSortArray(nums []int) bool {
	var l int = len(nums)
	var left int = 0
	var right int = 1
	var increase bool = true
	for left < l {
		for right < l && cal_bit(nums[left]) == cal_bit(nums[right]) {
			right++
		}
		sort.Ints(nums[left:right])
		if left > 0 {
			if nums[left] < nums[left-1] {
				increase = false
				break
			}
		}
		left = right
		right = left + 1
	}
	return increase
}
