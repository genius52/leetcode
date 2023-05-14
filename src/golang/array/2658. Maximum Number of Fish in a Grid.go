package array

func dfs_findMaxFish(grid [][]int, rows int, columns int, r int, c int) int {
	if r < 0 || r >= rows || c < 0 || c >= columns || grid[r][c] == 0 {
		return 0
	}
	var res int = grid[r][c]
	grid[r][c] = 0
	return res + dfs_findMaxFish(grid, rows, columns, r-1, c) +
		dfs_findMaxFish(grid, rows, columns, r+1, c) +
		dfs_findMaxFish(grid, rows, columns, r, c-1) +
		dfs_findMaxFish(grid, rows, columns, r, c+1)
}

func findMaxFish(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res int = 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == 0 {
				continue
			}
			cnt := dfs_findMaxFish(grid, rows, columns, i, j)
			if cnt > res {
				res = cnt
			}
		}
	}
	return res
}
