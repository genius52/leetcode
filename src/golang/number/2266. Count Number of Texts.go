package number

func CountTexts(pressedKeys string) int {
	var l int = len(pressedKeys)
	var MOD int = 1e9 + 7
	var dp []int = make([]int,l)//dp[i]: total count on index i
	dp[0] = 1
	for i := 1;i < l;i++{
		dp[i] = dp[i - 1]
		if pressedKeys[i] == pressedKeys[i - 1]{
			if i >= 2{
				dp[i] += dp[i - 2]
				if pressedKeys[i] == pressedKeys[i - 2]{
					if i >= 3{
						dp[i] += dp[i - 3]
						if pressedKeys[i] == '7' || pressedKeys[i] == '9'{
							if i >= 4{
								if pressedKeys[i] == pressedKeys[i - 3]{
									dp[i] += dp[i - 4]
								}
							}else{
								dp[i] += 1
							}
						}
					}else{
						dp[i] += 1
					}
				}
			}else{
				dp[i] += 1
			}
		}
		dp[i] %= MOD
	}
	return dp[l - 1]
}