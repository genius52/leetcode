package diagram

//输入：grid = [[0,1,1,0],[0,1,1,0],[0,0,0,0]]
//输出：2
//解释：至少需要 2 天才能得到分离的陆地。
//将陆地 grid[1][1] 和 grid[0][2] 更改为水，得到两个分离的岛屿。


//Input: grid = [[1,1,0,1,1],
//               [1,1,1,1,1],
//               [1,1,0,1,1],
//               [1,1,1,1,1]]
//Output: 2
func dfs_minDays(grid [][]int,rows int,columns int,r int,c int)bool{
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return false
	}
	if grid[r][c] == 0 || grid[r][c] == -1{
		return false
	}
	grid[r][c] = -1
	dfs_minDays(grid,rows,columns,r - 1,c)
	dfs_minDays(grid,rows,columns,r + 1,c)
	dfs_minDays(grid,rows,columns,r,c + 1)
	dfs_minDays(grid,rows,columns,r,c - 1)
	return true
}

func count_island(grid [][]int,rows int,columns int)int{
	var res int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 0 || grid[i][j] == -1{
				continue
			}
			if dfs_minDays(grid,rows,columns,i,j){
				res++
			}

		}
	}
	return res
}

func MinDays(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var grid2 [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		grid2[i] = make([]int,columns)
		for j := 0;j < columns;j++{
			grid2[i][j] = grid[i][j]
		}
	}
	island_cnt := count_island(grid2,rows,columns)
	if island_cnt == 0 || island_cnt >= 2{
		return 0
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			grid2[i][j] = grid[i][j]
		}
	}
	for i := 0;i < rows;i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == 0{
				continue
			}
			grid2[i][j] = 0
			island_cnt = count_island(grid2,rows,columns)
			if island_cnt == 0 || island_cnt > 1{
				return 1
			}
			for i := 0;i < rows;i++{
				for j := 0;j < columns;j++{
					grid2[i][j] = grid[i][j]
				}
			}
		}
	}
	return 2
}