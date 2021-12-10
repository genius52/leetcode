package diagram

func minTrioDegree(n int, edges [][]int) int {
	var graph [][]bool = make([][]bool, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]bool, n+1)
	}
	var degree []int = make([]int, n+1)
	for _, edge := range edges {
		graph[min_int(edge[0], edge[1])][max_int(edge[0], edge[1])] = true
		degree[edge[0]]++
		degree[edge[1]]++
	}
	var res int = 2147483647
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if !graph[i][j] {
				continue
			}
			for k := j + 1; k <= n; k++ {
				if graph[i][k] && graph[j][k] {
					res = min_int(res, degree[i]+degree[j]+degree[k]-6)
				}
			}
		}
	}
	if res == 2147483647 {
		return -1
	}
	return res
}
