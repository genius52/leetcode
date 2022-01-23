package array

import (
	"container/list"
	"sort"
)

type Move struct{
	steps int
	x int
	y int
}

func HighestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	var rows int = len(grid)
	var columns int = len(grid[0])
	var res [][]int
	start_x := start[0]
	start_y := start[1]
	my_sort := func(i, j int) bool {
		dis1 := res[i][0]
		dis2 := res[j][0]
		if dis1 != dis2{
			return dis1 < dis2
		}
		if grid[res[i][1]][res[i][2]] != grid[res[j][1]][res[j][2]] {
			return grid[res[i][1]][res[i][2]] < grid[res[j][1]][res[j][2]]
		}
		if res[i][1] != res[j][1]{
			return res[i][1] < res[j][1]
		}
		return res[i][2] < res[j][2]
	}
	var q list.List
	var first Move
	first.steps = 0
	first.x = start_x
	first.y = start_y
	q.PushBack(first)
	var visited map[int]bool = make(map[int]bool)
	visited[first.x * 100000 + first.y] = true
	var steps int = 0
	var dirs [][]int = [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			var cur Move = q.Front().Value.(Move)
			q.Remove(q.Front())
			if grid[cur.x][cur.y] >= pricing[0] && grid[cur.x][cur.y] <= pricing[1]{
				res = append(res,[]int{steps,cur.x,cur.y})
				if len(res) > k{
					sort.Slice(res,my_sort)
					res = res[:k]
				}
			}
			for _, dir := range dirs{
				var next Move
				next.steps = steps
				next.x = cur.x + dir[0]
				next.y = cur.y + dir[1]
				if next.x >= 0 && next.x < rows && next.y >= 0 && next.y < columns &&
					grid[next.x][next.y] != 0{
					if _,ok := visited[next.x * 100000 + next.y];!ok{
						q.PushBack(next)
						visited[next.x * 100000 + next.y] = true
					}
				}
			}
		}
		if len(res) == k{
			break
		}
		steps++
	}
	sort.Slice(res,my_sort)
	for i := 0;i < len(res);i++{
		res[i] = []int{res[i][1],res[i][2]}
	}
	return res
}