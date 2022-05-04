package array

func matrixBlockSum(mat [][]int, K int) [][]int {
	rows := len(mat)
	columns := len(mat[0])
	var sum [][]int = make([][]int,rows)//sum[i][j] means sum numbers from "0,0" to "i,j"
	for i := 0;i < rows;i++{
		sum[i] = make([]int,columns)
		for j := 0;j < columns;j++{
			upper_sum := 0
			if i - 1 >= 0{
				upper_sum =  sum[i - 1][j]
			}
			left_sum := 0
			if j - 1 >= 0{
				left_sum = sum[i][j - 1]
			}
			dup_sum := 0
			if j - 1 >= 0 && i - 1 >= 0{
				dup_sum = sum[i - 1][j - 1]
			}
			sum[i][j] = upper_sum + left_sum - dup_sum + mat[i][j]
		}
	}
	var res [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		res[i] = make([]int,columns)
		for j := 0;j < columns;j++ {
			var top_left_region int = 0
			var left_region int = 0
			var top_region int = 0
			var top int = i - K
			var left int = j - K
			if top - 1 >= 0 && left - 1 >= 0{
				top_region = 0
				top_left_region = sum[top - 1][left - 1]
			}
			var bottom int = i + K
			var right int = j + K
			if bottom >= rows{
				bottom = rows - 1
			}
			if right >= columns{
				right = columns - 1
			}
			if top - 1 >= 0{
				top_region = sum[top - 1][right]
			}
			if left - 1 >= 0{
				left_region = sum[bottom][left - 1]
			}
			res[i][j] = sum[bottom][right] - left_region - top_region + top_left_region
		}
	}
	return res
}