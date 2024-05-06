package array

import "sort"

func MinimumAddedInteger(nums1 []int, nums2 []int) int {
	var l1 int = len(nums1)
	var l2 int = len(nums2)
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res int = 1<<31 - 1
	for i := 2; i >= 0; i-- { //要保留的索引
		var idx1 int = 0
		var idx2 int = 0
		var diff_cnt int = 0
		var diff int = nums2[0] - nums1[i]
		for idx1 < l1 && idx2 < l2 {
			if nums1[idx1]+diff == nums2[idx2] {
				diff_cnt++
				idx1++
				idx2++
			} else {
				idx1++
			}
		}
		if diff_cnt == l2 {
			if diff < res {
				res = diff
			}
		}
	}
	return res
}
