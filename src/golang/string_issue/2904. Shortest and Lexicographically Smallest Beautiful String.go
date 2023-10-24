package string_issue

import "strings"

func ShortestBeautifulSubstring(s string, k int) string {
	var l int = len(s)
	var left int = 0
	var right int = 0
	var one_cnt int = 0
	var min_len int = l + 1
	var res string
	for left < l {
		for right < l && one_cnt < k {
			if s[right] == '1' {
				one_cnt++
			}
			right++
		}
		if one_cnt == k {
			cur_len := right - left
			if cur_len < min_len {
				min_len = cur_len
				res = s[left:right]
			} else if cur_len == min_len {
				if strings.Compare(s[left:right], res) < 0 {
					res = s[left:right]
				}
			}
		}
		if s[left] == '1' {
			one_cnt--
		}
		left++
	}
	return res
}
