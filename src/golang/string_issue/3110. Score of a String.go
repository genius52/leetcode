package string_issue

func scoreOfString(s string) int {
	var l int = len(s)
	var res int = 0
	for i := 1; i < l; i++ {
		diff := int(s[i]) - int(s[i-1])
		if diff < 0 {
			diff = -diff
		}
		res += diff
	}
	return res
}
