package string_issue

func firstPalindrome(words []string) string {
	for _, word := range words {
		var l int = len(word)
		var same bool = true
		for i := 0; i < l/2; i++ {
			if word[i] != word[l-1-i] {
				same = false
				break
			}
		}
		if same {
			return word
		}
	}
	return ""
}
