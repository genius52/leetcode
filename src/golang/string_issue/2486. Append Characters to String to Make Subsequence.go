package string_issue

func appendCharacters(s string, t string) int {
	var l1 int = len(s)
	var l2 int = len(t)
	var idx1 int = 0
	var idx2 int = 0
	for idx1 < l1 && idx2 < l2{
		for idx1 < l1 && s[idx1] != t[idx2]{
			idx1++
		}
		if idx1 == l1{
			break
		}
		idx1++
		idx2++
	}
	return l2 - idx2
}