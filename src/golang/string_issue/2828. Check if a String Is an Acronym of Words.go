package string_issue

func isAcronym(words []string, s string) bool {
	var l1 int = len(words)
	var l2 int = len(s)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		if s[i] != words[i][0] {
			return false
		}
	}
	return true
}
