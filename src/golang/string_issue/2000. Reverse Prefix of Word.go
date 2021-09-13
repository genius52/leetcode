package string_issue

func reversePrefix(word string, ch byte) string {
	var idx int = -1
	var l int = len(word)
	for i := 0;i < l;i++{
		if word[i] == ch{
			idx = i
			break
		}
	}
	if idx == -1{
		return word
	}
	var res string
	for i := idx;i >= 0;i--{
		res += string(word[i])
	}
	res += word[idx + 1:]
	return res
}