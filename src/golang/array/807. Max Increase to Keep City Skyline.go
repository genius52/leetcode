package array

func maxIncreaseKeepingSkyline(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var maxheight_on_rows []int = make([]int,rows)
	var maxheight_on_column []int = make([]int,columns)
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] > maxheight_on_rows[i]{
				maxheight_on_rows[i] = grid[i][j]
			}
			if grid[i][j] > maxheight_on_column[j]{
				maxheight_on_column[j] = grid[i][j]
			}
		}
	}
	var res int = 0
	for i := 0;i < rows;i++ {
		for j := 0; j < columns; j++ {
			res += min_int(maxheight_on_rows[i],maxheight_on_column[j]) - grid[i][j]
		}
	}
	return res
}