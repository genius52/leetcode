package string_issue

func doesAliceWin(s string) bool {
	var l int = len(s)
	var cnt int = 0
	for i := 0; i < l; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			cnt++
		}
	}
	if cnt == 0 {
		return false
	}
	return true
}
