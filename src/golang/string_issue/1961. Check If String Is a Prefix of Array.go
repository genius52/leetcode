package string_issue

func isPrefixString(s string, words []string) bool {
	var l int = len(words)
	var len_s int = len(s)
	var sub string
	for i := 0;i < l;i++{
		sub += words[i]
		if len(sub) == len_s{
			break
		}
	}
	if s == sub{
		return true
	}
	return false
}