package array

//Input:
//[
//  ["1","0","1","0","0"],
//  ["1","0","1","1","1"],
//  ["1","1","1","1","1"],
//  ["1","0","0","1","0"]
//]
//Output: 6
func MaximalRectangle(matrix [][]byte) int {
	var rows int = len(matrix)
	if rows == 0{
		return 0
	}
	var columns int = len(matrix[0])
	var record [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		record[i] = make([]int,columns)
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if i == 0{
				if matrix[i][j] == '1'{
					record[i][j] = 1
				}
			}else{
				if matrix[i][j] == '1'{
					record[i][j] = 1 + record[i - 1][j]
				}
			}
		}
	}
	var res int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if record[i][j] == 0{
				continue
			}
			var left int = j
			var right int = left
			var min_height int = rows
			for right < columns && record[i][right] != 0{
				min_height = min_int(min_height,record[i][right])
				area := min_height * (right - left + 1)
				res = max_int(res,area)
				right++
			}
		}

	}
	return res
}