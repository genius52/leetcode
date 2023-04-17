package array

import "strconv"

func findColumnWidth(grid [][]int) []int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res []int = make([]int, columns)
	for j := 0; j < columns; j++ {
		var max_width int = 0
		for i := 0; i < rows; i++ {
			s := strconv.Itoa(grid[i][j])
			var l int = len(s)
			if l > max_width {
				max_width = l
			}
		}
		res[j] = max_width
	}
	return res
}
