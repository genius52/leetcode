package array

func FindMissingAndRepeatedValues(grid [][]int) []int {
	var l int = len(grid)
	var res []int = make([]int, 2)
	var visited []bool = make([]bool, l*l+1)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if visited[grid[i][j]] {
				res[0] = grid[i][j]
			}
			visited[grid[i][j]] = true
		}
	}
	for i := 1; i <= l*l; i++ {
		if !visited[i] {
			res[1] = i
		}
	}
	return res
}
