package array

func numberOfRightTriangles(grid [][]int) int64 {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res int64 = 0
	var r_cnt []int = make([]int, rows)
	var c_cnt []int = make([]int, columns)
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if grid[r][c] == 1 {
				r_cnt[r]++
				c_cnt[c]++
			}
		}
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if grid[r][c] == 1 {
				res += int64(r_cnt[r]-1) * int64(c_cnt[c]-1)
			}
		}
	}
	return res
}
