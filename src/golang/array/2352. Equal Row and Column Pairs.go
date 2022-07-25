package array

func equalPairs(grid [][]int) int {
	var rows int = len(grid)
	var res int = 0
	for r := 0;r < rows;r++{
		for c := 0;c < rows;c++{
			var equal bool = true
			for i := 0;i < rows;i++{
				if grid[r][i] != grid[i][c]{
					equal = false
					break
				}
			}
			if equal{
				res++
			}
		}
	}
	return res
}