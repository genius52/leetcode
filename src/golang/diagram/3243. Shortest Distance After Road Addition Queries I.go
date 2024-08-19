package diagram

import "container/list"

func ShortestDistanceAfterQueries(n int, queries [][]int) []int {
	var graph [][]int = make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
	}
	//var distance_from_zero []int = make([]int, n)
	//for i := 1; i < n; i++ {
	//	distance_from_zero[i] = i
	//}
	var l int = len(queries)
	var res []int = make([]int, l)
	for i := 0; i < l; i++ {
		src := queries[i][0]
		dst := queries[i][1]
		graph[src] = append(graph[src], dst)
		//distance_from_zero[n - 1] = min(distance_from_zero[n - 1],1 + distance_from_zero[])
		var visited []bool = make([]bool, n)
		visited[0] = true
		var q list.List
		q.PushBack(0)
		var steps int = 0
	loop:
		for q.Len() > 0 {
			var cur_len int = q.Len()
			for j := 0; j < cur_len; j++ {
				cur := q.Front().Value.(int)
				q.Remove(q.Front())
				if cur == n-1 {
					res[i] = steps
					break loop
				}
				for _, next := range graph[cur] {
					if !visited[next] {
						q.PushBack(next)
						visited[next] = true
					}
				}
			}
			steps++
		}
	}
	return res
}
