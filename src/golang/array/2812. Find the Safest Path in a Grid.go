package array

import "container/list"

func maximumSafenessFactor(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	if grid[0][0] == 1 || grid[rows-1][columns-1] == 1 {
		return 0
	}
	var q list.List
	type point struct {
		x int
		y int
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == 1 {
				q.PushBack(point{i, j})
			}
		}
	}
	var dirs [][2]int = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var steps int = 2
	for q.Len() > 0 {
		var cur_len int = q.Len()
		for i := 0; i < cur_len; i++ {
			var cur point = q.Front().Value.(point)
			q.Remove(q.Front())
			for _, dir := range dirs {
				var next point
				next.x = cur.x + dir[0]
				next.y = cur.y + dir[1]
				if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= columns || grid[next.x][next.y] != 0 {
					continue
				}
				grid[next.x][next.y] = steps
			}
		}
		steps++
	}
	var q2 list.List
	q2.PushBack(point{0, 0})
	var max_factor int = 0
	for q2.Len() > 0 {
		var cur_len int = q2.Len()
		for i := 0; i < cur_len; i++ {
			var cur point = q2.Front().Value.(point)

		}
	}
	return max_factor
}
