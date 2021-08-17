package array

import "strconv"

func maxEqualRowsAfterFlips(matrix [][]int) int {
	var res int = 0
	var record map[string]int = make(map[string]int)
	for row := 0; row < len(matrix);row++{
		var s string = ""
		var reverse_s string = ""
		for column := 0; column < len(matrix[0]);column++{
			if matrix[row][column] == 1{
				reverse_s += "0"
			}else{
				reverse_s += "1"
			}
			s += strconv.Itoa(matrix[row][column])
		}
		if _,ok := record[reverse_s];ok{
			record[reverse_s]++
		}else{
			record[reverse_s] = 1
		}
		if _,ok := record[s];ok{
			record[s]++
		}else{
			record[s] = 1
		}
		if record[s] > res{
			res = record[s]
		}
	}
	return res
}