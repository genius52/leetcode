package array

//Input: mat = [[1,0,0],
//              [0,0,1],
//              [1,0,0]]
//Output: 1
func numSpecial(mat [][]int) int{
	rows := len(mat)
	columns := len(mat[0])
	var existed_row []int = make([]int,rows)
	var existed_col []int = make([]int,columns)
	var res int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if mat[i][j] == 1{
				existed_row[i]++
				existed_col[j]++
			}
		}
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if mat[i][j] == 1{
				if existed_row[i] == 1 && existed_col[j] == 1{
					res++
				}
			}
		}
	}
	return res
}

func NumSpecial(mat [][]int) int {
	var res int = 0
	var rows int = len(mat)
	var columns int = len(mat[0])
	var special_rows []int
	var special_columns []int
	for r := 0;r < rows;r++{
		var one_cnt int = 0
		for c := 0;c < columns;c++{
			if mat[r][c] == 1{
				one_cnt++
				if one_cnt > 1{
					break
				}
			}
		}
		if one_cnt == 1{
			special_rows = append(special_rows,r)
		}
	}
	for c := 0;c < columns;c++{
		var one_cnt int = 0
		for r := 0;r < rows;r++{
			if mat[r][c] == 1{
				one_cnt++
				if one_cnt > 1{
					break
				}
			}
		}
		if one_cnt == 1{
			special_columns = append(special_columns,c)
		}
	}
	for _,r := range special_rows{
		for _,c := range special_columns{
			if mat[r][c] == 1{
				res++
			}
		}
	}
	return res
}