package array

import "sort"

func check_minimizeMax(nums []int, l int, p int, target int) bool {
	var cnt int = 0
	var i int = 1
	for i < l {
		if (nums[i] - nums[i-1]) <= target {
			cnt++
			i += 2
		} else {
			i++
		}
	}
	return cnt >= p
}

func MinimizeMax(nums []int, p int) int {
	var l int = len(nums)
	sort.Ints(nums)
	var low int = 0
	var high int = nums[l-1] - nums[0]
	var res int = low
	for low < high {
		mid := (low + high) / 2 //假设存在p个数值对，最大的diff是mid
		ret := check_minimizeMax(nums, l, p, mid)
		if !ret {
			low = mid + 1
			res = mid + 1
		} else {
			high = mid
			res = mid
		}
	}
	return res
}
