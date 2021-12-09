package array

import "container/list"

//Input: isWater = [[0,0,1],[1,0,0],[0,0,0]]
//Output: [[1,1,0],[0,1,1],[1,2,2]]
func HighestPeak(isWater [][]int) [][]int {
	var rows int = len(isWater)
	var columns int = len(isWater[0])
	var res [][]int = make([][]int, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			res[i][j] = -1
		}
	}
	var q list.List
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if isWater[i][j] == 1 {
				var p point
				p.x = i
				p.y = j
				q.PushBack(p)
				res[i][j] = 0
			}
		}
	}
	var dirs [][]int = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var height int = 1
	for q.Len() > 0 {
		var l int = q.Len()
		for i := 0; i < l; i++ {
			var cur point = q.Front().Value.(point)
			q.Remove(q.Front())
			for _, dir := range dirs {
				next := cur
				next.x += dir[0]
				next.y += dir[1]
				if next.x >= 0 && next.x < rows && next.y >= 0 && next.y < columns {
					if res[next.x][next.y] == -1 {
						res[next.x][next.y] = height
						q.PushBack(next)
					}
				}
			}
		}
		height++
	}
	return res
}
