package string_issue

func canConstruct1400(s string, k int) bool {
	var l int = len(s)
	if k > l{
		return false
	}
	var record [26]int
	for _,c := range s{
		record[c - 'a']++
	}
	for i := 0;i < 26;i++{
		if (record[i] | 1) == record[i]{
			k--
		}
	}
	return k >= 0
}