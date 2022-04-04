package number

import "math"

//[
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
type point struct {
	r int
	c int
}
func dp_minimumTotal(row int,column int,rows int,triangle [][]int,record map[point]int)int{
	if row >= rows {
		return 0
	}
	if column < 0 || column > row{
		return math.MaxInt32
	}
	var p point
	p.r = row
	p.c = column
	if _,ok := record[p];ok{
		return record[p]
	}
	res := triangle[row][column] + min_int_number(dp_minimumTotal(row + 1,column,rows,triangle,record),dp_minimumTotal(row + 1,column + 1,rows,triangle,record))
	record[p] = res
	return res
}

func MinimumTotal(triangle [][]int) int {
	var record map[point]int = make(map[point]int)
	rows := len(triangle)
	res := dp_minimumTotal(0,0,rows,triangle,record)
	return res
}

//DP solution from bottom to top
func MinimumTotal2(triangle [][]int) int {
	rows := len(triangle)
	var dp [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		dp[i] = make([]int,i + 1)
	}
	for i := 0;i < rows;i++{
		dp[rows-1][i] = triangle[rows-1][i]
	}
	for i := rows - 2;i >= 0;i--{
		for j := 0;j <= i;j++{
			dp[i][j] = triangle[i][j] + min_int(dp[i + 1][j],dp[i + 1][j + 1])
		}
	}
	return dp[0][0]
}

func MinimumTotal3(triangle [][]int) int{
	var rows int = len(triangle)
	if rows == 1{
		return triangle[0][0]
	}
	var res int = 2147483647
	var pre []int = []int{triangle[0][0]}
	for i := 1;i < rows;i++{
		var cur []int = make([]int,i + 1)
		var columns int = len(triangle[i])
		for j := 0;j < columns;j++{
			if j == 0{
				cur[j] = triangle[i][j] + pre[j]
			}else if j == columns - 1{
				cur[j] = triangle[i][j] + pre[j - 1]
			}else{
				cur[j] = min_int(triangle[i][j] + pre[j - 1],triangle[i][j] + pre[j])
			}
		}
		pre = cur
		if i == rows - 1{
			for j := 0;j < len(triangle[i]);j++{
				res = min_int(res,cur[j])
			}
		}
	}
	return res
}
