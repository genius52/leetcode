package array

//TLE
func PossibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var dp [][]int = make([][]int,rows)
	var stamps [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		dp[i] = make([]int,columns)
		stamps[i] = make([]int,columns)
	}
	dp[0][0] = grid[0][0]
	for i := 1;i < columns;i++{
		dp[0][i] += dp[0][i - 1] + grid[0][i]
	}
	for i := 1;i < rows;i++{
		dp[i][0] += dp[i - 1][0] + + grid[i][0]
	}
	for i := 1;i < rows;i++{
		for j := 1;j < columns;j++{
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1] - dp[i - 1][j - 1] + grid[i][j]
		}
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++ {
			if grid[i][j] == 1 || stamps[i][j] == 1{
				continue
			}
			if i + 1 < stampHeight || j + 1 < stampWidth{
				continue
			}
			pre_r := i - stampHeight
			pre_c := j - stampWidth
			area := dp[i][j]
			if pre_r >= 0{
				area -= dp[pre_r][j]
			}
			if pre_c >= 0{
				area -= dp[i][pre_c]
			}
			if pre_r >= 0 && pre_c >= 0{
				area += dp[pre_r][pre_c]
			}
			if area == 0{
				for m := pre_r + 1;m <= i;m++{
					for n := pre_c + 1;n <= j;n++{
						stamps[m][n] = 1
					}
				}
			}
		}
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 0 && stamps[i][j] == 0{
				return false
			}
		}
	}
	return true
}