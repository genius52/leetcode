package string_issue

func minimumLength(s string) int {
	var char_cnt [26]int
	for _, c := range s {
		char_cnt[c-'a']++
	}
	var res int = 0
	for i := 0; i < 26; i++ {
		if char_cnt[i] >= 3 {
			if (char_cnt[i]-3)%2 == 1 {
				res += 2
			} else {
				res += 1
			}
		} else {
			res += char_cnt[i]
		}
	}
	return res
}
