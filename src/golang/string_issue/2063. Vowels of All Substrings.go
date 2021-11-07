package string_issue

func countVowels(word string) int64 {
	var l int = len(word)
	var res int64 = 0
	for i := 0;i < l;i++{
		if word[i] == 'a' || word[i] == 'a' || word[i] == 'a' || word[i] == 'a' || word[i] == 'a'{
			res += int64((i + 1) * (l - i))
		}
	}
	return res
}