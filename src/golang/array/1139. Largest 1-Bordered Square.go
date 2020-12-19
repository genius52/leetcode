package array

//Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
//Output: 9
//Example 2:
//
//Input: grid = [[1,1,0,0]]
//Output: 1
func dfs_mark(grid [][]int,r int,c int,rows int,columns int,search_val int,set_val int){
	if r < 0 || r >= rows || c < 0 || c >= columns{
		return
	}
	if grid[r][c] != search_val || grid[r][c] == set_val{
		return
	}
	grid[r][c] = set_val
	dfs_mark(grid,r - 1,c,rows,columns,search_val,set_val)
	dfs_mark(grid,r + 1,c,rows,columns,search_val,set_val)
	dfs_mark(grid,r,c - 1,rows,columns,search_val,set_val)
	dfs_mark(grid,r,c + 1,rows,columns,search_val,set_val)
}

func Largest1BorderedSquare(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var dp_rows [][]int = make([][]int,rows)
	var dp_columns [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		dp_rows[i] = make([]int,columns)
		dp_columns[i] = make([]int,columns)
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 0{
				continue
			}
			if j > 0{
				dp_rows[i][j] = 1 + dp_rows[i][j - 1]
			}else{
				dp_rows[i][j] = 1
			}
			if i > 0{
				dp_columns[i][j] = 1 + dp_columns[i - 1][j]
			}else{
				dp_columns[i][j] = 1
			}
		}
	}
	var res int = 0
	for i := rows - 1;i >= 0;i--{
		if res >= i + 1{
			return res * res
		}
		for j := columns - 1;j >= 0;j--{
			if dp_rows[i][j] == 0 || dp_columns[i][j] == 0{
				continue
			}
			distance := min_int_number(dp_rows[i][j],dp_columns[i][j])
			var cur_max int = 1
			for k := distance;k > 0;k--{
				if dp_rows[i - k + 1][j] >=k && dp_columns[i][j - k + 1] >= k{
					cur_max = k
					break
				}
			}
			if cur_max > res{
				res = cur_max
			}
		}
	}
	return res * res
}