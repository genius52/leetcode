package number

func findChampion(grid [][]int) int {
	var l int = len(grid)
	var is_best []bool = make([]bool, l)
	for i := 0; i < l; i++ {
		is_best[i] = true
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if grid[i][j] == 1 {
				is_best[j] = false
			}
		}
	}
	for i := 0; i < l; i++ {
		if is_best[i] {
			return i
		}
	}
	return -1
}
