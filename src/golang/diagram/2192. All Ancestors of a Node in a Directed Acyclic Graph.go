package diagram

import "sort"

func dfs_getAncestors(graph [][]int,cur int,ans int,visited []bool,res [][]int){
	visited[cur] = true
	for _,child := range graph[cur]{
		if !visited[child]{
			res[child] = append(res[child],ans)
			dfs_getAncestors(graph,child,ans,visited,res)
		}
	}
}

func GetAncestors(n int, edges [][]int) [][]int {
	var graph [][]int = make([][]int,n)
	//var indegree []int = make([]int,n)
	for _,edge := range edges{
		graph[edge[0]] = append(graph[edge[0]],edge[1])
		//indegree[edge[1]]++
	}
	var res [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		var visited []bool = make([]bool,n)
		dfs_getAncestors(graph,i,i,visited,res)
	}
	for i := 0;i < n;i++{
		sort.Ints(res[i])
	}
	return res
}