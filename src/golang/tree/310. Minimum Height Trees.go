package tree

import "container/list"

func FindMinHeightTrees(n int, edges [][]int) []int {
	if n == 0{
		return []int{0}
	}
	var graph map[int]map[int]bool = make(map[int]map[int]bool)
	for _,edge := range edges{
		if _,ok := graph[edge[0]];!ok{
			graph[edge[0]] = make(map[int]bool)
		}
		graph[edge[0]][edge[1]] = true
		if _,ok := graph[edge[1]];!ok{
			graph[edge[1]] = make(map[int]bool)
		}
		graph[edge[1]][edge[0]] = true
	}
	var q list.List
	var total int = n
	for node,neighbours := range graph {
		if len(neighbours) == 1 {
			q.PushBack(node)
			//visited[node] = true
		}
	}
	for total > 2{
		var l int = q.Len()
		total -= l
		for i := 0;i < l;i++{
			var node int = q.Front().Value.(int)
			q.Remove(q.Front())
			for neighbour,_ := range graph[node]{
				delete(graph[neighbour],node)
				if len(graph[neighbour]) == 1{
					q.PushBack(neighbour)
				}
			}
		}
	}
	var res []int
	for q.Len() > 0{
		var node int = q.Front().Value.(int)
		q.Remove(q.Front())
		res = append(res,node)
	}
	return res
}