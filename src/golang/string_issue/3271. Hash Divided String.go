package string_issue

func stringHash(s string, k int) string {
	var l int = len(s)
	var cnt int = l / k
	var res string
	for i := 0; i < cnt; i++ {
		var sum int = 0
		for j := 0; j < k; j++ {
			sum += int(s[j+i*k] - 'a')
		}
		sum %= 26
		res += string('a' + sum)
	}
	return res
}
