package array

import "container/list"

//Input: isWater = [[0,0,1],[1,0,0],[0,0,0]]
//Output: [[1,1,0],[0,1,1],[1,2,2]]
func HighestPeak(isWater [][]int) [][]int {
	var rows int = len(isWater)
	var columns int = len(isWater[0])
	var res [][]int = make([][]int,rows)
	var visited [][]bool = make([][]bool,rows)
	for i := 0;i < rows;i++{
		res[i] = make([]int,columns)
		visited[i] = make([]bool,columns)
	}
	var q list.List
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			if isWater[i][j] == 1{
				var p point
				p.x = i
				p.y = j
				q.PushBack(p)
				visited[i][j] = true
			}
		}
	}

	var height int = 0
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var p point = q.Front().Value.(point)
			q.Remove(q.Front())
			res[p.x][p.y] = height
			if p.x - 1 >= 0 && !visited[p.x - 1][p.y]{
				var p1 point
				p1.x = p.x - 1
				p1.y = p.y
				q.PushBack(p1)
				visited[p1.x][p1.y] = true
			}
			if p.x + 1 < rows && !visited[p.x + 1][p.y]{
				var p1 point
				p1.x = p.x + 1
				p1.y = p.y
				q.PushBack(p1)
				visited[p1.x][p1.y] = true
			}
			if p.y - 1 >= 0 && !visited[p.x][p.y - 1]{
				var p1 point
				p1.x = p.x
				p1.y = p.y - 1
				q.PushBack(p1)
				visited[p1.x][p1.y] = true
			}
			if p.y + 1 < columns && !visited[p.x][p.y + 1]{
				var p1 point
				p1.x = p.x
				p1.y = p.y + 1
				q.PushBack(p1)
				visited[p1.x][p1.y] = true
			}
		}
		height++
	}
	return res
}