package array

import "sort"

func deleteGreatestValue(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	for i := 0;i < rows;i++{
		sort.Ints(grid[i])
	}
	var res int = 0
	for j := columns - 1;j >= 0;j--{
		var cur int = grid[0][j]
		for i := 1;i < rows;i++{
			if grid[i][j] > cur{
				cur = grid[i][j]
			}
		}
		res += cur
	}
	return res
}