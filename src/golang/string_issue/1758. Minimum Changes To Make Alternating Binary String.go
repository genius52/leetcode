package string_issue

func min_int(a,b int)int{
	if a < b {
		return a
	}else{
		return b
	}
}

func MinOperations(s string) int {
	var l int = len(s)
	if l <= 1{
		return 0
	}
	var zero_one int = 0
	var one_zero int = 0
	for i := 0;i < l;i++{
		if (i | 1) == i{
			if s[i] != '1'{
				zero_one++
			}else{
				one_zero++
			}
		}else{
			if s[i] != '0'{
				zero_one++
			}else{
				one_zero++
			}
		}
	}
	return min_int(zero_one,one_zero)
}