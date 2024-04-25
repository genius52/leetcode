package array

//Equal to the cell below it, i.e. grid[i][j] == grid[i + 1][j] (if it exists).
//Different from the cell to its right, i.e. grid[i][j] != grid[i][j + 1] (if it exists).

func dp_minimumOperations3122(columns int, rows int, col int, pre_col_val int, col_cnt [][10]int, memo [][10]int) int {
	if col >= columns {
		return 0
	}
	if pre_col_val != -1 {
		if memo[col][pre_col_val] != -1 {
			return memo[col][pre_col_val]
		}
	}
	var res int = 1<<31 - 1
	for i := 0; i < 10; i++ {
		if pre_col_val == i {
			continue
		}
		ret := rows - col_cnt[col][i] + dp_minimumOperations3122(columns, rows, col+1, i, col_cnt, memo)
		if ret < res {
			res = ret
		}
	}
	if col != 0 {
		memo[col][pre_col_val] = res
	}
	return res
}

func minimumOperations3122(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var col_cnt [][10]int = make([][10]int, columns)
	for c := 0; c < columns; c++ {
		for r := 0; r < rows; r++ {
			col_cnt[c][grid[r][c]]++
		}
	}
	var memo [][10]int = make([][10]int, columns)
	for i := 0; i < columns; i++ {
		for j := 0; j < 10; j++ {
			memo[i][j] = -1
		}
	}
	return dp_minimumOperations3122(columns, rows, 0, -1, col_cnt, memo)
}
