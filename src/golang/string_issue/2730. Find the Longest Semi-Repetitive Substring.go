package string_issue

func longestSemiRepetitiveSubstring(s string) int {
	var l int = len(s)
	var left int = 0
	var right int = 1
	var res int = 1
	var find_dup bool = false
	var first_idx int = -1
	for left < l && right < l {
		if s[right] == s[right-1] {
			if find_dup { //???
				cur_len := right - left
				if cur_len > res {
					res = cur_len
				}
				left = first_idx
				first_idx = right
				right++
			} else {
				find_dup = true
				first_idx = right
				cur_len := right - left + 1
				if cur_len > res {
					res = cur_len
				}
				right++
			}
		} else {
			cur_len := right - left + 1
			if cur_len > res {
				res = cur_len
			}
			right++
		}
	}
	return res
}
