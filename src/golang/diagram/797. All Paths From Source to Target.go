package diagram

import "container/list"

func allPathsSourceTarget(graph [][]int) [][]int {
	var n int = len(graph)
	var route [][]bool = make([][]bool,n)
	for i := 0;i < n;i++{
		route[i] = make([]bool,n)
	}
	for i := 0;i < n;i++{
		for j := 0;j < len(graph[i]);j++{
			route[i][graph[i][j]] = true
		}
	}
	var res [][]int
	var q list.List
	q.PushBack([]int{0})
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var cur_trace []int = q.Front().Value.([]int)
			q.Remove(q.Front())
			var l2 int = len(cur_trace)
			if l2 > n{
				continue
			}
			if cur_trace[l2 - 1] == n - 1{
				res = append(res,cur_trace)
				continue
			}
			for j := 0;j < n;j++{
				if route[cur_trace[l2 - 1]][j]{
					var next_trace []int = make([]int,l2)
					copy(next_trace,cur_trace)
					next_trace = append(next_trace,j)
					q.PushBack(next_trace)
				}
			}
		}
	}
	return res
}

//dfs
func allPathsSourceTarget2(graph [][]int) [][]int {
	var res [][]int
	var target int = len(graph) - 1
	var route []int
	dfs(target,0,graph,route,&res)
	return res
}
func dfs(target int,cur_pos int,graph [][]int,route []int,res *[][]int) {
	if cur_pos == target{
		route = append(route, target)
		*res = append(*res, route)
		return
	}else{
		for i := 0; i < len(graph[cur_pos]);i++{
			var new_route []int = make([]int,len(route))
			copy(new_route,route)
			new_route = append(new_route, cur_pos)
			dfs(target,graph[cur_pos][i],graph,new_route,res)
		}
	}
}

