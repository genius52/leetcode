package string_issue

func checkZeroOnes(s string) bool {
	var l int = len(s)
	var max_zero_len int = 0
	var max_one_len int = 0
	var left int = 0
	for left < l{
		var right int = left
		for right < l && s[left] == s[right]{
			right++
		}
		cur := right - left
		if s[left] == '0'{
			if cur > max_zero_len{
				max_zero_len = cur
			}
		}else{
			if cur > max_one_len{
				max_one_len = cur
			}
		}
		left = right
	}
	return max_one_len > max_zero_len
}