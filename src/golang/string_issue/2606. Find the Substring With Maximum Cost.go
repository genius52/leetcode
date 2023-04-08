package string_issue

func maximumCostSubstring(s string, chars string, vals []int) int {
	var l1 int = len(s)
	var l2 int = len(chars)
	var char_val map[uint8]int = make(map[uint8]int)
	for i := 0; i < l2; i++ {
		char_val[chars[i]] = vals[i]
	}
	var res int = 0
	var pre int = 0
	for i := 0; i < l1; i++ {
		var c int = 0
		if val, ok := char_val[s[i]]; ok {
			c = val
		} else {
			c = int(s[i] - 'a' + 1)
		}
		pre = max_int(pre+c, c)
		res = max_int(pre, res)
	}
	return res
}
