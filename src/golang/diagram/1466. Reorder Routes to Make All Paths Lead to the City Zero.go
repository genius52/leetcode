package diagram

import "container/list"

func MinReorder(n int, connections [][]int) int {
	var out [][]int = make([][]int,n)
	var in [][]int = make([][]int,n)
	for _,con := range connections{
		out[con[0]] = append(out[con[0]],con[1])
		in[con[1]] = append(in[con[1]],con[0])
	}
	var q list.List
	var res int = 0
	var visited []bool = make([]bool,n)
	visited[0] = true
	q.PushBack(0)
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var cur int = q.Front().Value.(int)
			q.Remove(q.Front())
			for j := 0;j < len(out[cur]);j++{
				if !visited[out[cur][j]]{
					res++
					visited[out[cur][j]] = true
					q.PushBack(out[cur][j])
				}
			}
			for j := 0;j < len(in[cur]);j++{
				if !visited[in[cur][j]]{
					visited[in[cur][j]] = true
					q.PushBack(in[cur][j])
				}
			}
		}
	}
	return res
}