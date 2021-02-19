package array

//Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//Output: true
func searchMatrix(matrix [][]int, target int) bool {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var low int = 0
	var high int = rows * columns - 1
	for low <= high{
		mid := low + (high - low)/2
		r := mid / columns
		c := mid % columns
		if matrix[r][c] == target{
			return true
		}else if matrix[r][c] < target{
			low = mid + 1
		}else{
			high = mid - 1
		}
	}
	return false
}