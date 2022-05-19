package diagram

func dfs_longestIncreasingPath(matrix [][]int,rows int,columns int,r int,c int,last_val int,dp [][]int)int{
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return 0
	}
	if matrix[r][c] <= last_val{
		return 0
	}
	if dp[r][c] != 0{
		return dp[r][c]
	}
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	var res int = 0
	for _,dir := range dirs{
		cur := dfs_longestIncreasingPath(matrix,rows,columns,r + dir[0],c + dir[1],matrix[r][c],dp)
		if cur > res{
			res = cur
		}
	}
	dp[r][c] = 1 + res
	return dp[r][c]
}

func longestIncreasingPath(matrix [][]int) int {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var dp [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		dp[i] = make([]int,columns)
	}
	var max_length int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if dp[i][j] == 0{
				cur := dfs_longestIncreasingPath(matrix,rows,columns,i,j,-1,dp)
				if cur > max_length{
					max_length = cur
				}
			}
		}
	}
	return max_length
}