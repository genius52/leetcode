package diagram

//Input: grid = [[1,2],[3,4]]
//Output: 34
func surfaceArea(grid [][]int) int {
	var rows int = len(grid[0])
	var columns int = len(grid[0])
	var total int = 0
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 0{
				continue
			}
			total += 2 + grid[i][j] * 4
			if i == 0 && j == 0{
				continue
			}
			var dup_cnt int = 0
			if i > 0{
				dup_cnt += min_int(grid[i][j],grid[i - 1][j])
			}
			if j > 0{
				dup_cnt += min_int(grid[i][j],grid[i][j - 1])
			}
			total -= dup_cnt * 2
		}
	}
	return total
}