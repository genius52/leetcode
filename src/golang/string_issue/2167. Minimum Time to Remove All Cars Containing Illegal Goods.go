package string_issue

//输入：s = "1100101"
//输出：5
func MinimumTime(s string) int {
	var l int = len(s)
	if l == 1{
		if s[0] == '0'{
			return 0
		}else{
			return 1
		}
	}
	var dp_left []int = make([]int,l)//在i位置最少的耗时
	var dp_right []int = make([]int,l)
	for i := 0;i < l;i++{
		if s[i] == '1'{
			if i == 0{
				dp_left[i] = 1
			}else{
				dp_left[i] = min_int(i + 1,dp_left[i - 1] + 2)
			}
		}else{
			if i != 0 {
				dp_left[i] = dp_left[i-1]
			}
		}
	}
	for i := l - 1;i >= 0;i--{
		if s[i] == '1'{
			if i == l - 1{
				dp_right[i] = 1
			}else{
				dp_right[i] = min_int(l - i,dp_right[i + 1] + 2)
			}
		}else{
			if i != l - 1 {
				dp_right[i] = dp_right[i + 1]
			}
		}
	}
	var res int = 1 << 31
	for i := 0;i < l - 1;i++{
		res = min_int(res,dp_left[i] + dp_right[i + 1])
	}
	return res
}