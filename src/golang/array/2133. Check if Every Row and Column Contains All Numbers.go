package array

func checkValid(matrix [][]int) bool {
	var n int = len(matrix)
	for r := 0; r < n; r++ {
		var record []bool = make([]bool, n)
		for c := 0; c < n; c++ {
			if record[matrix[r][c]-1] {
				return false
			}
			record[matrix[r][c]-1] = true
		}
	}
	for c := 0; c < n; c++ {
		var record []bool = make([]bool, n)
		for r := 0; r < n; r++ {
			if record[matrix[r][c]-1] {
				return false
			}
			record[matrix[r][c]-1] = true
		}
	}
	return true
}
