package array

import "sort"

func minimumArrayLength(nums []int) int {
	var l int = len(nums)
	sort.Ints(nums)
	var find bool = false
	for i := 1; i < l; i++ {
		if nums[i]%nums[0] != 0 {
			find = true //用更小的数将目前最小的数字剔除
			break
		}
	}
	if find {
		return 1
	}
	var cnt int = 1
	for i := 1; i < l; i++ {
		if nums[i] == nums[0] {
			cnt++
		} else {
			break
		}
	}
	return cnt/2 + cnt%2
}
