package string_issue

func minimizedStringLength(s string) int {
	var l int = len(s)
	var cnt [26]int
	for i := 0; i < l; i++ {
		cnt[s[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if cnt[i] > 1 {
			l = l - cnt[i] + 1
		}
	}
	return l
}
