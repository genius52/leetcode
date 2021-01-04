package array

//Input: grid = [[1,1,0,0],[0,0,1,0],[0,0,1,0],[0,0,0,1]]
//Output: 4
func countServers(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var rows_cnt []int = make([]int,rows)
	var columns_cnt []int = make([]int,columns)
	for r := 0;r < rows;r++{
		for c := 0;c < columns;c++{
			if grid[r][c] == 1{
				rows_cnt[r]++
				columns_cnt[c]++
			}
		}
	}
	var res int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 1{
				if rows_cnt[i] > 1 || columns_cnt[j] > 1{
					res++
				}
			}
		}
	}
	return res
}