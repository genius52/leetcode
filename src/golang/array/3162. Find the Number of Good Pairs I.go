package array

func numberOfPairs3162(nums1 []int, nums2 []int, k int) int {
	var num_cnt1 map[int]int = make(map[int]int)
	var num_cnt2 map[int]int = make(map[int]int)
	var res int = 0
	for _, n := range nums1 {
		num_cnt1[n]++
	}
	for _, n := range nums2 {
		num_cnt2[n]++
	}
	for n1, cnt1 := range num_cnt1 {
		for n2, cnt2 := range num_cnt2 {
			if n1 < n2*k {
				continue
			}
			if n1%(n2*k) == 0 {
				res += cnt1 * cnt2
			}
		}
	}
	return res
}
