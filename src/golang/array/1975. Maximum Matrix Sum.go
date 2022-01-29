package array

func maxMatrixSum(matrix [][]int) int64 {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var neg_cnt int = 0
	var min_val int64 = 2147483647
	var sum int64 = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			abs := abs_int(matrix[i][j])
			sum += int64(abs)
			if matrix[i][j] < 0 {
				neg_cnt++
			}
			if int64(abs) < min_val{
				min_val = int64(abs)
			}
		}
	}
	if neg_cnt | 1 == neg_cnt{ //odd cnt
		return sum - min_val * 2
	}else{
		return sum
	}
}