package string_issue

func CanMakeSubsequence(str1 string, str2 string) bool {
	var l1 int = len(str1)
	var l2 int = len(str2)
	if l1 < l2 {
		return false
	}
	var idx1 int = 0
	var idx2 int = 0
	for idx1 < l1 && idx2 < l2 {
		if str1[idx1] == str2[idx2] {
			idx1++
			idx2++
		} else {
			c1 := (str1[idx1] - 'a' + 1) % 26
			c2 := str2[idx2] - 'a'
			if c1 == c2 {
				idx1++
				idx2++
			} else {
				idx1++
			}
		}
	}
	return idx2 == l2
}
