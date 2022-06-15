package diagram

func MinPathCost(grid [][]int, moveCost [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var dp [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		dp[i] = make([]int,columns)
	}
	for j := 0;j < columns;j++{
		dp[0][j] = grid[0][j]
	}

	for i := 1;i < rows;i++{
		for j := 0;j < columns;j++{
			dp[i][j] = 2147483647
			for k := 0;k < columns;k++{
				dp[i][j] = min_int(dp[i][j],grid[i][j] + dp[i - 1][k] + moveCost[grid[i - 1][k]][j])
			}
		}
	}
	var res int = 2147483647
	for j := 0;j < columns;j++{
		res = min_int(res,dp[rows - 1][j])
	}
	return res
}