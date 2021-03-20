package string_issue

func firstUniqChar(s string) int {
	var record [26]int
	var l int = len(s)
	for i := 0;i < l;i++{
		record[s[i] - 'a']++
	}
	var res int = -1
	for i := 0;i < l;i++{
		if record[s[i] - 'a'] == 1{
			res = i
			break
		}
	}
	return res
}