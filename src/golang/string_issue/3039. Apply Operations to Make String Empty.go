package string_issue

func lastNonEmptyString(s string) string {
	var l int = len(s)
	var cnt [26]int
	for i := 0; i < l; i++ {
		cnt[s[i]-'a']++
	}
	var max_cnt int = 0
	for i := 0; i < 26; i++ {
		if cnt[i] > max_cnt {
			max_cnt = cnt[i]
		}
	}
	var res string
	var cnt2 [26]int
	for i := 0; i < l; i++ {
		c := s[i] - 'a'
		if cnt[c] == max_cnt {
			cnt2[c]++
			if cnt2[c] == max_cnt {
				res += string(s[i])
			}
		}
	}
	return res
}
