package array

func dp_cherryPickup2(grid [][]int,rows int,columns int,r int,c1 int,c2 int,memo *[][][]int)int{
	if r >= rows{
		return 0
	}
	if (*memo)[r][c1][c2] != -1{
		return (*memo)[r][c1][c2]
	}
	var res int = 0
	for dir_left := -1;dir_left <= 1;dir_left++{
		for dir_right := -1;dir_right <= 1;dir_right++{
			next_c1 := c1 + dir_left
			next_c2 := c2 + dir_right
			if next_c1 >= 0 && next_c1 < columns && next_c2 >= 0 && next_c2 < columns{
				res = max_int(res,dp_cherryPickup2(grid,rows,columns,r + 1,next_c1 ,next_c2,memo))
			}
		}
	}
	if c1 == c2{
		res += grid[r][c1]
	}else{
		res += grid[r][c1] + grid[r][c2]
	}
	(*memo)[r][c1][c2] = res
	return res
}

func CherryPickup2(grid [][]int) int{
	var rows int = len(grid)
	var columns int = len(grid[0])
	var memo [][][]int = make([][][]int,rows)
	for i := 0;i < rows;i++{
		memo[i] = make([][]int,columns)
		for j := 0;j < columns;j++{
			memo[i][j] = make([]int,columns)
			for k := 0;k < columns;k++{
				memo[i][j][k] = -1
			}
		}
	}
	return dp_cherryPickup2(grid,rows,columns,0,0,columns - 1,&memo)
}

//func cherryPickup(grid [][]int) int {
//	var rows int = len(grid)
//	var columns int = len(grid[0])
//	var dp [][][]int = make([][][]int,rows)
//	for i := 0;i < rows;i++{
//		dp[i] = make([][]int,columns)
//		for j := 0;j < columns;j++{
//			dp[i][j] = make([]int,columns)
//		}
//	}
//	dp[0][0][columns - 1] = grid[0][0] + grid[0][columns]
//	for i := 1;i < rows;i++{
//		for left := 0;left < columns;left++{
//			for right := columns - 1;right >= 0;right--{
//				for dir1 := -1;dir1 <= 1;dir1++{
//					for dir2 := -1;dir2 <= 1;dir2++{
//						last_c1 := left + dir1
//						last_c2 := right + dir2
//						if last_c1 >= 0 && last_c1 < columns && last_c2 >= 0 && last_c2 < columns{
//							if left == right{
//								dp[i][left][right] = max_int(dp[i][left][right],grid[i][left] + dp[i - 1][])
//							}
//							 = 0
//						}
//					}
//				}
//			}
//
//		}
//	}
//	return 0
//}