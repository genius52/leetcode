package string_issue

func balancedStringSplit2(s string) int{
	var l int = len(s)
	var res int = 0
	var l_cnt int = 0
	var r_cnt int = 0
	var visit int = 0
	for visit < l{
		if s[visit] == 'L'{
			l_cnt++
		}else{
			r_cnt++
		}
		if l_cnt == r_cnt{
			res++
			l_cnt = 0
			r_cnt = 0
		}
		visit++
	}
	return res
}

func check_balance(s string) bool{
	var L_cnt int = 0
	var R_cnt int = 0
	for i := 0 ;i < len(s);i++{
		if s[i] == 'L'{
			L_cnt++
		}else{
			R_cnt++
		}
	}
	return L_cnt == R_cnt
}

//"RLRRLLRLRL"
func balancedStringSplit(s string) int {
	var cnt int = 0
	begin := 0
	cur := 2
	for ;begin < len(s) && cur <= len(s);{
		sub := s[begin:cur]
		if check_balance(sub) {
			begin = cur
			cur = begin + 1
			cnt++
		}else{
			cur++
		}
	}
	return cnt
}