package string_issue

func checkString(s string) bool {
	var l int = len(s)
	var find_b bool = false
	for i := 0; i < l; i++ {
		if s[i] == 'a' && find_b {
			return false
		}
		if s[i] == 'b' {
			find_b = true
		}
	}
	return true
}
