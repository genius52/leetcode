package string_issue

func canConstruct(ransomNote string, magazine string) bool {
	var dict [26]int
	for _,c := range magazine{
		dict[c - 'a']++
	}
	for _,c := range ransomNote{
		dict[c - 'a']--
		if dict[c - 'a'] < 0{
			return false
		}
	}
	return true
}