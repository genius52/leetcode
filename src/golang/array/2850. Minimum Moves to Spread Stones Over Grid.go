package array

func dfs_minimumMoves(grid [][]int, zero [][2]int, l1 int, idx int, full [][2]int, l2 int) int {
	if idx == l1 {
		return 0
	}
	var res int = 999
	for i := 0; i < l2; i++ {
		x := full[i][0]
		y := full[i][1]
		if grid[x][y] == 1 {
			continue
		}
		grid[x][y] -= 1
		cur := abs_int(zero[idx][0]-x) + abs_int(zero[idx][1]-y) + dfs_minimumMoves(grid, zero, l1, idx+1, full, l2)
		if cur < res {
			res = cur
		}
		grid[x][y] += 1
	}
	return res
}

func minimumMoves(grid [][]int) int {
	var zero [][2]int
	var full [][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i][j] == 0 {
				zero = append(zero, [2]int{i, j})
			}
			if grid[i][j] > 1 {
				full = append(full, [2]int{i, j})
			}
		}
	}
	return dfs_minimumMoves(grid, zero, len(zero), 0, full, len(full))
}
