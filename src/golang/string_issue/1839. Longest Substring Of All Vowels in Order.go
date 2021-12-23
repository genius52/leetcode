package string_issue

func LongestBeautifulSubstring(word string) int {
	var l int = len(word)
	var find_a bool = false
	var find_e bool = false
	var find_i bool = false
	var find_o bool = false
	var find_u bool = false
	var left int = 0
	var res int = 0
	for i := 0; i < l; i++ {
		if word[i] == 'a' {
			if find_a {
				if find_e {
					left = i
					find_e = false
					find_i = false
					find_o = false
					find_u = false
				}
			} else {
				left = i
				find_a = true
				find_e = false
				find_i = false
				find_o = false
				find_u = false
			}
		} else if word[i] == 'e' {
			if find_a && !find_i {
				find_e = true
			} else {
				find_a = false
				find_e = false
				find_i = false
				find_o = false
				find_u = false
			}
		} else if word[i] == 'i' {
			if find_e && !find_o {
				find_i = true
			} else {
				find_a = false
				find_e = false
				find_i = false
				find_o = false
				find_u = false
			}
		} else if word[i] == 'o' {
			if find_i && !find_u {
				find_o = true
			} else {
				find_a = false
				find_e = false
				find_i = false
				find_o = false
				find_u = false
			}
		} else if word[i] == 'u' {
			if find_a && find_e && find_i && find_o {
				find_u = true
				res = max_int(res, i-left+1)
			} else {
				find_a = false
				find_e = false
				find_i = false
				find_o = false
				find_u = false
			}
		}
	}
	return res
}
