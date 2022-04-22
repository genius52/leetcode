package diagram

//Input: grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
//Output: 1
func dfs_closedIsland(grid [][]int,rows int,columns int,row int,col int){
	if row < 0 || row >= rows || col < 0 || col >= columns {
		return
	}
	if grid[row][col] != 0{
		return
	}
	grid[row][col] = 1
	dfs_closedIsland(grid,rows,columns,row - 1,col)
	dfs_closedIsland(grid,rows,columns,row + 1,col)
	dfs_closedIsland(grid,rows,columns,row,col - 1)
	dfs_closedIsland(grid,rows,columns,row,col + 1)
}

func ClosedIsland(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	for i := 0;i < rows;i++{
		dfs_closedIsland(grid,rows,columns,i,0)
		dfs_closedIsland(grid,rows,columns,i,columns - 1)
	}
	for j := 0;j < columns;j++{
		dfs_closedIsland(grid,rows,columns,0,j)
		dfs_closedIsland(grid,rows,columns,rows - 1,j)
	}
	var cnt int = 0
	for i := 1;i < rows - 1;i++{
		for j := 1;j < columns - 1;j++{
			if grid[i][j] == 0{
				dfs_closedIsland(grid,rows,columns,i,j)
				cnt++
			}
		}
	}
	return cnt
}