package string_issue


func AppealSum(s string) int64 {
	var last_char_idx [26]int
	for i := 0;i < 26;i++{
		last_char_idx[i] = -1
	}
	var res int64
	var l int = len(s)
	for i := 0;i < l;i++{
		idx := s[i] - 'a'
		var cur int64 = 0
		if last_char_idx[idx] == -1{
			cur = int64(i + 1) * int64(l - i)
		}else{
			cur = int64(i - last_char_idx[idx]) * int64(l - i)
		}
		last_char_idx[idx] = i
		res += cur
	}
	return res
}