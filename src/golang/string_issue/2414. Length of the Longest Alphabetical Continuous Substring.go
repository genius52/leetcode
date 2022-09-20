package string_issue

func longestContinuousSubstring(s string) int {
	var l int = len(s)
	if l <= 1{
		return l
	}
	var res int = 1
	var cur_len int = 1
	for i := 1;i < l;i++{
		if s[i] - s[i - 1] == 1{
			cur_len++
			if cur_len > res{
				res = cur_len
				if res == 26{
					return res
				}
			}
		}else{
			cur_len = 1
		}
	}
	return res
}