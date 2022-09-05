package number

func dfs_numberOfWays2400(cur_pos int,endPos int,k int,memo *[3000][3000]int)int{
	if k == 0 && cur_pos == endPos{
		return 1
	}
	if k == 0 || abs_int(cur_pos - endPos) > k{
		return 0
	}
	if memo[cur_pos + 1000][k] != -1{
		return memo[cur_pos + 1000][k]
	}
	var MOD int = 1e9 + 7
	memo[cur_pos + 1000][k] = dfs_numberOfWays2400(cur_pos + 1,endPos,k - 1,memo) + dfs_numberOfWays2400(cur_pos - 1,endPos,k - 1,memo)
	memo[cur_pos + 1000][k] %= MOD
	return memo[cur_pos + 1000][k]
}

func NumberOfWays2400(startPos int, endPos int, k int) int {
	var memo [3000][3000]int
	for i := 0;i < 3000;i++{
		for j := 0;j < 3000;j++{
			memo[i][j] = -1
		}
	}
	return dfs_numberOfWays2400(startPos,endPos,k,&memo)
}