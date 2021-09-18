package array

func DiagonalSort(mat [][]int) [][]int {
	rows := len(mat)
	if rows <= 1{
		return mat
	}
	columns := len(mat[0])
	if columns <= 1{
		return mat
	}
	//把最大的数往下移动
	for r := rows - 1;r >= 0;r--{
		for i := 0;i < r;i++{
			for j := 0;j < columns - 1;j++{
				if mat[i][j] > mat[i + 1][j + 1]{
					mat[i][j], mat[i + 1][j + 1] = mat[i + 1][j + 1],mat[i][j]
				}
			}
		}
	}
	return mat
}