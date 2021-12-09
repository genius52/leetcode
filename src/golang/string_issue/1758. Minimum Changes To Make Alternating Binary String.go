package string_issue

func min_int(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minOperations(s string) int {
	var l int = len(s)
	var start_zero_opt_cnt int = 0 //010101
	var start_one_opt_cnt int = 0  //101010
	for i := 0; i < l; i++ {
		if i|1 == i { //odd index
			if s[i] == '0' {
				start_zero_opt_cnt++
			} else {
				start_one_opt_cnt++
			}
		} else { //even index
			if s[i] == '0' {
				start_one_opt_cnt++
			} else {
				start_zero_opt_cnt++
			}
		}
	}
	if start_zero_opt_cnt < start_one_opt_cnt {
		return start_zero_opt_cnt
	}
	return start_one_opt_cnt
}

func MinOperations(s string) int {
	var l int = len(s)
	if l <= 1 {
		return 0
	}
	var zero_one int = 0
	var one_zero int = 0
	for i := 0; i < l; i++ {
		if (i | 1) == i {
			if s[i] != '1' {
				zero_one++
			} else {
				one_zero++
			}
		} else {
			if s[i] != '0' {
				zero_one++
			} else {
				one_zero++
			}
		}
	}
	return min_int(zero_one, one_zero)
}
