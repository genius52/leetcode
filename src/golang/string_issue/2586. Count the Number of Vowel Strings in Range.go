package string_issue

func vowelStrings2586(words []string, left int, right int) int {
	var res int = 0
	for i := left; i <= right; i++ {
		var l int = len(words[i])
		if (words[i][0] == 'a' || words[i][0] == 'e' || words[i][0] == 'i' || words[i][0] == 'o' || words[i][0] == 'u') &&
			(words[i][l-1] == 'a' || words[i][l-1] == 'e' || words[i][l-1] == 'i' || words[i][l-1] == 'o' || words[i][l-1] == 'u') {
			res++
		}
	}
	return res
}
