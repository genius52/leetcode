package string_issue

func findTheLongestBalancedSubstring(s string) int {
	var l int = len(s)
	var res int = 0
	var zero_len int = 0
	var one_len int = 0
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			if i > 0 && s[i-1] == '1' {
				zero_len = 0
			}
			zero_len++
			one_len = 0
		} else if s[i] == '1' {
			one_len++
			cur := min_int(zero_len, one_len) * 2
			res = max_int(res, cur)
		}
	}
	return res
}
