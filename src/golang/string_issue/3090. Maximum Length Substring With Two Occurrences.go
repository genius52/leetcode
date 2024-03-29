package string_issue

func MaximumLengthSubstring(s string) int {
	var l int = len(s)
	var cnt [26]int
	var left int = 0
	var right int = 0
	var res int = 0
	for right < l {
		for right < l {
			cnt[s[right]-'a']++
			if cnt[s[right]-'a'] == 3 {
				right++
				break
			}
			cur_len := right - left + 1
			if cur_len > res {
				res = cur_len
			}
			right++
		}
		for left < right && cnt[s[left]-'a'] != 3 {
			cnt[s[left]-'a']--
			left++
		}
		if left < l {
			cnt[s[left]-'a']--
			left++
		}
	}
	return res
}
