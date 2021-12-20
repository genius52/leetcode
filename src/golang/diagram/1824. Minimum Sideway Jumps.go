package diagram

//输入：obstacles = [0,1,2,3,0]
//输出：2

//DP from bottom to top
func MinSideJumps2(obstacles []int) int {
	var l int = len(obstacles)
	var dp [][3]int = make([][3]int, l)
	dp[0][0] = 1
	dp[0][1] = 0
	dp[0][2] = 1
	for i := 1; i < l; i++ {
		if obstacles[i] == 0 {
			dp[i][0] = min_int_number(dp[i-1][0], 1+dp[i-1][1], 1+dp[i-1][2])
			dp[i][1] = min_int_number(dp[i-1][1], 1+dp[i-1][0], 1+dp[i-1][2])
			dp[i][2] = min_int_number(dp[i-1][2], 1+dp[i-1][0], 1+dp[i-1][1])
		} else {
			dp[i][obstacles[i]-1] = 1e6
			if obstacles[i] == 1 { //block on first line
				if obstacles[i-1] == 2 {
					dp[i][1] = 1e6
				} else {
					dp[i][1] = min_int_number(dp[i-1][1], dp[i-1][0]+1, dp[i-1][2]+1) //second line
				}
				if obstacles[i-1] == 3 {
					dp[i][2] = 1e6
				} else {
					dp[i][2] = min_int_number(dp[i-1][2], dp[i-1][0]+1, dp[i-1][1]+1) //third line
				}
			} else if obstacles[i] == 2 {
				if obstacles[i-1] == 1 {
					dp[i][0] = 1e6
				} else {
					dp[i][0] = min_int_number(dp[i-1][0], dp[i-1][1]+1, dp[i-1][2]+1)
				}
				if obstacles[i-1] == 3 {
					dp[i][2] = 1e6
				} else {
					dp[i][2] = min_int_number(dp[i-1][2], dp[i-1][0]+1, dp[i-1][1]+1)
				}
			} else if obstacles[i] == 3 {
				if obstacles[i-1] == 1 {
					dp[i][0] = 1e6
				} else {
					dp[i][0] = min_int_number(dp[i-1][0], dp[i-1][1]+1, dp[i-1][2]+1)
				}
				if obstacles[i-1] == 2 {
					dp[i][1] = 1e6
				} else {
					dp[i][1] = min_int_number(dp[i-1][1], dp[i-1][0]+1, dp[i-1][2]+1)
				}
			}
		}
	}
	return min_int_number(dp[l-1][0], dp[l-1][1], dp[l-1][2])
}

//DP from top to bottom
func dfs_minSideJumps(obstacles []int, l int, pos int, line int, memo [][3]int) int {
	if pos == l-1 {
		return 0
	}
	if memo[pos][line-1] != 0 {
		return memo[pos][line-1]
	}
	if obstacles[pos+1] == line {
		var res int = 2147483647
		for i := 1; i <= 3; i++ {
			if i != line && obstacles[pos] != i {
				steps := dfs_minSideJumps(obstacles, l, pos+1, i, memo)
				if steps != 2147483647 {
					res = min_int(res, 1+steps)
				}
			}
		}
		memo[pos][line-1] = res
	} else {
		memo[pos][line-1] = dfs_minSideJumps(obstacles, l, pos+1, line, memo)
	}
	return memo[pos][line-1]
}

func minSideJumps(obstacles []int) int {
	var l int = len(obstacles)
	var memo [][3]int = make([][3]int, l)
	return dfs_minSideJumps(obstacles, l, 0, 2, memo)
}
