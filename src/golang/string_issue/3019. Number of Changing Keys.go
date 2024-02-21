package string_issue

func CountKeyChanges(s string) int {
	var res int = 0
	var gap int = abs_int('a' - 'A')
	for i := 1; i < len(s); i++ {
		diff := int(s[i]) - int(s[i-1])
		if diff == 0 || abs_int(diff) == gap {
			continue
		}
		res++
	}
	return res
}
