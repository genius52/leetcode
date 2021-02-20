package string_issue

func NumDecodings(s string) int {
	var l int = len(s)
	var dp []int = make([]int,l)
	if s[0] != '0'{
		dp[0] = 1
	}else{
		return 0
	}
	for i := 1;i < l;i++{
		if s[i] == '0'{
			if s[i - 1] == '2' || s[i - 1] == '1'{
				if i > 1{
					dp[i] += dp[i - 2]
				}else{
					dp[i] = 1
				}
			}else{
				return 0
			}
		}else {
			dp[i] = dp[i - 1]
			if s[i - 1] == '1'{
				if i > 1{
					dp[i] += dp[i - 2]
				}else{
					dp[i]++
				}
			}else if s[i - 1] == '2'{
				if s[i] >= '0' && s[i] <= '6'{
					if i > 1{
						dp[i] += dp[i - 2]
					}else{
						dp[i]++
					}
				}
			}
		}
	}
	return dp[l - 1]
}