package array

func minFlips3240(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res int = 0
	for r := 0; r < rows/2; r++ {
		for c := 0; c < columns/2; c++ {
			var zero_cnt int = 0
			if grid[r][c] == 1 {
				zero_cnt++
			}
			if grid[r][columns-1-c] == 1 {
				zero_cnt++
			}
			if grid[rows-1-r][c] == 1 {
				zero_cnt++
			}
			if grid[rows-1-r][columns-1-c] == 1 {
				zero_cnt++
			}
			res += min(zero_cnt, 4-zero_cnt)
			//dp_cnt[r][c][0] = zero_cnt
			//dp_cnt[r][c][1] = 4 - zero_cnt
		}
	}
	var cnt_00 int = 0
	var cnt_11 int = 0
	var cnt_01 int = 0
	if rows|1 == rows {
		for c := 0; c < columns/2; c++ {
			sum := grid[rows/2][c] + grid[rows/2][columns-1-c]
			if sum == 0 {
				cnt_00++
			} else if sum == 1 {
				cnt_01++
				res++
			} else if sum == 2 {
				cnt_11++
			}
		}
	}
	if columns|1 == columns {
		for r := 0; r < rows/2; r++ {
			sum := grid[r][columns/2] + grid[rows-1-r][columns/2]
			if sum == 0 {
				cnt_00++
			} else if sum == 1 {
				cnt_01++
				res++
			} else if sum == 2 {
				cnt_11++
			}
		}
	}
	if (rows|1 == rows) && (columns|1 == columns) {
		if grid[rows/2][columns/2] == 1 {
			res++
		}
	}
	if cnt_11%2 == 0 {
		return res
	} else { //缺2个1
		if cnt_01 == 0 {
			return res + 2
		} else {
			return res
		}
	}
}
