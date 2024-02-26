package array

func modifiedMatrix(matrix [][]int) [][]int {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var res [][]int = make([][]int, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, columns)
	}
	for j := 0; j < columns; j++ {
		var minus_one []int
		var max_value int = -1
		for i := 0; i < rows; i++ {
			if matrix[i][j] == -1 {
				minus_one = append(minus_one, i)
			} else {
				res[i][j] = matrix[i][j]
				if matrix[i][j] > max_value {
					max_value = matrix[i][j]
				}
			}
		}
		for _, idx := range minus_one {
			res[idx][j] = max_value
		}
	}
	return res
}
