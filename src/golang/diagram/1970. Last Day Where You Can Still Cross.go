package diagram

import "container/list"

//Input: row = 2, col = 2, cells = [[1,1],[2,1],[1,2],[2,2]]
//Output: 2
func bfs_latestDayToCross(row int,col int,cells [][]int,days int)bool{
	var graph [][]int = make([][]int,row)
	for i := 0;i < row;i++{
		graph[i] = make([]int,col)
	}
	for i := 0;i < days;i++{
		graph[cells[i][0] - 1][cells[i][1] - 1] = 1
	}
	var q list.List
	for i := 0;i < col;i++{
		if graph[0][i] != 1{
			var p point
			p.x = 0
			p.y = i
			q.PushBack(p)
			graph[0][i] = 1
		}
	}
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			cur := q.Front().Value.(point)
			q.Remove(q.Front())
			if cur.x == (row - 1){
				return true
			}
			for _,dir := range dirs{
				var next point
				next.x = cur.x + dir[0]
				next.y = cur.y + dir[1]
				if next.x >= 0 && next.x < row && next.y >= 0 && next.y < col && graph[next.x][next.y] != 1{
					q.PushBack(next)
					graph[next.x][next.y] = 1
				}
			}
		}
	}
	return false
}

func LatestDayToCross(row int, col int, cells [][]int) int {
	var left int = 0
	var right int = len(cells)
	var res int = 0
	for left <= right{
		mid := left + (right - left)/2
		if bfs_latestDayToCross(row,col,cells,mid){
			res = mid
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return res
}