package diagram

//1 represents the starting square.  There is exactly one starting square.
//2 represents the ending square.  There is exactly one ending square.
//0 represents empty squares we can walk over.
//-1 represents obstacles that we cannot walk over.
func dfs_uniquePathsIII(grid [][]int,row int,column int,visited [][]int, total *int){
	rows := len(grid)
	columns := len(grid[0])
	if row < 0 || row >= rows || column < 0 || column >= columns || visited[row][column] != 0{
		return
	}
	if grid[row][column] == 2{
		for i := 0;i < rows;i++{
			for j := 0;j < columns;j++{
				if i == row && j == column{
					continue
				}
				if visited[i][j] == 0{
					return
				}
			}
		}
		*total++
		return
	}
	if grid[row][column] == -1{
		visited[row][column] = -1
		return
	}
	visited[row][column] = 1
	dfs_uniquePathsIII(grid,row - 1,column,visited,total)
	dfs_uniquePathsIII(grid,row,column - 1,visited,total)
	dfs_uniquePathsIII(grid,row + 1,column,visited,total)
	dfs_uniquePathsIII(grid,row,column + 1,visited,total)
	visited[row][column] = 0
}

func reset_visit(visited [][]int){
	for i := 0;i < len(visited);i++{
		for j := 0;j < len(visited[0]);j++ {
			if visited[i][j] == 1{
				visited[i][j] = 0
			}
		}
	}
}

func uniquePathsIII(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])
	var visited [][]int = make([][]int,rows)
	for ii := 0;ii < rows;ii++{
		visited[ii] = make([]int,columns)
	}
	var total int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] != 1{
				continue
			}
			visited[i][j] = 1
			dfs_uniquePathsIII(grid,i - 1,j,visited,&total)
			reset_visit(visited)
			visited[i][j] = 1
			dfs_uniquePathsIII(grid,i + 1,j,visited,&total)
			reset_visit(visited)
			visited[i][j] = 1
			dfs_uniquePathsIII(grid,i,j - 1,visited,&total)
			reset_visit(visited)
			visited[i][j] = 1
			dfs_uniquePathsIII(grid,i,j + 1,visited,&total)
			break
		}
	}
	return total
}