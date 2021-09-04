package number

func dfs_dieSimulator(n int,last_num int,num_cnt int,rollMax []int,memo *[][6][16]int)int{
	if n == 0{
		return 1
	}
	if last_num >= 0{
		if (*memo)[n][last_num][num_cnt] > 0{
			return (*memo)[n][last_num][num_cnt]
		}
	}
	var res int = 0
	for i := 0;i < 6;i++{
		if i == last_num && num_cnt == rollMax[i]{
			continue
		}
		if i == last_num{
			res = (res + dfs_dieSimulator(n - 1,i,num_cnt + 1,rollMax,memo)) % (1e9 + 7)
		}else{
			res = (res + dfs_dieSimulator(n - 1,i,1,rollMax,memo)) % (1e9 + 7)
		}
	}
	if last_num >= 0{
		(*memo)[n][last_num][num_cnt] = res
	}
	return res
}

func DieSimulator(n int, rollMax []int) int{
	var memo [][6][16]int = make([][6][16]int,n)
	return dfs_dieSimulator(n,-1,0,rollMax,&memo)
}

func DieSimulator2(n int, rollMax []int) int {
	var MOD int = 1e9 + 7
	var dp [][7]int = make([][7]int,n)
	dp[0][0] = 1//dp[i][2]: 连续掷骰子i次，连续是2的排列数. dp[3][2]
	dp[0][1] = 1
	dp[0][2] = 1
	dp[0][3] = 1
	dp[0][4] = 1
	dp[0][5] = 1
	dp[0][6] = 6
	for i := 1;i < n;i++{
		for j := 0;j < 6;j++{
			dp[i][j] = dp[i - 1][6]
			if i >= rollMax[j]{
				if i >= (rollMax[j] + 1){
					dp[i][j] = (dp[i][j] - (dp[i - rollMax[j] - 1][6] - dp[i - rollMax[j] - 1][j]) + MOD) % MOD
				}else{
					dp[i][j]--
				}
			}
			dp[i][6] = (dp[i][6] + dp[i][j]) % MOD
		}
	}
	return dp[n - 1][6] % MOD
}