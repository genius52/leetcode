package string_issue

func checkStrings(s1 string, s2 string) bool {
	var l int = len(s1)
	var even_group [26]int
	for i := 0; i < l; i += 2 {
		even_group[s1[i]-'a']++
		even_group[s1[i]-'a']++
		even_group[s2[i]-'a']--
		even_group[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if even_group[i] != 0 {
			return false
		}
	}
	var odd_group [26]int
	for i := 1; i < l; i += 2 {
		odd_group[s1[i]-'a']++
		odd_group[s1[i]-'a']++
		odd_group[s2[i]-'a']--
		odd_group[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if odd_group[i] != 0 {
			return false
		}
	}
	return true
}
