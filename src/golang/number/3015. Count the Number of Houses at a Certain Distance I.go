package number

func countOfPairs(n int, x int, y int) []int {
	var graph [][]int = make([][]int, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if i != j {
				graph[i][j] = abs_int(i - j)
			}
		}
	}
	graph[x][y] = 1
	graph[y][x] = 1
	for mid := 1; mid <= n; mid++ {
		for start := 1; start <= n; start++ {
			for end := 1; end <= n; end++ {
				if graph[start][mid]+graph[mid][end] <= graph[start][end] {
					graph[start][end] = graph[start][mid] + graph[mid][end]
				}
			}
		}
	}
	var res []int = make([]int, n)
	for start := 1; start <= n; start++ {
		for end := 1; end <= n; end++ {
			if start != end {
				res[graph[start][end]-1]++
			}
		}
	}
	return res
}
