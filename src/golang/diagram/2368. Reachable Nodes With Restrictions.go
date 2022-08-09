package diagram

import "container/list"

func reachableNodes(n int, edges [][]int, restricted []int) int {
	var visited []bool = make([]bool,n)
	var graph [][]int = make([][]int,n)
	for _,edge := range edges{
		graph[edge[0]] = append(graph[edge[0]],edge[1])
		graph[edge[1]] = append(graph[edge[1]],edge[0])
	}
	var forbidden map[int]bool = make(map[int]bool)
	for _,r := range restricted{
		forbidden[r] = true
	}
	var res int = 1
	var q list.List
	q.PushBack(0)
	visited[0] = true
	for q.Len() > 0{
		var cur int = q.Front().Value.(int)
		q.Remove(q.Front())
		for _,next := range graph[cur]{
			if _,ok := forbidden[next];!ok && !visited[next]{
				q.PushBack(next)
				visited[next] = true
				res++
			}
		}
	}
	return res
}