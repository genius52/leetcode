package number

import "math"

func NumberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	var num_cnt1 map[int]int = make(map[int]int)
	var num_cnt2 map[int]int = make(map[int]int)
	for _, n := range nums1 {
		num_cnt1[n]++
	}
	for _, n := range nums2 {
		num_cnt2[n*k]++
	}
	var res int64 = 0
	for n1, cnt1 := range num_cnt1 {
		for i := 1; i <= int(math.Sqrt(float64(n1))); i++ {
			if n1%i == 0 {
				val := n1 / i
				if cnt2, ok := num_cnt2[val]; ok {
					res += int64(cnt1) * int64(cnt2)
				}
				if i != val {
					if cnt2, ok := num_cnt2[i]; ok {
						res += int64(cnt1) * int64(cnt2)
					}
				}
			}
		}
	}
	return res
}
