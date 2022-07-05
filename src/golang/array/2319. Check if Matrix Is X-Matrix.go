package array

func checkXMatrix(grid [][]int) bool {
	var rows int = len(grid)
	for i := 0;i < rows;i++{
		for j := 0;j < rows;j++{
			if i == j || i + j == rows - 1{
				if grid[i][j] == 0{
					return false
				}
			}else{
				if grid[i][j] != 0{
					return false
				}
			}
		}
	}
	return true
}