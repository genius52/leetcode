package string_issue

func countSubstrings3084(s string, c byte) int64 {
	var l int = len(s)
	var cnt int64 = 0
	for i := 0; i < l; i++ {
		if s[i] == c {
			cnt++
		}
	}
	return cnt*(cnt-1)/2 + cnt
}
