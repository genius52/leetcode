package diagram

func check(grid1 [][]int,visited [][]int)bool{
	var l int = len(visited)
	for i := 0;i < l;i++{
		if grid1[visited[i][0]][visited[i][1]] == 0{
			return false
		}
	}
	return true
}

func dfs_countSubIslands(grid2 [][]int,rows int,columns int,r int,c int,visited *[][]int){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if grid2[r][c] != 1 || grid2[r][c] == 2{
		return
	}
	grid2[r][c] = 2
	*visited = append(*visited,[]int{r,c})
	dfs_countSubIslands(grid2,rows,columns,r - 1,c,visited)
	dfs_countSubIslands(grid2,rows,columns,r + 1,c,visited)
	dfs_countSubIslands(grid2,rows,columns,r,c - 1,visited)
	dfs_countSubIslands(grid2,rows,columns,r,c + 1,visited)
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
			var visited [][]int
			dfs_countSubIslands(grid2,rows,columns,i,j,&visited)
			if check(grid1,visited){
				res++
			}
		}
	}
	return res
}