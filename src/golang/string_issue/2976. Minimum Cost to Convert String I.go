package string_issue

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	var l1 int = len(source)
	var l2 int = len(cost)
	var graph [26][26]int
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i != j {
				graph[i][j] = 1<<31 - 1
			}
		}
	}
	for i := 0; i < l2; i++ {
		graph[int(original[i]-'a')][int(changed[i]-'a')] = min_int(graph[int(original[i]-'a')][int(changed[i]-'a')], cost[i])
	}
	for mid := 0; mid < 26; mid++ {
		for start := 0; start < 26; start++ {
			for end := 0; end < 26; end++ {
				if start == end || start == mid || end == mid {
					continue
				}
				if graph[start][mid] != 1<<31-1 && graph[mid][end] != 1<<31-1 {
					graph[start][end] = min_int(graph[start][end], graph[start][mid]+graph[mid][end])
				}
			}
		}
	}
	var res int64 = 0
	for i := 0; i < l1; i++ {
		if source[i] != target[i] {
			val := graph[int(source[i]-'a')][int(target[i]-'a')]
			if val == 1<<31-1 {
				return -1
			}
			res += int64(val)
		}
	}
	return res
}
