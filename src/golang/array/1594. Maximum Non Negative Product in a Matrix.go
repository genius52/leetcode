package array

import "math"

var mod int64 = 1000000007
func dfs_maxProductPath(grid [][]int,rows int,columns int,r int,c int,product int64)(int64){
	if r == rows - 1 && c == columns - 1{
		return product * int64(grid[r][c])
	}
	product = product * int64(grid[r][c])
	if product == 0{
		return 0
	}
	if r < rows - 1 && c < columns - 1{
		move_down := dfs_maxProductPath(grid,rows,columns,r + 1,c,product)
		move_right := dfs_maxProductPath(grid,rows,columns,r,c + 1,product)
		if move_down > move_right{
			return move_down
		}else {
			return move_right
		}
	}else if c < columns - 1{
		return dfs_maxProductPath(grid,rows,columns,r,c + 1,product)
	}else if r < rows - 1{
		return dfs_maxProductPath(grid,rows,columns,r + 1,c,product)
	}
	return 0
}

func MaxProductPath(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])
	res := dfs_maxProductPath(grid,rows,columns,0,0,1)
	if res < 0 {
		return -1
	}
	return int(res) % int(mod)
}

func min_int64_number(nums ...int64)int64{
	var min int64 = math.MaxInt64
	for _,n := range nums{
		if n < min{
			min = n
		}
	}
	return min
}

func max_int64_number(nums ...int64)int64{
	var max int64 = math.MinInt64
	for _,n := range nums{
		if n > max{
			max = n
		}
	}
	return max
}

//DP from bottom to top
func MaxProductPath2(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])
	var dp [][][2]int64 = make([][][2]int64,rows)
	for i := 0;i < rows;i++{
		dp[i] = make([][2]int64,columns)
	}
	dp[0][0][0] = int64(grid[0][0])//min product
	dp[0][0][1] = int64(grid[0][0])//max product
	var MOD int = 1e9 + 7
	for i := 1;i < rows;i++{
		dp[i][0][0] = min_int64_number(dp[i - 1][0][0] * int64(grid[i][0]),dp[i - 1][0][1] * int64(grid[i][0]))
		dp[i][0][1] = max_int64_number(dp[i - 1][0][0] * int64(grid[i][0]),dp[i - 1][0][1] * int64(grid[i][0]))
	}
	for i := 1;i < columns;i++{
		dp[0][i][0] = min_int64_number(dp[0][i - 1][0] * int64(grid[0][i]),dp[0][i - 1][1] * int64(grid[0][i]))
		dp[0][i][1] = max_int64_number(dp[0][i - 1][0] * int64(grid[0][i]),dp[0][i - 1][1] * int64(grid[0][i]))
	}
	for i := 1;i < rows;i++{
		for j := 1;j < columns;j++{
			dp[i][j][0] = min_int64_number(int64(grid[i][j]) * dp[i - 1][j][0],
				int64(grid[i][j]) * dp[i][j - 1][0],int64(grid[i][j]) * dp[i - 1][j][1],
				int64(grid[i][j]) * dp[i][j - 1][1])
			dp[i][j][1] = max_int64_number(int64(grid[i][j]) * dp[i - 1][j][0],
				int64(grid[i][j]) * dp[i][j - 1][0],int64(grid[i][j]) * dp[i - 1][j][1],
				int64(grid[i][j]) * dp[i][j - 1][1])
		}
	}
	if dp[rows - 1][columns - 1][1] < 0{
		return -1
	}
	return int(dp[rows - 1][columns - 1][1]) % MOD
}