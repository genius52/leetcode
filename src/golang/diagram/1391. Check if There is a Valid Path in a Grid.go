package diagram

//1 which means a street connecting the left cell and the right cell.
//2 which means a street connecting the upper cell and the lower cell.
//3 which means a street connecting the left cell and the lower cell.
//4 which means a street connecting the right cell and the lower cell.
//5 which means a street connecting the left cell and the upper cell.
//6 which means a street connecting the right cell and the upper cell.
func dfs_hasValidPath(grid *[][]int,rows int,columns int,r int,c int,pre_r int,pre_c int,visited *[][]bool)bool{
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return false
	}
	if (*visited)[r][c] {
		return false
	}
	var dirs [6][2][2]int = [6][2][2]int{{{0,-1},{0,1}},{{-1,0},{1,0}},{{0,-1},{1,0}},
									{{1,0},{0,1}},{{-1,0},{0,-1}},{{-1,0},{0,1}}}
	if pre_r != -1 && pre_c != -1{
		var connected bool = false
		for i := 0;i < 2;i++{
			next_r := r + dirs[(*grid)[r][c] - 1][i][0]
			next_c := c + dirs[(*grid)[r][c] - 1][i][1]
			if next_r == pre_r && next_c == pre_c{
				connected = true
				break
			}
		}
		if !connected{
			return false
		}
	}
	(*visited)[r][c] = true
	if r == rows - 1 && c == columns - 1{
		return true
	}
	for i := 0;i < 2;i++{
		next_r := r + dirs[(*grid)[r][c] - 1][i][0]
		next_c := c + dirs[(*grid)[r][c] - 1][i][1]
		res := dfs_hasValidPath(grid,rows,columns,next_r,next_c,r,c,visited)
		if res{
			return true
		}
	}
	return false
}

func HasValidPath(grid [][]int) bool {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var visited [][]bool = make([][]bool,rows)
	for i := 0;i < rows;i++{
		visited[i] = make([]bool,columns)
	}
	return dfs_hasValidPath(&grid,rows,columns,0,0,-1,-1,&visited)
}