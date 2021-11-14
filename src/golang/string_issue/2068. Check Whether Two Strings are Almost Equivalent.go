package string_issue

func checkAlmostEquivalent(word1 string, word2 string) bool {
	var cnt1 [26]int
	var cnt2 [26]int
	for _,w := range word1{
		cnt1[w - 'a']++
	}
	for _,w := range word2{
		cnt2[w - 'a']++
	}
	for i := 0;i < 26;i++{
		if abs_int(cnt1[i] - cnt2[i]) > 3{
			return false
		}
	}
	return true
}