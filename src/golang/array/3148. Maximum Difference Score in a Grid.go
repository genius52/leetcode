package array

func maxScore3148(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var prefix_max [][]int = make([][]int, rows)
	for i := 0; i < rows; i++ {
		prefix_max[i] = make([]int, columns)
	}
	prefix_max[rows-1][columns-1] = grid[rows-1][columns-1]
	for c := columns - 2; c >= 0; c-- {
		prefix_max[rows-1][c] = max_int(grid[rows-1][c], prefix_max[rows-1][c+1])
	}
	for r := rows - 2; r >= 0; r-- {
		prefix_max[r][columns-1] = max_int(grid[r][columns-1], prefix_max[r+1][columns-1])
	}
	var res int = -(1 << 31)
	for r := rows - 2; r >= 0; r-- {
		for c := columns - 2; c >= 0; c-- {
			prefix_max[r][c] = max_int_number(grid[r][c], prefix_max[r+1][c], prefix_max[r][c+1])
		}
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			var diff1 int = -(1 << 31)
			var diff2 int = -(1 << 31)
			if r != rows-1 {
				diff1 = prefix_max[r+1][c] - grid[r][c]
			}
			if c != columns-1 {
				diff2 = prefix_max[r][c+1] - grid[r][c]
			}
			cur_diff := max_int(diff1, diff2)
			if cur_diff > res {
				res = cur_diff
			}
		}
	}
	return res
}
