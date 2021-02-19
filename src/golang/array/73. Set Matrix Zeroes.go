package array


//73
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//Output:
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]

func setZeroes(matrix [][]int)  {
	if len(matrix) == 0{
		return
	}
	if len(matrix[0]) == 0{
		return
	}
	var first_row_is_zero bool = false
	var first_col_is_zero bool = false
	for i := 0;i < len(matrix);i++{
		if matrix[i][0] == 0{
			first_col_is_zero = true
			break
		}
	}
	for i := 0;i < len(matrix[0]);i++{
		if matrix[0][i] == 0{
			first_row_is_zero = true
			break
		}
	}
	for i := 1;i < len(matrix);i++{
		for j := 1;j < len(matrix[0]);j++{
			if matrix[i][j] == 0{
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1;i < len(matrix);i++{
		for j := 1;j < len(matrix[0]);j++{
			if matrix[0][j] == 0 || matrix[i][0] == 0{
				matrix[i][j] = 0
			}
		}
	}
	if first_row_is_zero{
		for i := 0;i < len(matrix[0]);i++{
			matrix[0][i] = 0
		}
	}
	if first_col_is_zero{
		for i := 0;i < len(matrix);i++{
			matrix[i][0] = 0
		}
	}
}

func setZeroes2(matrix [][]int){
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var zero_rows map[int]bool = make(map[int]bool)
	var zero_columns map[int]bool = make(map[int]bool)
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if matrix[i][j] == 0{
				zero_rows[i] = true
				zero_columns[j] = true
			}
		}
	}
	for r,_ := range zero_rows{
		for j := 0;j < columns;j++{
			matrix[r][j] = 0
		}
	}
	for c,_ := range zero_columns{
		for i := 0;i < rows;i++{
			matrix[i][c] = 0
		}
	}
}