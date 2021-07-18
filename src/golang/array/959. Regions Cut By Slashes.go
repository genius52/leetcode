package array


func dfs_regionsBySlashes(visited [][]bool,i int,j int)bool{
	l := len(visited)
	if i < 0 || j < 0 || i >= l || j >= l || visited[i][j]{
		return false
	}
	visited[i][j] = true
	dfs_regionsBySlashes(visited,i - 1,j)
	dfs_regionsBySlashes(visited,i + 1,j)
	dfs_regionsBySlashes(visited,i,j - 1)
	dfs_regionsBySlashes(visited,i,j + 1)
	return true
}

func regionsBySlashes(grid []string) int {
	l := len(grid)
	var visited [][]bool = make([][]bool,l * 3)
	for i := 0;i < l;i++ {
		visited[i * 3] = make([]bool,l * 3)
		visited[i * 3 + 1] = make([]bool,l * 3)
		visited[i * 3 + 2] = make([]bool,l * 3)
		for j := 0;j < l;j++{
			if grid[i][j] == '\\'{
				visited[i * 3][j * 3] = true
				visited[i * 3 + 1][j * 3 + 1] = true
				visited[i * 3 + 2][j * 3 + 2] = true
			}else if grid[i][j] == '/'{
				visited[i * 3][j * 3 + 2] = true
				visited[i * 3 + 1][j * 3 + 1] = true
				visited[i * 3 + 2][j * 3] = true
			}
		}
	}
	var res int = 0
	for i := 0;i < l * 3;i++{
		for j := 0;j < l * 3;j++{
			if dfs_regionsBySlashes(visited,i,j){
				res++
			}
		}
	}
	return res
}