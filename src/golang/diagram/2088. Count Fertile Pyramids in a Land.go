package diagram

func CountPyramids(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var grid2 [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		grid2[i] = make([]int,columns)
		for j := 0;j < columns;j++{
			grid2[i][j] = grid[i][j]
		}
	}
	var res int = 0
	for i := 1;i < rows;i++{
		for j := 1;j < columns - 1;j++{
			if grid[i][j] == 0{
				continue
			}
			if grid[i - 1][j - 1] != 0 && grid[i - 1][j] != 0 && grid[i - 1][j + 1] != 0{
				grid[i][j] = 1 + min_int_number(grid[i - 1][j - 1],grid[i - 1][j],grid[i - 1][j + 1])
				res += grid[i][j] - 1
			}
		}
	}
	for i := 0;i < rows/2;i++{
		grid2[i],grid2[rows - 1 - i] = grid2[rows - 1 - i],grid2[i]
	}
	for i := 1;i < rows;i++{
		for j := 1;j < columns - 1;j++{
			if grid2[i][j] == 0{
				continue
			}
			if grid2[i - 1][j - 1] != 0 && grid2[i - 1][j] != 0 && grid2[i - 1][j + 1] != 0{
				grid2[i][j] = 1 + min_int_number(grid2[i - 1][j - 1],grid2[i - 1][j],grid2[i - 1][j + 1])
				res += grid2[i][j] - 1
			}
		}
	}
	return res
}