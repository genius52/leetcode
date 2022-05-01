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


//DFS
func dfs_minReorder(cur int,last int,graph [][]int,start_end []map[int]bool,res *int){
	for _,next := range graph[cur]{
		if next == last{
			continue
		}
		if _,ok := start_end[cur][next];ok{
			*res++
		}
		dfs_minReorder(next,cur,graph,start_end,res)
	}
}

func MinReorder2(n int, connections [][]int) int {
	var graph [][]int = make([][]int,n)
	var start_end []map[int]bool = make([]map[int]bool,n)
	for i := 0;i < n;i++{
		start_end[i] = make(map[int]bool)
	}
	for _,con := range connections{
		graph[con[0]] = append(graph[con[0]],con[1])
		graph[con[1]] = append(graph[con[1]],con[0])
		start_end[con[0]][con[1]] = true
	}
	var res int = 0
	dfs_minReorder(0,-1,graph,start_end,&res)
	return res
}