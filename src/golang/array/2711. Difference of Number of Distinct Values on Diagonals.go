package array

//func differenceOfDistinctValues(grid [][]int) [][]int {
//	var rows int = len(grid)
//	var columns int = len(grid[0])
//	var left_side []map[int]bool = make([]map[int]bool, rows) //左边的部分
//	for i := 0; i < rows; i++ {
//		left_side[i] = make(map[int]bool)
//	}
//	var right_side []map[int]bool = make([]map[int]bool, columns) //右边部分
//	for i := 0; i < columns; i++ {
//		right_side[i] = make(map[int]bool)
//	}
//	var res [][]int = make([][]int, rows)
//	for i := 0; i < rows; i++ {
//		res[i] = make([]int, columns)
//	}
//	for i := 0; i < rows; i++ {
//		for j := 0; j < columns; j++ {
//			if i >= j {
//				left_side[i][grid[i][j]] = true
//			} else {
//				right_side[j][grid[i][j]] = true
//			}
//		}
//	}
//	for i := 0; i < rows; i++ {
//		for j := 0; j < columns; j++ {
//			if i >= j {
//				res[i][j] = len(left_side[rows-1]) - len(left_side[i])
//				if i > 0 {
//					res[i][j] = abs_int(res[i][j] - len(left_side[i-1]))
//				}
//			} else {
//				//res[i][j] = len()
//			}
//		}
//	}
//	return res
//}

func differenceOfDistinctValues(grid [][]int) [][]int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res [][]int = make([][]int, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, columns)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			var left_up map[int]bool = make(map[int]bool)
			var right_bottom map[int]bool = make(map[int]bool)
			var r int = i - 1
			var c int = j - 1
			for r >= 0 && c >= 0 {
				left_up[grid[r][c]] = true
				r--
				c--
			}
			r = i + 1
			c = j + 1
			for r < rows && c < columns {
				right_bottom[grid[r][c]] = true
				r++
				c++
			}
			res[i][j] = abs_int(len(right_bottom) - len(left_up))
		}
	}
	return res
}
