package string_issue

func CheckPartitioning(s string) bool {
	var l int = len(s)
	var dp [][]bool = make([][]bool, l) //dp[i][j] = true,s[i:j + 1] is palin..
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}
	for i := 0; i < l; i++ {
		for j := i; j >= 0; j-- {
			if s[i] != s[j] {
				dp[j][i] = false
			} else {
				if i-j <= 1 {
					dp[j][i] = true
				} else {
					dp[j][i] = dp[j+1][i-1]
				}
			}
		}
	}
	for i := 0; i < l-1; i++ {
		if !dp[0][i] {
			continue
		}
		for j := i + 1; j < l-1; j++ {
			if dp[0][i] && dp[i+1][j] && dp[j+1][l-1] {
				return true
			}
		}
	}
	return false
}
