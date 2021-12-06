package array

//Input: adjacentPairs = [[2,1],[3,4],[3,2]]
//Output: [1,2,3,4]
func dfs_restoreArray(graph map[int][]int, cur_num int, idx *int, res *[]int, visited map[int]bool) {
	if visited[cur_num] {
		return
	}
	(*res)[*idx] = cur_num
	visited[cur_num] = true
	*idx++
	for _, next := range graph[cur_num] {
		dfs_restoreArray(graph, next, idx, res, visited)
	}
}

func restoreArray(adjacentPairs [][]int) []int {
	var l int = len(adjacentPairs)
	var res []int = make([]int, l+1)
	var graph map[int][]int = make(map[int][]int)
	var start_num int = 0
	for i := 0; i < l; i++ {
		graph[adjacentPairs[i][0]] = append(graph[adjacentPairs[i][0]], adjacentPairs[i][1])
		graph[adjacentPairs[i][1]] = append(graph[adjacentPairs[i][1]], adjacentPairs[i][0])
	}
	for k, v := range graph {
		if len(v) == 1 {
			start_num = k
		}
	}
	var visited map[int]bool = make(map[int]bool)
	idx := 0
	dfs_restoreArray(graph, start_num, &idx, &res, visited)
	return res
}
