package string_issue

func MinCharacters(a string, b string) int {
	var a_cnt [26]int
	var b_cnt [26]int
	var same_cnt_a int = 0
	var same_cnt_b int = 0
	var len_a int = len(a)
	var len_b int = len(b)
	for _, c := range a {
		a_cnt[c-'a']++
	}
	for _, c := range b {
		b_cnt[c-'a']++
	}
	var prefix_a [26]int
	var prefix_b [26]int
	var res int = len_a + len_b
	for i := 0; i < 26; i++ {
		if i > 0 {
			res = min_int(res, len_a-prefix_a[i-1]+prefix_b[i-1]) //a变成小字符，b变成大字符
			res = min_int(res, len_b-prefix_b[i-1]+prefix_a[i-1]) //b变成小字符，a变成大字符
		}
		if i == 0 {
			prefix_a[i] = a_cnt[i]
			prefix_b[i] = b_cnt[i]
		} else {
			prefix_a[i] = prefix_a[i-1] + a_cnt[i]
			prefix_b[i] = prefix_b[i-1] + b_cnt[i]
		}
		if a_cnt[i] > same_cnt_a {
			same_cnt_a = a_cnt[i]
		}
		if b_cnt[i] > same_cnt_b {
			same_cnt_b = b_cnt[i]
		}

	}
	res = min_int(res, len_a-same_cnt_a+len_b-same_cnt_b)
	return res
}
