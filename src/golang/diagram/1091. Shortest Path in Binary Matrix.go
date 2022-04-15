package diagram

import "container/list"

//func ShortestPathBinaryMatrix(grid [][]int) int {
//	var l int = len(grid)
//	if grid[0][0] == 1 || grid[l - 1][l - 1] == 1{
//		return -1
//	}
//	if l == 1{
//		return 1
//	}
//	var visited [][]bool = make([][]bool,l)
//	for i := 0;i < l;i++{
//		visited[i] = make([]bool,l)
//	}
//	visited[0][0] = true
//	var dir [][]int = [][]int{[]int{-1,-1},[]int{-1,0},[]int{-1,1},[]int{0,-1},[]int{0,1},[]int{1,-1},[]int{1,0},[]int{1,1}}
//	var q list.List
//	var p point
//	p.x = 0
//	p.y = 0
//	q.PushFront(p)
//	var node_cnt int = 1
//	for q.Len() > 0{
//		var size int = q.Len()
//		for i := 0;i < size;i++{
//			var node = q.Back().Value.(point)
//			q.Remove(q.Back())
//			for i := 0;i < 8;i++{
//				next_x := node.x + dir[i][0]
//				next_y := node.y + dir[i][1]
//				if next_x == l - 1 && next_y == l - 1{
//					return node_cnt + 1
//				}
//				if next_x < 0 || next_y < 0 || next_x >= l || next_y >= l || visited[next_x][next_y] || grid[next_x][next_y] == 1{
//					continue
//				}
//				var p2 point
//				p2.x = next_x
//				p2.y = next_y
//				q.PushFront(p2)
//				visited[next_x][next_y] = true
//			}
//		}
//		node_cnt++
//	}
//	return -1
//}

func ShortestPathBinaryMatrix(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	if grid[0][0] == 1 || grid[rows - 1][columns - 1] == 1{
		return -1
	}
	var dirs [][]int = [][]int{[]int{-1,-1},[]int{-1,0},[]int{-1,1},[]int{0,-1},[]int{0,1},[]int{1,-1},[]int{1,0},[]int{1,1}}
	var q list.List
	var p point
	p.x = 0
	p.y = 0
	q.PushFront(p)
	grid[0][0] = 1
	var steps int = 1
	for q.Len() > 0{
		cur_len := q.Len()
		for i := 0;i < cur_len;i++{
			cur := q.Front().Value.(point)
			q.Remove(q.Front())
			if cur.x == rows - 1 && cur.y == columns - 1{
				return steps
			}
			for _,dir := range dirs{
				var next point
				next.x = cur.x + dir[0]
				next.y = cur.y + dir[1]
				if next.x >= 0 && next.x < rows && next.y >= 0 && next.y < columns && grid[next.x][next.y] == 0{
					grid[next.x][next.y] = 1
					q.PushBack(next)
				}
			}
		}
		steps++
	}
	return -1
}