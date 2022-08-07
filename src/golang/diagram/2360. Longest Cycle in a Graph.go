package diagram


func dfs_longestCycle(edges []int,l int,cur int,timestamp int,visited []bool,trace []int,total_removed map[int]bool,removed map[int]bool,max_val *int){
	if visited[cur] {
		if _,ok := total_removed[cur];ok{
			return
		}
		trace[cur] = timestamp - trace[cur]
		if trace[cur] > *max_val{
			*max_val = trace[cur]
		}
		return
	}
	visited[cur] = true
	removed[cur] = true
	if edges[cur] == -1{
		return
	}
	trace[cur] = timestamp
	dfs_longestCycle(edges,l,edges[cur],timestamp + 1,visited,trace,total_removed,removed,max_val)
}

func LongestCycle(edges []int) int {
	var l int = len(edges)
	var visited []bool = make([]bool,l)
	var trace []int = make([]int,l)
	var total_removed map[int]bool = make(map[int]bool)
	var res int = -1
	for i := 0;i < l;i++{
		if visited[i] || edges[i] == -1{
			continue
		}
		var removed map[int]bool = make(map[int]bool)
		dfs_longestCycle(edges,l,i,0,visited,trace,total_removed,removed,&res)
		for k,_ := range removed{
			total_removed[k] = true
		}
	}
	return res
}