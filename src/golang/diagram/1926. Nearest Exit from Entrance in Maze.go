package diagram

import "container/list"

func NearestExit(maze [][]byte, entrance []int) int {
	var rows int = len(maze)
	var columns int = len(maze[0])
	var q list.List
	for i := 0;i < rows;i++{
		if maze[i][0] == '.'{
			var p point
			p.x = i
			p.y = 0
			if p.x != entrance[0] || p.y != entrance[1]{
				q.PushBack(p)
			}
		}
		if maze[i][columns - 1] == '.'{
			var p point
			p.x = i
			p.y = columns - 1
			if p.x != entrance[0] || p.y != entrance[1]{
				q.PushBack(p)
			}
		}
	}
	for j := 1;j < columns - 1;j++{
		if maze[0][j] == '.'{
			var p point
			p.x = 0
			p.y = j
			if p.x != entrance[0] || p.y != entrance[1]{
				q.PushBack(p)
			}
		}
		if maze[rows - 1][j] == '.'{
			var p point
			p.x = rows - 1
			p.y = j
			if p.x != entrance[0] || p.y != entrance[1]{
				q.PushBack(p)
			}
		}
	}
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	var steps int = 0
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			var cur point = q.Front().Value.(point)
			q.Remove(q.Front())
			if cur.x == entrance[0] && cur.y == entrance[1]{
				return steps
			}
			for _,d := range dirs{
				var next point
				next.x = cur.x + d[0]
				next.y = cur.y + d[1]
				if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= columns ||
					maze[next.x][next.y] == '+'{
					continue
				}
				maze[next.x][next.y] = '+'
				q.PushBack(next)
			}
		}
		steps++
	}
	return -1
}