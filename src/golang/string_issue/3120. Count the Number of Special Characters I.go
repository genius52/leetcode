package string_issue

func numberOfSpecialChars(word string) int {
	var res int = 0
	var lower [26]bool
	var upper [26]bool
	for _, w := range word {
		if w >= 'a' && w <= 'z' && !lower[w-'a'] {
			lower[w-'a'] = true
			if upper[w-'a'] {
				res++
			}
		} else if w >= 'A' && w <= 'Z' && !upper[w-'A'] {
			upper[w-'A'] = true
			if lower[w-'A'] {
				res++
			}
		}
		if res == 26 {
			return res
		}
	}
	return res
}
