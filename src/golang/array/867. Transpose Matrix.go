package array

//Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
//Output: [[1,4,7],[2,5,8],[3,6,9]]
func transpose(matrix [][]int) [][]int {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var res [][]int = make([][]int,columns)
	for j := 0;j < columns;j++{
		res[j] = make([]int,rows)
		for i := 0;i < rows;i++{
			res[j][i] = matrix[i][j]
		}
	}
	return res
}