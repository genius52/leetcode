package string_issue

func countKConstraintSubstrings(s string, k int) int {
	var l int = len(s)
	var res int = 0
	for left := 0; left < l; left++ {
		var zero_cnt int = 0
		var one_cnt int = 0
		for right := left; right < l; right++ {
			if s[right] == '0' {
				zero_cnt++
			} else {
				one_cnt++
			}
			if zero_cnt > k && one_cnt > k {
				break
			}
			res++
		}
	}
	return res
}

func CountKConstraintSubstrings(s string, k int) int {
	var l int = len(s)
	var res int = 0
	var left int = 0
	var right int = 0
	var zero_cnt int = 0
	var one_cnt int = 0
	for left < l && right < l {
		if s[right] == '0' {
			zero_cnt++
		} else {
			one_cnt++
		}
		for zero_cnt > k && one_cnt > k {
			if s[left] == '0' {
				zero_cnt--
			} else {
				one_cnt--
			}
			left++
		}
		res += right - left + 1
		right++
	}
	return res
}
