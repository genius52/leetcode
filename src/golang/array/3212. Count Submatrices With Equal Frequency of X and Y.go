package array

func NumberOfSubmatrices(grid [][]byte) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var prefix_x [][]int = make([][]int, rows)
	var prefix_y [][]int = make([][]int, rows)
	for r := 0; r < rows; r++ {
		prefix_x[r] = make([]int, columns)
		prefix_y[r] = make([]int, columns)
	}
	if grid[0][0] == 'X' {
		prefix_x[0][0] = 1
	} else if grid[0][0] == 'Y' {
		prefix_y[0][0] = 1
	}
	var res int = 0
	for r := 1; r < rows; r++ {
		prefix_x[r][0] = prefix_x[r-1][0]
		prefix_y[r][0] = prefix_y[r-1][0]
		if grid[r][0] == 'X' {
			prefix_x[r][0]++
		} else if grid[r][0] == 'Y' {
			prefix_y[r][0]++
		}
		if prefix_x[r][0] > 0 && prefix_y[r][0] > 0 && prefix_x[r][0] == prefix_y[r][0] {
			res++
		}
	}
	for c := 1; c < columns; c++ {
		prefix_x[0][c] = prefix_x[0][c-1]
		prefix_y[0][c] = prefix_y[0][c-1]
		if grid[0][c] == 'X' {
			prefix_x[0][c]++
		} else if grid[0][c] == 'Y' {
			prefix_y[0][c]++
		}
		if prefix_x[0][c] > 0 && prefix_y[0][c] > 0 && prefix_x[0][c] == prefix_y[0][c] {
			res++
		}
	}
	for r := 1; r < rows; r++ {
		for c := 1; c < columns; c++ {
			prefix_x[r][c] = prefix_x[r-1][c] + prefix_x[r][c-1] - prefix_x[r-1][c-1]
			prefix_y[r][c] = prefix_y[r-1][c] + prefix_y[r][c-1] - prefix_y[r-1][c-1]
			if grid[r][c] == 'X' {
				prefix_x[r][c]++
			} else if grid[r][c] == 'Y' {
				prefix_y[r][c]++
			}
			if prefix_x[r][c] > 0 && prefix_y[r][c] > 0 && prefix_x[r][c] == prefix_y[r][c] {
				res++
			}
		}
	}
	return res
}
