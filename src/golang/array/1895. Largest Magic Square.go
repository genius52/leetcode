package array

func check_largestMagicSquare(grid [][]int,rows int,columns,target_len int,
	prefix_row [][]int,prefix_col [][]int)bool{
	for i := 0;(i + target_len) <= rows;i++{
		for j := 0;(j + target_len) <= columns;j++{
			var sum int = prefix_row[i][j + target_len] - prefix_row[i][j]
			var equal bool = true
			for x := i + 1;x <(i + target_len);x++{
				if (prefix_row[x][j + target_len] - prefix_row[x][j]) != sum{
					equal = false
					break
				}
			}
			if !equal{
				continue
			}
			for y := j;y < (j + target_len);y++{
				if (prefix_col[i + target_len][y] - prefix_col[i][y]) != sum{
					equal = false
					break
				}
			}
			if !equal{
				continue
			}
			var sum_tilt1 int = 0
			for x,y := i, j;x < (i + target_len);x++{
				sum_tilt1 += grid[x][y]
				y++
			}
			if sum_tilt1 != sum{
				continue
			}
			var sum_tilt2 int = 0
			for x,y := i,j + target_len - 1;x < (i + target_len);x++{
				sum_tilt2 += grid[x][y]
				y--
			}
			if sum_tilt2 != sum{
				continue
			}
			return true
		}
	}
	return false
}

func LargestMagicSquare(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var prefix_row [][]int = make([][]int,rows + 1)
	var prefix_col [][]int = make([][]int,rows + 1)
	for i := 0;i <= rows;i++{
		prefix_row[i] = make([]int,columns + 1)
	}
	for i := 0;i <= rows;i++{
		prefix_col[i] = make([]int,columns + 1)
	}
	//x axis
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			prefix_row[i][j + 1] = prefix_row[i][j] + grid[i][j]
		}
	}
	//y axis
	for j := 0;j < columns;j++{
		for i := 0;i < rows;i++{
			prefix_col[i + 1][j] = prefix_col[i][j] + grid[i][j]
		}
	}
	var limit int = min_int(rows,columns)
	for i := limit;i > 1;i--{
		if check_largestMagicSquare(grid,rows,columns,i,prefix_row,prefix_col){
			return i
		}
	}
	return 1
}