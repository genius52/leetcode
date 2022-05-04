package array

func maximalSquare(matrix [][]byte) int {
	rows := len(matrix)
	columns := len(matrix[0])
	var dp [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		dp[i] = make([]int,columns)
	}
	var res int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if matrix[i][j] == '1'{
				dp[i][j] = 1
				res = 1
			}
		}
	}
	for i := 1;i < len(matrix);i++{
		for j := 1;j < len(matrix[0]);j++{
			if dp[i][j] == 0{
				continue
			}
			dp[i][j] = 1 + min_int_number(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])
			if dp[i][j] > res{
				res = dp[i][j]
			}
		}
	}
	return res * res
}