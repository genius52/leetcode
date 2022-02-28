package string_issue

func checkAlmostEquivalent(word1 string, word2 string) bool {
	var cnt [26]int
	for _,w := range word1{
		cnt[w - 'a']++
	}
	for _,w := range word2{
		cnt[w - 'a']--
	}
	for i := 0;i < 26;i++{
		if abs_int(cnt[i]) > 3{
			return false
		}
	}
	return true
}