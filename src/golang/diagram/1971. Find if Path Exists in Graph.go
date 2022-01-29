package diagram

import "container/list"

func validPath(n int, edges [][]int, source int, destination int) bool {
	var graph [][]int = make([][]int,n)
	for _,edge := range edges{
		graph[edge[0]] = append(graph[edge[0]],edge[1])
		graph[edge[1]] = append(graph[edge[1]],edge[0])
	}
	var q list.List
	q.PushBack(source)
	var visited []bool = make([]bool,n)
	visited[source] = true
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			cur := q.Front().Value.(int)
			q.Remove(q.Front())
			if cur == destination{
				return true
			}
			for j := 0;j < len(graph[cur]);j++{
				if visited[graph[cur][j]]{
					continue
				}
				visited[graph[cur][j]] = true
				q.PushBack(graph[cur][j])
			}
		}
	}
	return false
}