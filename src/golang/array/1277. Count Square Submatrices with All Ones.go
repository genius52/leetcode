package array

func countSquares(matrix [][]int) int {
	var cnt int = 0
	for i := 1;i < len(matrix);i++{
		for j := 1;j < len(matrix[0]);j++{
			if matrix[i][j] == 0{
				continue
			}
			matrix[i][j] = 1 + min_int_number(matrix[i-1][j],matrix[i][j-1],matrix[i-1][j-1])
		}
	}
	for i := 0;i < len(matrix);i++{
		for j := 0;j < len(matrix[0]);j++{
			cnt += matrix[i][j]
		}
	}
	return cnt
}