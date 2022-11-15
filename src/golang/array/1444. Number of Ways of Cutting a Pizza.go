package array

func dp_ways(dp [][]int,rows int,columns int,r int,c int,k int,memo [][][]int)int{
	if k == 0{
		return 1
	}
	if memo[r][c][k] != -1{
		return memo[r][c][k]
	}
	var res int = 0
	var MOD int = 1e9 + 7
	//水平切
	for i := r + 1;i < rows;i++{
		if dp[i][c] == 0{
			break
		}
		cnt := dp[r][c] - dp[i][c]
		if cnt == 0{
			continue
		}
		res += dp_ways(dp,rows,columns,i,c,k - 1,memo)
		res %= MOD
	}
	//垂直切
	for j := c + 1;j < columns;j++{
		if dp[r][j] == 0{
			break
		}
		cnt := dp[r][c] - dp[r][j]
		if cnt == 0{
			continue
		}
		res += dp_ways(dp,rows,columns,r,j,k - 1,memo)
		res %= MOD
	}
	memo[r][c][k] = res
	return res
}

func Ways(pizza []string, k int) int {
	var rows int = len(pizza)
	var columns int = len(pizza[0])
	var dp [][]int = make([][]int,rows)
	var memo [][][]int = make([][][]int,rows)
	for i := 0;i < rows;i++{
		dp[i] = make([]int,columns)
		memo[i] = make([][]int,columns)
		for j := 0;j < columns;j++{
			memo[i][j] = make([]int,k)
			for m := 0;m < k;m++{
				memo[i][j][m] = -1
			}
		}
	}
	if pizza[rows - 1][columns - 1] == 'A'{
		dp[rows - 1][columns - 1] = 1
	}
	for i := rows - 2;i >= 0;i--{
		dp[i][columns - 1] = dp[i + 1][columns - 1]
		if pizza[i][columns - 1] == 'A'{
			dp[i][columns - 1]++
		}
	}
	for j := columns - 2;j >= 0;j--{
		dp[rows - 1][j] = dp[rows - 1][j + 1]
		if pizza[rows - 1][j] == 'A'{
			dp[rows - 1][j]++
		}
	}
	for i := rows - 2;i >= 0;i--{
		for j := columns - 2;j >= 0;j--{
			dp[i][j] = dp[i + 1][j] + dp[i][j + 1] - dp[i + 1][j + 1]
			if pizza[i][j] == 'A'{
				dp[i][j]++
			}
		}
	}
	return dp_ways(dp,rows,columns,0,0,k - 1,memo)
}