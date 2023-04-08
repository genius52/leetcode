package number

import "sort"

func minNumber(nums1 []int, nums2 []int) int {
	var cnt [10]int
	for _, n := range nums1 {
		cnt[n]++
	}
	for _, n := range nums2 {
		cnt[n]++
	}
	for i := 0; i <= 9; i++ {
		if cnt[i] > 1 {
			return i
		}
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	n1 := nums1[0]*10 + nums2[0]
	n2 := nums2[0]*10 + nums1[0]
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
