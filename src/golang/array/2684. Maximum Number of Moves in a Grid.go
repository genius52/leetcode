package array

import "container/list"

func MaxMoves(grid [][]int) int {
	type point struct {
		x int
		y int
	}
	var rows int = len(grid)
	var columns int = len(grid[0])
	var visited [][]bool = make([][]bool, rows)
	var q list.List
	for i := 0; i < rows; i++ {
		var p point
		p.x = i
		p.y = 0
		q.PushBack(p)
		visited[i] = make([]bool, columns)
	}
	var dirs [][]int = [][]int{{-1, 1}, {0, 1}, {1, 1}}
	var steps int = 0
	for q.Len() > 0 {
		var cur_len int = q.Len()
		for i := 0; i < cur_len; i++ {
			var cur point = q.Front().Value.(point)
			q.Remove(q.Front())
			for _, dir := range dirs {
				var next point
				next.x = cur.x + dir[0]
				next.y = cur.y + dir[1]
				if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= columns || visited[next.x][next.y] ||
					grid[cur.x][cur.y] >= grid[next.x][next.y] {
					continue
				}
				visited[next.x][next.y] = true
				q.PushBack(next)
			}
		}
		if q.Len() == 0 {
			break
		}
		steps++
	}
	return steps
}
