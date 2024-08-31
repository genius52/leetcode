package number

import "strconv"

func countPairs3265(nums []int) int {
	var l int = len(nums)
	var res int = 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] == nums[j] {
				res++
			} else {
				s1 := strconv.Itoa(nums[i])
				s2 := strconv.Itoa(nums[j])
				l1 := len(s1)
				l2 := len(s2)
				//if abs_int(l1-l2) >= 2 {
				//	continue
				//}
				//if l1 == l2+1 {
				//	s2 = "0" + s2
				//} else if l1+1 == l2 {
				//	s1 = "0" + s1
				//}
				if l1 > l2 {
					diff := l1 - l2
					for diff > 0 {
						s2 = "0" + s2
						diff--
					}
				}
				if l1 < l2 {
					diff := l2 - l1
					for diff > 0 {
						s1 = "0" + s2
						diff--
					}
				}
				var cur_len int = max(l1, l2)
				var diff_idx []int
				for k := 0; k < cur_len; k++ {
					if s1[k] != s2[k] {
						diff_idx = append(diff_idx, k)
					}
				}
				if len(diff_idx) == 2 {
					if s1[diff_idx[0]] == s2[diff_idx[1]] && s1[diff_idx[1]] == s2[diff_idx[0]] {
						res++
					}
				}
			}
		}
	}
	return res
}
