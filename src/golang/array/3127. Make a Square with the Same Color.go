package array

func canMakeSquare(grid [][]byte) bool {
	for r := 0; r <= 1; r++ {
		for c := 0; c <= 1; c++ {
			var w_cnt int = 0
			if grid[r][c] == 'W' {
				w_cnt++
			}
			if grid[r+1][c] == 'W' {
				w_cnt++
			}
			if grid[r][c+1] == 'W' {
				w_cnt++
			}
			if grid[r+1][c+1] == 'W' {
				w_cnt++
			}
			if w_cnt == 0 || w_cnt == 1 || w_cnt == 3 || w_cnt == 4 {
				return true
			}
		}
	}
	return false
}
