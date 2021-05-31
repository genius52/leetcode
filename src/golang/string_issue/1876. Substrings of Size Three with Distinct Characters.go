package string_issue

func countGoodSubstrings(s string) int {
	var l int = len(s)
	if l < 3{
		return 0
	}
	var res int = 0
	var visit int = 1
	for visit <= l - 2{
		if s[visit - 1] != s[visit] && s[visit - 1] != s[visit + 1] && s[visit] != s[visit + 1]{
			res++
		}
		visit++
	}
	return res
}