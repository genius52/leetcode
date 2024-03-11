package array

func CountSubmatrices(grid [][]int, k int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var prefix [][]int = make([][]int, rows)
	for i := 0; i < rows; i++ {
		prefix[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			prefix[i][j] = grid[i][j]
		}
	}
	for i := 1; i < rows; i++ {
		prefix[i][0] += prefix[i-1][0]
	}
	for j := 1; j < columns; j++ {
		prefix[0][j] += prefix[0][j-1]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			prefix[i][j] += prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1]
		}
	}
	var res int = 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if prefix[i][j] <= k {
				res++
			}
		}
	}
	return res
}
