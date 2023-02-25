package number

func MinOperations2571(n int) int {
	var s string
	for n > 0 {
		if (n | 1) == n {
			s += "1"
		} else {
			s += "0"
		}
		n >>= 1
	}
	var l int = len(s)
	var res int = 0
	var one_cnt int = 0
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			if one_cnt == 1 {
				res += 1
				one_cnt = 0
			} else if one_cnt > 1 {
				res += 1
				one_cnt = 1
			}
		} else if s[i] == '1' {
			one_cnt++
		}
	}
	if one_cnt > 1 {
		res += 2
	} else {
		res += 1
	}
	return res
}
