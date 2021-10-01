package diagram

import "container/list"

type Mytrace struct {
	node int
	rate float64
	visited []bool
}

func FrogPosition(n int, edges [][]int, t int, target int) float64 {
	var graph [][]int = make([][]int,n + 1)
	for _,edge := range edges{
		graph[edge[0]] = append(graph[edge[0]],edge[1])
		graph[edge[1]] = append(graph[edge[1]],edge[0])
	}
	var first Mytrace
	first.node = 1
	first.visited = make([]bool,n + 1)
	first.visited[1] = true
	first.rate = 1.0
	var res float64 = 0
	var q list.List
	q.PushBack(first)
	for q.Len() > 0 && t >= 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			cur := q.Front().Value.(Mytrace)
			q.Remove(q.Front())
			if t == 0 && cur.node == target{
				res += cur.rate
			}
			var cnt int = 0
			for j := 0;j < len(graph[cur.node]);j++{
				if cur.visited[graph[cur.node][j]]{
					continue
				}
				cnt++
			}
			if cnt > 0{
				var next_rate float64 = cur.rate / float64(cnt)
				for j := 0;j < len(graph[cur.node]);j++{
					if cur.visited[graph[cur.node][j]]{
						continue
					}
					var next Mytrace
					next.node = graph[cur.node][j]
					next.rate = next_rate
					next.visited = make([]bool,n + 1)
					copy(next.visited,cur.visited)
					next.visited[graph[cur.node][j]] = true
					q.PushBack(next)
				}
			}else{
				if t >= 1 && cur.node == target{
					res += cur.rate
				}
			}
		}
		t--
	}
	return res
}