package array

type NumMatrix struct {
	dp [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	var obj NumMatrix
	rows := len(matrix)
	if rows == 0{
		return obj
	}
	columns := len(matrix[0])
	obj.dp = make([][]int,rows)
	for r := 0;r < rows;r++{
		obj.dp[r] = make([]int,columns)
		for c := 0;c < columns;c++{
			if r == 0 && c == 0{
				obj.dp[r][c] = matrix[r][c]
			}else if r == 0 && c > 0{
				obj.dp[r][c] = matrix[r][c] + obj.dp[r][c - 1]
			}else if r > 0 && c == 0{
				obj.dp[r][c] = matrix[r][c] + obj.dp[r - 1][c]
			}else{
				obj.dp[r][c] = matrix[r][c] + obj.dp[r][c - 1] + obj.dp[r - 1][c] - obj.dp[r - 1][c - 1]
			}
		}
	}
	return obj
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var left_arr int = 0
	var upper_arr int = 0
	var left_upper_arr int = 0
	if row1 > 0{
		upper_arr = this.dp[row1 - 1][col2]
	}
	if col1 > 0{
		left_arr = this.dp[row2][col1 - 1]
	}
	if row1 > 0 && col1 > 0{
		left_upper_arr = this.dp[row1 - 1][col1 - 1]
	}
	return this.dp[row2][col2] - upper_arr - left_arr + left_upper_arr
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
