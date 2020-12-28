package string_issue

func halvesAreAlike(s string) bool {
	var l int = len(s)
	var cnt int = 0
	for i := 0;i < l/2;i++{
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U'{
			cnt++
		}
	}
	for i := l/2;i < l;i++{
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U'{
			cnt--
		}
	}
	return cnt == 0
}