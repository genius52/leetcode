package diagram

import "container/list"

func MinScore(n int, roads [][]int) int {
	var graph map[int][][2]int = make(map[int][][2]int)
	for _,edge := range roads{
		graph[edge[0]] = append(graph[edge[0]],[2]int{edge[1],edge[2]})
		graph[edge[1]] = append(graph[edge[1]],[2]int{edge[0],edge[2]})
	}
	var visited []bool = make([]bool,n + 1)
	var q list.List
	q.PushBack(1)
	var res int = 1 << 31 - 1
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			var cur int = q.Front().Value.(int)
			q.Remove(q.Front())
			visited[cur] = true
			for _,next := range graph[cur]{
				if visited[next[0]] || next[1] == 0{
					continue
				}
				if next[1] < res{
					res = next[1]
				}
				q.PushBack(next[0])
			}
		}
	}
	return res
}