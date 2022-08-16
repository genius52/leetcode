package array

func LargestLocal(grid [][]int) [][]int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res [][]int = make([][]int,rows - 2)
	for i := 0;i < rows - 2;i++{
		res[i] = make([]int,columns - 2)
	}
	var dirs [][2]int = [][2]int{{-1,-1},{-1,0},{-1,1},{0,1},{0,-1},{1,-1},{1,0},{1,1},{0,0}}
	for i := 0;i < rows - 2 ;i++{
		for j := 0;j < columns - 2;j++{
			for _,dir := range dirs{
				var next_r int = i + 1 + dir[0]
				var next_c int = j + 1 + dir[1]
				if grid[next_r][next_c] > res[i][j]{
					res[i][j] = grid[next_r][next_c]
				}
			}
		}
	}
	return res
}