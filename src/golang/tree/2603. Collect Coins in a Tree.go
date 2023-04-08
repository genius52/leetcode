package tree

func collectTheCoins(coins []int, edges [][]int) int {
	var l1 int = len(coins)
	var res int = 0
	var graph [][]int = make([][]int, l1)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return res
}
