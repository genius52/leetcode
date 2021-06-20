package diagram

func dfs_countSubIslands(grid1 [][]int,grid2 [][]int,rows int,columns int,r int,c int)bool{
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return true
	}
	if grid2[r][c] == 0{
		return true
	}
	grid2[r][c] = 0
	ret1 := dfs_countSubIslands(grid1,grid2,rows,columns,r - 1,c)
	ret2 := dfs_countSubIslands(grid1,grid2,rows,columns,r + 1,c)
	ret3 := dfs_countSubIslands(grid1,grid2,rows,columns,r,c - 1)
	ret4 := dfs_countSubIslands(grid1,grid2,rows,columns,r,c + 1)
	if grid1[r][c] == 0 || !ret1 || !ret2 || !ret3 || !ret4{
		return false
	}
	return true
}

func CountSubIslands(grid1 [][]int, grid2 [][]int) int {
	var rows int = len(grid1)
	var columns int = len(grid1[0])
	var res int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid2[i][j] != 1{
				continue
			}
			if dfs_countSubIslands(grid1,grid2,rows,columns,i,j){
				res++
			}
		}
	}
	return res
}