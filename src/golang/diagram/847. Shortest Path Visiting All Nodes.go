package diagram

import "container/list"

type State struct {
	cur_node int
	trace int
	dis int
}

func ShortestPathLength(graph [][]int) int {
	var l int = len(graph)
	var visited [][]bool = make([][]bool,l)//memo[i][j]: current is i,state is j
	//var state int = 1 << l
	for i := 0;i < l;i++{
		visited[i] = make([]bool, 1 << l)
	}
	var q list.List
	for i := 0;i < l;i++{
		q.PushBack(State{i,1 << i,0})
		visited[i][1 << i] = true
	}
	var end int = 1 << l - 1
	//var res int = 2147483647
	for q.Len() > 0{
		cur := q.Front().Value.(State)
		q.Remove(q.Front())
		if cur.trace == end{
			return cur.dis
			//res = min_int(res,cur.dis)
		}
		for _,neighbour := range graph[cur.cur_node]{
			var next State
			next.cur_node = neighbour
			next.trace = cur.trace | (1 << neighbour)
			next.dis = cur.dis + 1
			if visited[neighbour][next.trace] {
				continue
			}
			visited[neighbour][next.trace] = true
			q.PushBack(next)
		}
	}
	return 0
}