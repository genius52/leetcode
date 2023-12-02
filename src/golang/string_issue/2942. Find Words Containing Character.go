package string_issue

func findWordsContaining(words []string, x byte) []int {
	var res []int
	for idx, word := range words {
		for i := 0; i < len(word); i++ {
			if word[i] == x {
				res = append(res, idx)
				break
			}
		}
	}
	return res
}
