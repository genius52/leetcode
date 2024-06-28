package diagram

func minimumArea(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var up int = rows - 1
	var down int = 0
	var left int = columns - 1
	var right int = 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == 1 {
				if i < up {
					up = i
				}
				if i > down {
					down = i
				}
				if j < left {
					left = j
				}
				if j > right {
					right = j
				}
			}
		}
	}
	return (right - left + 1) * (down - up + 1)
}
