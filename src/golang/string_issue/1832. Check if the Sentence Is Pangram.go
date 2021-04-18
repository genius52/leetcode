package string_issue

func checkIfPangram(sentence string) bool {
	var l int = len(sentence)
	if l < 26{
		return false
	}
	var record [26]bool
	for i := 0;i < l;i++{
		record[sentence[i] - 'a'] = true
	}
	for i := 0;i < 26;i++{
		if !record[i]{
			return false
		}
	}
	return true
}