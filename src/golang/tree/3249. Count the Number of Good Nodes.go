package tree

func dfs_countGoodNodes(graph [][]int, cur int, last int, res *int) int {
	var node_cnt int = 0
	var same bool = true
	var sum int = 0
	for _, next := range graph[cur] {
		if next == last {
			continue
		}
		cur := dfs_countGoodNodes(graph, next, cur, res)
		sum += cur
		if node_cnt == 0 {
			node_cnt = cur
		} else {
			if cur != node_cnt {
				same = false
			}
		}
	}
	if same {
		*res++
	}
	return 1 + sum
}

func CountGoodNodes(edges [][]int) int {
	var n int = len(edges)
	var graph [][]int = make([][]int, n+1)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	var res int = 0
	dfs_countGoodNodes(graph, 0, -1, &res)
	return res
}
