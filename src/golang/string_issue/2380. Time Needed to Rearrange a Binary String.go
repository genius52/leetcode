package string_issue

func SecondsToRemoveOccurrences(s string) int {
	var l int = len(s)
	var dp []int = make([]int,l + 1)//dp[i]:
	var zero_cnt int = 0
	for i := 0;i < l;i++{
		if s[i] == '0'{
			dp[i + 1] = dp[i]
			zero_cnt++
		}else{
			if zero_cnt > 0{
				dp[i + 1] = max_int(1 + dp[i],zero_cnt)
			}
		}
	}
	return dp[l]
}