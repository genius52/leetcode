package array

func dfs_isPossibleToCutPath(grid [][]int,rows int,columns int,r int,c int,visited [][]bool)bool{
	if visited[r][c] {
		return false
	}
	visited[r][c] = true
	if r == rows - 1 && c == columns - 1{
		return true
	}
	var dirs [][]int = [][]int{{1,0},{0,1}}
	for _,dir := range dirs{
		next_r := r + dir[0]
		next_c := c + dir[1]
		if next_r < 0 || next_r >= rows || next_c < 0 || next_c >= columns || visited[next_r][next_c] || grid[next_r][next_c] == 0{
			continue
		}
		ret := dfs_isPossibleToCutPath(grid,rows,columns,next_r,next_c,visited)
		if ret{
			grid[r][c] = 0
			return true
		}
	}
	return false
}

func IsPossibleToCutPath(grid [][]int) bool {
	var rows int = len(grid)
	var columns int = len(grid[0])
	dis := rows - 1 +  columns - 1
	if dis == 1{
		return false
	}
	var visited [][]bool = make([][]bool,rows)
	for i := 0;i < rows;i++{
		visited[i] = make([]bool,columns)
	}
	ret := dfs_isPossibleToCutPath(grid,rows,columns,0,0,visited)
	if !ret{
		return true
	}
	grid[0][0] = 1
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			visited[i][j] = false
		}
	}
	return !dfs_isPossibleToCutPath(grid,rows,columns,0,0,visited)
}