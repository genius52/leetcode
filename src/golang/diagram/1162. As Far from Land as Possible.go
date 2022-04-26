package diagram

import "container/list"

//BFS solution
func MaxDistance2(grid [][]int) int{
	var rows int = len(grid)
	var columns int = len(grid[0])
	var q list.List
	var zero_cnt int = rows * columns
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 1{
				var p point = point{i,j}
				q.PushBack(p)
				zero_cnt--
			}
		}
	}
	if q.Len() == 0 || q.Len() == rows * columns{
		return -1
	}
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	var distance int = 1
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			cur := q.Front().Value.(point)
			q.Remove(q.Front())
			for _,dir := range dirs{
				var next point
				next.x = cur.x + dir[0]
				next.y = cur.y + dir[1]
				if next.x < 0 || next.y < 0 || next.x >= rows || next.y >= columns{
					continue
				}
				if grid[next.x][next.y] != 0{
					continue
				}
				q.PushBack(next)
				grid[next.x][next.y] = distance
				zero_cnt--
			}
		}
		if zero_cnt == 0{
			break
		}
		distance++
	}
	return distance
}

//DFS solution
func search_neighbour(p point,grid [][]int,rows int,columns int,q *list.List){
	if p.x - 1 >= 0{
		if grid[p.x - 1][p.y] == 0{
			q.PushBack(point{p.x - 1,p.y})
			grid[p.x - 1][p.y] = 1
		}
	}
	if p.y - 1 >= 0 && grid[p.x][p.y - 1] == 0{
		q.PushBack(point{p.x,p.y - 1})
		grid[p.x][p.y - 1] = 1
	}
	if p.y + 1 < columns && grid[p.x][p.y + 1] == 0{
		q.PushBack(point{p.x,p.y + 1})
		grid[p.x][p.y + 1] = 1
	}
	if p.x + 1 < rows{
		if grid[p.x + 1][p.y] == 0{
			q.PushBack(point{p.x + 1,p.y})
			grid[p.x + 1][p.y] = 1
		}
	}
}

func MaxDistance(grid [][]int) int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var q list.List
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if grid[i][j] == 1{
				var p point = point{i,j}
				q.PushBack(p)
			}
		}
	}
	if q.Len() == 0 || q.Len() == rows * columns{
		return -1
	}
	var steps int = 0
	for q.Len() > 0{
		var q_len int = q.Len()
		for i := 0;i < q_len;i++{
			var p point = q.Front().Value.(point)
			q.Remove(q.Front())
			search_neighbour(p,grid,rows,columns,&q)
		}
		if q.Len() == 0{
			return steps
		}
		steps++
	}
	return steps
}