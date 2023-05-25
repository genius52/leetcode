package string_issue

func makeSmallestPalindrome(s string) string {
	var l int = len(s)
	var res string
	for i := 0; i < l/2; i++ {
		if s[i] == s[l-1-i] {
			res += string(s[i])
		} else {
			if s[i] < s[l-1-i] {
				res += string(s[i])
			} else {
				res += string(s[l-1-i])
			}
		}
	}
	if (l | 1) == l {
		res += string(s[l/2])
	}
	for i := l/2 - 1; i >= 0; i-- {
		res += string(res[i])
	}
	return res
}
