package string_issue

func repeatedCharacter(s string) byte {
	var find [26]bool
	for _,c := range s{
		if find[c - 'a']{
			return byte(c)
		}
		find[c - 'a'] = true
	}
	return 'a'
}