package number

func dp_distinctSequences(prepre int,pre int,n int,graph [7][7]int,memo [][7][7]int)int{
	if n == 0{
		return 1
	}
	if prepre != -1 && pre != -1 && memo[n][prepre][pre] != -1{
		return memo[n][prepre][pre]
	}
	var res int = 0
	for i := 1;i <= 6;i++{
		if i != pre && i != prepre {
			if pre == -1{
				res += dp_distinctSequences(pre,i,n - 1,graph,memo)
			}else if graph[pre][i] == 1{
				res += dp_distinctSequences(pre,i,n - 1,graph,memo)
			}
			res %= 1e9 + 7
		}
	}
	if prepre != -1 && pre != -1{
		memo[n][prepre][pre] = res
	}
	return res
}

func DistinctSequences(n int) int{
	var graph [7][7]int = [7][7]int{{0,0,0,0,0,0,0},
									{0,0,1,1,1,1,1},
									{0,1,0,1,0,1,0},
									{0,1,1,0,1,1,0},
									{0,1,0,1,0,1,0},
									{0,1,1,1,1,0,1},
									{0,1,0,0,0,1,0}}
	var memo [][7][7]int = make([][7][7]int,n + 1)
	for i := 0;i <= n;i++{
		for j := 0;j < 7;j++{
			for k := 0;k < 7;k++{
				memo[i][j][k] = -1
			}
		}
	}
	return dp_distinctSequences(-1,-1,n,graph,memo)
}

//func DistinctSequences(n int) int {
//	var dp [][7]int = make([][7]int,n)
//	for i := 1;i < 7;i++{
//		dp[0][i] = 1
//	}
//	for i := 1;i < n;i++{
//		dp[i][1] = dp[i - 1][2] + dp[i - 1][3] + dp[i - 1][4] + dp[i - 1][5] + dp[i - 1][6]
//		dp[i][2] = dp[i - 1][1] + dp[i - 1][3] + dp[i - 1][5]
//		dp[i][3] = dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][4] + dp[i - 1][5]
//		dp[i][4] = dp[i - 1][1] + dp[i - 1][3] + dp[i - 1][5]
//		dp[i][5] = dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3] + dp[i - 1][4] + dp[i - 1][6]
//		dp[i][6] = dp[i - 1][1] + dp[i - 1][5]
//		if i >= 2{
//			dp[i][1] -= dp[i - 2][1]
//			dp[i][2] -= dp[i - 2][2]
//			dp[i][3] -= dp[i - 2][3]
//			dp[i][4] -= dp[i - 2][4]
//			dp[i][5] -= dp[i - 2][5]
//			dp[i][6] -= dp[i - 2][6]
//		}
//	}
//	var res int = 0
//	for i := 1;i <= 6;i++{
//		res += dp[n - 1][i]
//	}
//	return res
//}