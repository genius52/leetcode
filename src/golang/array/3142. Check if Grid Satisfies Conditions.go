package array

func satisfiesConditions(grid [][]int) bool {
	var rows int = len(grid)
	var columns int = len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i+1 < rows && grid[i][j] != grid[i+1][j] {
				return false
			}
			if j+1 < columns && grid[i][j] == grid[i][j+1] {
				return false
			}
		}
	}
	return true
}
