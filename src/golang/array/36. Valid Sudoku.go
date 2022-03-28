package array

func isValidSudoku2(board [][]byte) bool{
	var row_record [10][10]int
	var col_record [10][10]int
	var matrix_record [10][10]int
	for i := 0;i < 10;i++{
		for j := 0;j < 10;j++{
			if board[i][j] != '.'{
				val := board[i][j] - '0'
				row_record[i][val]++
				if row_record[i][val] > 1{
					return false
				}
				col_record[j][val]++
				if col_record[j][val] > 1{
					return false
				}
				idx := i /3 * 3 + j / 3
				matrix_record[idx][val]++
				if matrix_record[idx][val] > 1{
					return false
				}
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	var row_map [9][]byte
	var column_map [9][]byte
	var threeplusthree_map [9][]byte
	for i := 0;i < len(board);i++{
		for j := 0;j < len(board[0]);j++{
			if board[i][j] != '.'{
				for _,val := range row_map[i]{
					if val == board[i][j]{
						return false
					}
				}
				row_map[i] = append(row_map[i], board[i][j])
				for _,val := range column_map[j]{
					if val == board[i][j]{
						return false
					}
				}
				column_map[j] = append(column_map[j], board[i][j])
				index := i/3 * 3  + j / 3
				for _,val := range threeplusthree_map[index]{
					if val == board[i][j]{
						return false
					}
				}
				threeplusthree_map[index] = append(threeplusthree_map[index],board[i][j])
			}
		}
	}
	return true
}
