package diagram

import "container/list"


func ShortestPathBinaryMatrix(grid [][]int) int {
	var l int = len(grid)
	if grid[0][0] == 1 || grid[l - 1][l - 1] == 1{
		return -1
	}
	if l == 1{
		return 1
	}
	var visited [][]bool = make([][]bool,l)
	for i := 0;i < l;i++{
		visited[i] = make([]bool,l)
	}
	visited[0][0] = true
	var dir [][]int = [][]int{[]int{-1,-1},[]int{-1,0},[]int{-1,1},[]int{0,-1},[]int{0,1},[]int{1,-1},[]int{1,0},[]int{1,1}}
	var q list.List
	var p point
	p.x = 0
	p.y = 0
	q.PushFront(p)
	var node_cnt int = 1
	for q.Len() > 0{
		var size int = q.Len()
		for i := 0;i < size;i++{
			var node = q.Back().Value.(point)
			q.Remove(q.Back())
			for i := 0;i < 8;i++{
				next_x := node.x + dir[i][0]
				next_y := node.y + dir[i][1]
				if next_x == l - 1 && next_y == l - 1{
					return node_cnt + 1
				}
				if next_x < 0 || next_y < 0 || next_x >= l || next_y >= l || visited[next_x][next_y] || grid[next_x][next_y] == 1{
					continue
				}
				var p2 point
				p2.x = next_x
				p2.y = next_y
				q.PushFront(p2)
				visited[next_x][next_y] = true
			}
		}
		node_cnt++
	}
	return -1
}