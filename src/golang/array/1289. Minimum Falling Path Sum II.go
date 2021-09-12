package array

func MinFallingPathSum2(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var dp [][]int = make([][]int,rows)//dp[i][j]: 从0行到i行，j列最小的和
	for i := 0;i < rows;i++{
		dp[i] = make([]int,columns)
	}
	var pre_first int = 2147483647
	var pre_first_idx int = -1
	var pre_second int = 2147483647
	for j := 0;j < columns;j++{
		if grid[0][j] < pre_first{
			pre_second = pre_first
			pre_first = grid[0][j]
			pre_first_idx = j
		}else{
			if grid[0][j] < pre_second{
				pre_second = grid[0][j]
			}
		}
	}
	for i := 1;i < rows;i++{
		var cur_first int = 2147483647
		var cur_first_idx int = -1
		var cur_second int = 2147483647
		for j := 0;j < columns;j++{
			if j != pre_first_idx{
				grid[i][j] += pre_first
			}else{
				grid[i][j] += pre_second
			}
			if grid[i][j] < cur_first{
				cur_second = cur_first
				cur_first = grid[i][j]
				cur_first_idx = j
			}else{
				if grid[i][j] < cur_second{
					cur_second = grid[i][j]
				}
			}
		}
		pre_first = cur_first
		pre_first_idx = cur_first_idx
		pre_second = cur_second
	}
	var res int = 2147483647
	for j := 0;j < columns;j++{
		if grid[rows - 1][j] < res{
			res = grid[rows - 1][j]
		}
	}
	return res
}