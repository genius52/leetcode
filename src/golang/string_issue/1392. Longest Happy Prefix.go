package string_issue

//Input: s = "ababab"
//Output: "abab"
func longestPrefix(s string) string {
	l := len(s)
	var mod int = 10e9 + 7
	var res string
	var head int = 0
	var tail int = 0
	var times int = 1
	for i := 0;i < l - 1;i++{
		var s1 int = int((s[i] - 'a') * 26)
		var s2 int = int((s[l - 1 - i] - 'a') * 26)
		head = (head * 26 + s1) % mod
		tail = (s2 * times + tail) % mod
		times = times * 26 % mod
		if head == tail{
			res = s[:i+1]
		}
	}
	return res
}