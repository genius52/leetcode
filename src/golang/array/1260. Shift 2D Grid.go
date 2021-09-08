package array

func shiftGrid(grid [][]int, k int) [][]int{
	var rows int = len(grid)
	var columns int = len(grid[0])
	k %= rows * columns
	if k == 0{
		return grid
	}
	var res [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		res[i] = make([]int,columns)
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			next_r := (i + (j + k) / columns) % rows
			next_c := (j + k) % columns
			res[next_r][next_c] = grid[i][j]
		}
	}
	return res
}
// [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]] K = 4
//func shiftGrid(grid [][]int, k int) [][]int {
//	var row_num int = len(grid)
//	var col_num int = len(grid[0])
//	var total = row_num * col_num
//	var res []int = make([]int,total)
//	for i := 0;i < row_num;i++{
//		for j := 0;j < col_num;j++{
//			res[(i * col_num + j + k) % total] = grid[i][j]
//		}
//	}
//	var grid2 [][]int = make([][]int,len(grid))
//	for i := 0;i < row_num;i++{
//		grid2[i] = make([]int,len(grid[0]))
//		for j := 0;j < col_num;j++{
//			grid2[i][j] = res[i * col_num + j]
//		}
//	}
//	return grid2
//}