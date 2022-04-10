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

func SearchMatrix(matrix [][]int, target int) bool{
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	for i := 0;i < rows;i++{
		if target < matrix[i][0]{
			break
		}
		if target > matrix[i][columns - 1]{
			continue
		}
		var left int = 0
		var right int = columns - 1
		for left <= right{
			mid := (left + right)/2
			if matrix[i][mid] == target{
				return true
			}else if matrix[i][mid] < target{
				left = mid + 1
			}else if matrix[i][mid] > target{
				right = mid - 1
			}
		}
	}
	return false
}