package array

func ConstructProductMatrix(grid [][]int) [][]int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var MOD int = 12345
	var prefix [][]int = make([][]int, rows)
	var suffix [][]int = make([][]int, rows)
	var res [][]int = make([][]int, rows)
	for i := 0; i < rows; i++ {
		prefix[i] = make([]int, columns)
		suffix[i] = make([]int, columns)
		res[i] = make([]int, columns)
	}
	var sum int = 1
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			prefix[i][j] = sum
			sum *= grid[i][j]
			sum %= MOD
		}
	}
	sum = 1
	for i := rows - 1; i >= 0; i-- {
		for j := columns - 1; j >= 0; j-- {
			suffix[i][j] = sum
			sum *= grid[i][j]
			sum %= MOD
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			res[i][j] = prefix[i][j] * suffix[i][j] % MOD
		}
	}
	return res
}
