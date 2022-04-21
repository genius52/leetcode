package string_issue

//DP bottom to top
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

//DP from top to bottom
func dfs_numDecodings(s string,l int,idx int,memo []int)int{
	if idx == l{
		return 1
	}
	if s[idx] == '0'{
		return 0
	}
	if memo[idx] != -1{
		return memo[idx]
	}
	if s[idx] == '2'{
		if idx + 1 < l{
			if s[idx + 1] >= '0' && s[idx + 1] <= '6'{
				memo[idx] = dfs_numDecodings(s,l,idx + 1,memo) + dfs_numDecodings(s,l,idx + 2,memo)
			}else{
				memo[idx] = dfs_numDecodings(s,l,idx + 1,memo)
			}
		}else{
			memo[idx] = dfs_numDecodings(s,l,idx + 1,memo)
		}
	}else if s[idx] == '1'{
		if idx + 1 < l{
			memo[idx] = dfs_numDecodings(s,l,idx + 1,memo) + dfs_numDecodings(s,l,idx + 2,memo)
		}else{
			memo[idx] = dfs_numDecodings(s,l,idx + 1,memo)
		}
	}else{
		memo[idx] = dfs_numDecodings(s,l,idx + 1,memo)
	}
	return memo[idx]
}

func numDecodings(s string) int{
	var l int = len(s)
	if s[0] == '0'{
		return 0
	}
	var memo []int = make([]int,l)
	for i := 0;i < l;i++{
		memo[i] = -1
	}
	return dfs_numDecodings(s,l,0,memo)
}