package diagram

//Input: grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
//Output: 1
func dfs_closedIsland(grid [][]int,rows int,columns int,row int,col int,ori_val int,mark_val int)bool{
	if row < 0 || row >= rows || col < 0 || col >= columns {
		return false
	}
	if grid[row][col] != ori_val{
		return false
	}
	grid[row][col] = mark_val
	dfs_closedIsland(grid,rows,columns,row - 1,col,ori_val,mark_val)
	dfs_closedIsland(grid,rows,columns,row + 1,col,ori_val,mark_val)
	dfs_closedIsland(grid,rows,columns,row,col - 1,ori_val,mark_val)
	dfs_closedIsland(grid,rows,columns,row,col + 1,ori_val,mark_val)
	return true
}

func ClosedIsland(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	for i := 0;i < rows;i++{
		dfs_closedIsland(grid,rows,columns,i,0,0,1)
		dfs_closedIsland(grid,rows,columns,i,columns - 1,0,1)
	}
	for j := 0;j < columns;j++{
		dfs_closedIsland(grid,rows,columns,0,j,0,1)
		dfs_closedIsland(grid,rows,columns,rows - 1,j,0,1)
	}
	var cnt int = 0
	for i := 1;i < rows - 1;i++{
		for j := 1;j < columns - 1;j++{
			ret := dfs_closedIsland(grid,rows,columns,i,j,0,1)
			if ret{
				cnt++
			}
		}
	}
	return cnt
}