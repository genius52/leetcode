package diagram

func dfs_minimumTime(graph [][][2]int, cur_node int, cur_duration int, disappear []int, duration_from_zero []int) {
	for _, next := range graph[cur_node] {
		if cur_duration+next[1] >= disappear[next[0]] {
			continue
		}
		if duration_from_zero[next[0]] == -1 || cur_duration+next[1] < duration_from_zero[next[0]] {
			duration_from_zero[next[0]] = cur_duration + next[1]
			dfs_minimumTime(graph, next[0], cur_duration+next[1], disappear, duration_from_zero)
		}
	}
}

func minimumTime(n int, edges [][]int, disappear []int) []int {
	var graph [][][2]int = make([][][2]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], [2]int{edge[1], edge[2]})
		graph[edge[1]] = append(graph[edge[1]], [2]int{edge[0], edge[2]})
	}
	var duration_from_zero []int = make([]int, n)
	for i := 1; i < n; i++ {
		duration_from_zero[i] = -1
	}
	dfs_minimumTime(graph, 0, 0, disappear, duration_from_zero)
	return duration_from_zero
}
