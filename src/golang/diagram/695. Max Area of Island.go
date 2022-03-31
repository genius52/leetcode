package diagram

func dfs_maxAreaOfIsland(grid [][]int,rows int,columns int,r int,c int)int{
	if r < 0 || c < 0 || r >= rows || c >= columns{
		return 0
	}
	if grid[r][c] == 0{
		return 0
	}
	grid[r][c] = 0
	return 1 + dfs_maxAreaOfIsland(grid,rows,columns,r - 1,c) + dfs_maxAreaOfIsland(grid,rows,columns,r + 1,c) +
		dfs_maxAreaOfIsland(grid,rows,columns,r,c - 1) + dfs_maxAreaOfIsland(grid,rows,columns,r,c + 1)
}

func maxAreaOfIsland(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res int = 0
	for r := 0;r < rows;r++{
		for c := 0;c < columns;c++{
			if grid[r][c] != 1{
				continue
			}
			areas := dfs_maxAreaOfIsland(grid,rows,columns,r,c)
			if areas > res{
				res = areas
			}
		}
	}
	return res
}