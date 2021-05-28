package array

func isToeplitzMatrix(matrix [][]int) bool {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	for i := 0;i < rows;i++{
		var target int = matrix[i][0]
		for j := 1;j < columns && (i + j) < rows;j++{
			if matrix[i + j][j] != target{
				return false
			}
		}
	}
	for j := 1;j < columns;j++{
		var target int = matrix[0][j]
		for i := 0;i < rows && (i + j) < columns;i++{
			if matrix[i][i + j] != target{
				return false
			}
		}
	}
	return true
}