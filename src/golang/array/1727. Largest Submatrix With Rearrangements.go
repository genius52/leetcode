package array

import "sort"

func largestSubmatrix(matrix [][]int) int {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var record [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		record[i] = make([]int,columns)
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if i == 0{
				record[i][j] = matrix[i][j]
			}else{
				if matrix[i][j] == 1{
					record[i][j] = 1 + record[i - 1][j]
				}
			}
		}
	}
	var res int = 0
	for i := 0;i < rows;i++{
		sort.Slice(record[i], func(a, b int) bool {
			return record[i][a] < record[i][b]
		})
		for j := 0;j < columns;j++{
			area := record[i][j] * (columns - j)
			res = max_int(res,area)
		}
	}
	return res
}