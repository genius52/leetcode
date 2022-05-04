package array

import "sort"

func findUnsortedSubarray(nums []int) int {
	var l int = len(nums)
	var sorted []int = make([]int,l)
	copy(sorted,nums)
	sort.Ints(sorted)
	var left int = 0
	var right int = l - 1
	for left < l && nums[left] == sorted[left]{
		left++
	}
	if left == l{
		return 0
	}
	for right >= 0 && nums[right] == sorted[right]{
		right--
	}
	return right - left + 1
}

func findUnsortedSubarray2(nums []int) int{
	var l int = len(nums)
	var pre_max_val int = -2147483648
	var right int = -1
	for i := 0;i < l;i++{
		if nums[i] < pre_max_val{
			right = i
		}else{
			pre_max_val = nums[i]
		}
	}
	var pre_min_val int = 2147483647
	var left int = -1
	for i := l - 1;i >= 0;i--{
		if nums[i] > pre_min_val{
			left = i
		}else{
			pre_min_val = nums[i]
		}
	}
	if right == -1{
		return 0
	}
	return right - left + 1
}