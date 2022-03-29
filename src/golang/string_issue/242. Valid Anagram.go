package string_issue

func isAnagram(s string, t string) bool {
	var l1 int = len(s)
	var l2 int = len(t)
	if l1 != l2{
		return false
	}
	var cnt [26]int
	for _,c := range s{
		cnt[c - 'a']++
	}
	for _,c := range t{
		cnt[c - 'a']--
	}
	for i := 0;i < 26;i++{
		if cnt[i] != 0{
			return false
		}
	}
	return true
}