package diagram

import "container/list"

func findSafeWalk(grid [][]int, health int) bool {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var cost [][]int = make([][]int, rows)
	for r := 0; r < rows; r++ {
		cost[r] = make([]int, columns)
	}
	var q list.List
	
}
