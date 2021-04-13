package array

//Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
//Output: [1,2,4,7,5,3,6,8,9]
func findDiagonalOrder(mat [][]int) []int {
	var rows int = len(mat)
	var columns int = len(mat[0])
	var total int = rows * columns
	var res []int = make([]int,total)
	var index int = 0
	var r int = 0
	var c int = 0
	var row_step int = -1
	var column_step int = 1
	for index < total{
		res[index] = mat[r][c]
		index++
		r += row_step
		c += column_step
		if r >= rows{
			r = rows - 1
			c += 2
			row_step = -row_step
			column_step = -column_step
		}
		if c >= columns{
			r += 2
			c = columns - 1
			row_step = -row_step
			column_step = -column_step
		}
		if r < 0{
			r = 0
			row_step = -row_step
			column_step = -column_step
		}
		if c < 0{
			c = 0
			row_step = -row_step
			column_step = -column_step
		}
	}
	return res
}