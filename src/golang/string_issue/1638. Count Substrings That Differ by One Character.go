package string_issue

func isvalid_countSubstrings(sub string,t string)int{
	var t_len int = len(t)
	var sub_len int = len(sub)
	var res int = 0
	for i := 0;i < t_len - sub_len + 1;i++{
		var diff_cnt int = 0
		for j := 0;j < sub_len;j++{
			if sub[j] != t[i + j]{
				diff_cnt++
			}
			if diff_cnt > 1{
				break
			}
		}
		if diff_cnt == 1{
			res++
		}
	}
	return res
}

func CountSubstrings(s string, t string) int {
	var s_len int = len(s)
	var t_len int = len(t)
	var res int = 0
	for l := 1;l <= s_len;l++{
		if l > t_len{
			break
		}
		for i := 0;i + l <= s_len;i++{
			var sub string = s[i:i + l]
			res += isvalid_countSubstrings(sub,t)
		}
	}
	return res
}