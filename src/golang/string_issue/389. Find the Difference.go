package string_issue

func findTheDifference(s string, t string) byte {
	var cnt [26]int
	for _,c := range t{
		cnt[c - 'a']++
	}
	for _,c := range s{
		cnt[c - 'a']--
	}
	var res byte
	for i := 0;i < 26;i++{
		if cnt[i] == 1{
			res = byte(i) + byte('a')
			break
		}
	}
	return res
}