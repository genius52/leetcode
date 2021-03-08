package array

func SearchMatrix240(matrix [][]int, target int) bool {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var r int = 0
	var c int = columns - 1
	for r >= 0 && r < rows && c >= 0 && c < columns{
		if matrix[r][c] == target{
			return true
		}
		if matrix[r][c] > target{
			c--
		}else{
			r++
		}
	}
	return false
}