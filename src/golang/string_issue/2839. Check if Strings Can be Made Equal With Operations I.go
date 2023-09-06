package string_issue

func canBeEqual(s1 string, s2 string) bool {
	var group1 [26]int
	group1[s1[0]-'a']++
	group1[s1[2]-'a']++
	group1[s2[0]-'a']--
	group1[s2[2]-'a']--
	if group1[s1[0]-'a'] != 0 || group1[s1[2]-'a'] != 0 {
		return false
	}
	var group2 [26]int
	group2[s1[1]-'a']++
	group2[s1[3]-'a']++
	group2[s2[1]-'a']--
	group2[s2[3]-'a']--
	if group2[s1[1]-'a'] != 0 || group2[s1[3]-'a'] != 0 {
		return false
	}
	return true
}
