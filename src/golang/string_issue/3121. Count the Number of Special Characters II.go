package string_issue

func numberOfSpecialChars2(word string) int {
	var lower [26]bool
	var upper [26]bool
	var failure [26]bool
	for _, w := range word {
		if w >= 'a' && w <= 'z' && !failure[w-'a'] {
			lower[w-'a'] = true
			if upper[w-'a'] {
				failure[w-'a'] = true
			}
		} else if w >= 'A' && w <= 'Z' && !failure[w-'A'] {
			upper[w-'A'] = true
			if !lower[w-'A'] {
				failure[w-'A'] = true
			}
		}
	}
	var res int = 0
	for i := 0; i < 26; i++ {
		if lower[i] && upper[i] && !failure[i] {
			res++
		}
	}
	return res
}
