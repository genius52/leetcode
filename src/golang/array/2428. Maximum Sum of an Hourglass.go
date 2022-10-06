package array

func maxSum(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res int = 0
	for i := 1;i < rows - 1;i++{
		for j := 1;j < columns - 1;j++{
			sum := grid[i][j] + grid[i - 1][j - 1] + grid[i - 1][j] + grid[i - 1][j + 1] +
				grid[i + 1][j - 1] + grid[i + 1][j] + grid[i + 1][j + 1]
			if sum > res{
				res = sum
			}
		}
	}
	return res
}