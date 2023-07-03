package string_issue

func MaximumNumberOfStringPairs(words []string) int {
	var record map[int]int = make(map[int]int)
	for _, word := range words {
		record[int(word[0]-'a')+int(word[1]-'a')*26]++
	}
	var res int = 0
	for _, word := range words {
		target := int(word[1]-'a') + int(word[0]-'a')*26

		if cnt, ok := record[target]; ok {
			if word[0] == word[1] {
				res += cnt - 1
			} else {
				res += cnt
			}
			val2 := int(word[0]-'a') + int(word[1]-'a')*26
			delete(record, val2)
		}
	}
	return res
}
