package string_issue

//balloon
func maxNumberOfBalloons(text string) int {
	var cnt [26]int
	for _,c := range text{
		cnt[c - 'a']++
	}
	var res int = 2147483647
	if cnt['b' - 'a'] < res{
		res = cnt['b' - 'a']
	}
	if cnt['a' - 'a'] < res{
		res = cnt['a' - 'a']
	}
	if cnt['l' - 'a']/2 < res{
		res = cnt['l' - 'a']/2
	}
	if cnt['o' - 'a']/2 < res{
		res = cnt['o' - 'a']/2
	}
	if cnt['n' - 'a'] < res{
		res = cnt['n' - 'a']
	}
	return res
}