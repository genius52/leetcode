package array

//200
//Input:
//11000
//11000
//00100
//00011
//
//Output: 3
func dfs_numIslands(grid [][]byte,row int,col int){
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != '1'{
		return
	}
	grid[row][col] = '0'
	dfs_numIslands(grid,row - 1,col)
	dfs_numIslands(grid,row + 1,col)
	dfs_numIslands(grid,row,col - 1)
	dfs_numIslands(grid,row,col + 1)
}

func numIslands(grid [][]byte) int {
	var res int = 0
	for i := 0;i < len(grid);i++{
		for j := 0;j < len(grid[0]);j++{
			if grid[i][j] == '1'{
				dfs_numIslands(grid,i,j)
				res++
			}
		}
	}
	return res
}