package tree

// 从 c 到 a 的距离是可以被 signalSpeed 整除的。
// 从 c 到 b 的距离是可以被 signalSpeed 整除的。
// 从 c 到 b 的路径与从 c 到 a 的路径没有任何公共边。
func dfs_countPairsOfConnectableServers(graph [][][2]int, cur int, parent int, sum_len int, signalSpeed int) int {
	var res int = 0
	if sum_len != 0 && sum_len%signalSpeed == 0 {
		res++
	}
	if len(graph[cur]) == 1 {
		return res
	}
	//var cnt []int
	for _, neigh := range graph[cur] {
		if neigh[0] == parent {
			continue
		}
		res += dfs_countPairsOfConnectableServers(graph, neigh[0], cur, sum_len+neigh[1], signalSpeed)
	}
	//if graph[cur][0][0] != last {
	//	res += dfs_countPairsOfConnectableServers(graph, graph[cur][0][0], cur, sum_len+graph[cur][0][1], signalSpeed)
	//}
	//if graph[cur][1][0] != last {
	//	res += dfs_countPairsOfConnectableServers(graph, graph[cur][1][0], cur, sum_len+graph[cur][1][1], signalSpeed)
	//}
	return res
}

func CountPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	var l int = len(edges)
	var count []int = make([]int, l+1)
	var graph [][][2]int = make([][][2]int, l+1)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], [2]int{edge[1], edge[2]})
		graph[edge[1]] = append(graph[edge[1]], [2]int{edge[0], edge[2]})
	}
	for i := 0; i <= l; i++ {
		if len(graph[i]) == 1 {
			continue
		}
		var sum int = 0
		for _, neigh := range graph[i] {
			ret := dfs_countPairsOfConnectableServers(graph, neigh[0], i, neigh[1], signalSpeed)
			count[i] += ret * sum
			sum += ret
		}
	}
	return count
}
