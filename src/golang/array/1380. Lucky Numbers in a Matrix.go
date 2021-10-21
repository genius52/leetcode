package array

func luckyNumbers (matrix [][]int) []int{
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var min_row []int = make([]int,rows)
	for i := 0;i < rows;i++{
		min_row[i] = 2147483647
	}
	var max_column []int = make([]int,columns)
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			max_column[j] = max_int(max_column[j],matrix[i][j])
			min_row[i] = min_int(min_row[i],matrix[i][j])
		}
	}
	var res []int
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if min_row[i] == max_column[j]{
				res = append(res,min_row[i])
				break
			}
		}
	}
	return res
}

//func luckyNumbers (matrix [][]int) []int {
//	var rows int = len(matrix)
//	var columns int = len(matrix[0])
//	var res []int
//	for i := 0;i < rows;i++{
//		var min_num_column int = 0
//		var cur_row_min int = matrix[i][0]
//		for j := 0;j < columns;j++{
//			if matrix[i][j] < cur_row_min{
//				cur_row_min = matrix[i][j]
//				min_num_column = j
//			}
//		}
//
//		var pass bool = true
//		for j := 0;j < rows;j++{
//			if j == i{
//				continue
//			}
//			if matrix[j][min_num_column] > matrix[i][min_num_column]{
//				pass = false
//				break
//			}
//		}
//		if pass{
//			res = append(res, matrix[i][min_num_column])
//		}
//	}
//	return res
//}