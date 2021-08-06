package array

//Input: grid = [[1,2,2],[2,3,2]], r0 = 0, c0 = 1, color = 3
//Output: [[1, 3, 3], [2, 3, 3]]
func dfs_colorBorder(grid [][]int,rows int,columns int,r int, c int, old_color int,color int){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if grid[r][c] != old_color{
		return
	}
	grid[r][c] = color
	dfs_colorBorder(grid,rows,columns,r - 1,c,old_color,color)
	dfs_colorBorder(grid,rows,columns,r + 1,c,old_color,color)
	dfs_colorBorder(grid,rows,columns,r,c - 1,old_color,color)
	dfs_colorBorder(grid,rows,columns,r,c + 1,old_color,color)
}

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	old := grid[r0][c0]
	if old == color{
		return grid
	}
	dfs_colorBorder(grid,rows,columns,r0,c0,old,-color)
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	var internal [][]int
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == -color{
				var is_intenal bool = true
				for _,dir := range dirs{
					next_r := i + dir[0]
					next_c := j + dir[1]
					if next_r < 0 || next_r >= rows || next_c < 0 || next_c >= columns{
						is_intenal = false
						break
					}
					if grid[next_r][next_c] != -color{
						is_intenal = false
						break
					}
				}
				if is_intenal{
					internal = append(internal,[]int{i,j})
				}
			}
		}
	}
	for _,pos := range internal{
		grid[pos[0]][pos[1]] = old
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == -color{
				grid[i][j] = color
			}
		}
	}
	return grid
}