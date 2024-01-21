package number

import "sort"

func MaximumSetSize(nums1 []int, nums2 []int) int {
	var l int = len(nums1)
	var record1 map[int]int = make(map[int]int)
	var record2 map[int]int = make(map[int]int)
	for i := 0; i < l; i++ {
		record1[nums1[i]]++
		record2[nums2[i]]++
	}
	//var l1 int = len(record1)
	//var l2 int = len(record2)
	var single1 []int
	var single2 []int
	var share_cnt int = 0
	for n, cnt1 := range record1 {
		if _, ok := record2[n]; !ok {
			single1 = append(single1, cnt1)
		} else {
			share_cnt++
		}
	}
	for n, cnt2 := range record2 {
		if _, ok := record1[n]; !ok {
			single2 = append(single2, cnt2)
		}
	}
	sort.Ints(single1)
	sort.Ints(single2)
	var res int = 0
	if len(single1) >= l/2 {
		res += l / 2
	} else {
		res += len(single1)
	}
	if len(single2) > l/2 {
		res += l / 2
	} else {
		res += len(single2)
	}
	if res < l {
		res += share_cnt
	}
	if res > l {
		return l
	}
	return res
}
