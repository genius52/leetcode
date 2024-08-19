package array

func minFlips(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res int = rows * columns
	var cnt int = 0
	for r := 0; r < rows; r++ {
		for c := 0; c < columns/2; c++ {
			if grid[r][c] != grid[r][columns-1-c] {
				cnt++
			}
		}
	}
	res = min(res, cnt)
	cnt = 0
	for c := 0; c < columns; c++ {
		for r := 0; r < rows/2; r++ {
			if grid[r][c] != grid[rows-1-r][c] {
				cnt++
			}
		}
	}
	res = min(res, cnt)
	return res
}
