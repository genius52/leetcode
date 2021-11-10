package array

func diagonalSum(mat [][]int) int {
	var rows int = len(mat)
	var sum int = 0
	for i := 0;i < rows;i++{
		sum += mat[i][i]
		sum += mat[i][rows - 1 - i]
	}
	if rows | 1 == rows{
		sum -= mat[rows/2][rows/2]
	}
	return sum
}