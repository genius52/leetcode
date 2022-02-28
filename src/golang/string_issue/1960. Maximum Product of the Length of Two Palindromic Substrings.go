package string_issue

func maxProduct(s string) int64 {
	var l int = len(s)
	var dp []int = make([]int,l)
	for i := 0;i < l;i++{
		var left int = i - 1
		var right int = i + 1
		var max_len int = 1
		for left >= 0 && right < l && s[left] == s[right]{
			max_len += 2
		}
		dp[i] = max_len
	}
	return 0
}