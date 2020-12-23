package diagram

import "container/list"

//Input: grid = [[1,0,0],[0,0,0],[0,0,0]]
//Output: 4
//type point struct {
//	x int
//	y int
//}

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