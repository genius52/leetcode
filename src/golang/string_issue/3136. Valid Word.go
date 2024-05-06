package string_issue

func isValid3136(word string) bool {
	var l int = len(word)
	if l < 3 {
		return false
	}
	//var find_digit bool = false
	//var find_lower bool = false
	//var find_upper bool = false
	var find_vowel bool = false
	var find_consonant bool = false
	for _, c := range word {
		if c >= '0' && c <= '9' {
			//find_digit = true
		} else if c >= 'a' && c <= 'z' {
			//find_lower = true
			if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
				find_vowel = true
			} else {
				find_consonant = true
			}
		} else if c >= 'A' && c <= 'Z' {
			//find_upper = true
			if c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
				find_vowel = true
			} else {
				find_consonant = true
			}
		} else {
			return false
		}
	}
	return find_vowel && find_consonant
}
