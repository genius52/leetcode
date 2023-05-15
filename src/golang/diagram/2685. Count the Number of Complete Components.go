package diagram

func dfs_countCompleteComponents(cur int, no int, graph [][]int, nodes *[]int, visited []bool) {
	*nodes = append(*nodes, cur)
	visited[cur] = true
	for _, next := range graph[cur] {
		if visited[next] {
			continue
		}
		dfs_countCompleteComponents(next, no, graph, nodes, visited)
	}
}

func CountCompleteComponents(n int, edges [][]int) int {
	var graph [][]int = make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	var groups []int = make([]int, n)
	var no int = 1
	var res int = 0
	for i := 0; i < n; i++ {
		if groups[i] != 0 {
			continue
		}
		var visited []bool = make([]bool, n)
		var nodes []int
		dfs_countCompleteComponents(i, no, graph, &nodes, visited)
		var meet bool = true
		node_cnt := len(nodes)
		for j := 0; j < node_cnt; j++ {
			groups[nodes[j]] = no
			if len(graph[nodes[j]]) != (node_cnt - 1) {
				meet = false
			}
		}
		if meet {
			res++
		}
		no++
	}
	return res
}
