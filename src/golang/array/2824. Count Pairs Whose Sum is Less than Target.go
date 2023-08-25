package array

import "sort"

// binary search
func CountPairs2824(nums []int, target int) int {
	sort.Ints(nums)
	var res int = 0
	var l int = len(nums)
	for i := 0; i < l; i++ {
		find_idx := sort.Search(l, func(j int) bool {
			return nums[i]+nums[j] >= target
		})
		if (find_idx - 1) > i {
			res += find_idx - 1 - i
		} else {
			break
		}
	}
	return res
}

// two pointer
func countPairs2824(nums []int, target int) int {
	sort.Ints(nums)
	var res int = 0
	var l int = len(nums)
	var left int = 0
	var right int = l - 1
	for left < right {
		if nums[left]+nums[right] < target {
			res += right - left
			left++
		} else {
			right--
		}
	}
	return res
}
