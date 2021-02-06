package array

//Given input matrix =
//[
//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//],
//
//rotate the input matrix in-place such that it becomes:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]
//]
func rotate(matrix [][]int)  {
	//swap by lines
	rows := len(matrix)
	for row := 0;row < rows/2;row++{
		matrix[row],matrix[rows - 1 - row] = matrix[rows - 1 - row],matrix[row]
	}
	//swap by diagonal line
	for i := 1;i < rows;i++{
		for j := 0;j < i;j++{
			matrix[i][j],matrix[j][i] = matrix[j][i],matrix[i][j]
		}
	}
}
