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
//	var dp_left [][]int = make([][]int,rows)
//	var dp_right [][]int = make([][]int,rows)
//	for i := 0;i < rows;i++{
//		dp_left[i] = make([]int,columns)
//		dp_right[i] = make([]int,columns)
//	}
//	dp_left[0][0] = grid[0][0]
//	dp_right[0][columns - 1] = grid[0][columns - 1]
//	for i := 0;i < rows - 1;i++{
//		for j := 0;j < columns;j++{
//			if dp_left[i][j] == 0{
//				break
//			}
//			dp_left[i + 1][j] = max_int(dp_left[i + 1][j],dp_left[i][j] + grid[i + 1][j])
//			if j < columns - 1{
//				dp_left[i + 1][j + 1] = max_int(dp_left[i][j + 1],dp_left[i][j] + grid[i + 1][j + 1])
//			}
//			if j > 0{
//				dp_left[i + 1][j - 1] = max_int(dp_left[i][j - 1],dp_left[i][j] + grid[i + 1][j - 1])
//			}
//		}
//	}
//	for i := 0;i < rows - 1;i++{
//		for j := columns - 1;j >= 0;j--{
//			if dp_right[i][j] == 0{
//				break
//			}
//			dp_right[i + 1][j] = max_int(dp_right[i + 1][j],dp_right[i][j] + grid[i + 1][j])
//			if j < columns - 1{
//				dp_right[i + 1][j + 1] = max_int(dp_right[i][j + 1],dp_right[i][j] + grid[i + 1][j + 1])
//			}
//			if j > 0{
//				dp_right[i + 1][j - 1] = max_int(dp_right[i][j - 1],dp_right[i][j] + grid[i + 1][j - 1])
//			}
//		}
//	}
//	var dp_total [][][2]int = make([][][2]int,rows)
//	for i := 0;i < rows;i++{
//		dp_total[i] = make([][2]int,columns)
//	}
//	for i := 0;i < rows - 1;i++{
//		//var max_sum int = 0
//		for j := 0;j < columns;j++{
//			//max_sum = max_int(max_sum,)
//		}
//	}
//	return 0
//}