package tree

func max_int64(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func dfs_maximumScoreAfterOperations(graph [][]int, values []int, cur int, parent int) (int64, int64) {
	if len(graph[cur]) == 1 {
		return int64(values[cur]), 0
	}
	var take_node int64 = 0
	var not_take_node int64 = 0
	for _, next := range graph[cur] {
		if next != parent {
			ret1, ret2 := dfs_maximumScoreAfterOperations(graph, values, next, cur)
			take_node += ret1
			not_take_node += ret2
		}
	}
	//take_node = max_int64(not_take_node, int64(values[cur]))
	return take_node, not_take_node
}

func MaximumScoreAfterOperations(edges [][]int, values []int) int64 {
	var l int = len(values)
	var graph [][]int = make([][]int, l)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return 0
}
