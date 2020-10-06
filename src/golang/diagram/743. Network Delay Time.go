package diagram

func dfs_networkDelayTime(graph [][]int,N int,cur_node int,visited []bool)int{
	if visited[cur_node]{
		return 0
	}
	visited[cur_node] = true
	var res int = 0
	for i := 0;i <= N;i++{
		if graph[cur_node][i] == 0 || visited[i]{
			continue
		}
		depth := graph[cur_node][i] + dfs_networkDelayTime(graph,N,i,visited)
		if depth > res{
			res = depth
		}
	}
	return res
}

func NetworkDelayTime(times [][]int, N int, K int) int {
	var graph [][]int = make([][]int,N + 1)
	for i := 0;i <= N;i++{
		graph[i] = make([]int,N + 1)
	}
	var l int = len(times)
	for i := 0;i < l;i++{
		graph[times[i][0]][times[i][1]] = times[i][2]
	}
	var visited []bool = make([]bool,N + 1)
	depth := dfs_networkDelayTime(graph,N,K,visited)
	for i := 1;i <= N;i++{
		if !visited[i]{
			return -1
		}
	}
	return depth
}