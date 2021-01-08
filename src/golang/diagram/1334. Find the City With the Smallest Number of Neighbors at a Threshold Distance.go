package diagram

//Input: n = 4, edges = [[0,1,3],[1,2,1],[1,3,4],[2,3,1]], distanceThreshold = 4
//Output: 3
//Explanation: The figure above describes the graph.
//The neighboring cities at a distanceThreshold = 4 for each city are:
//City 0 -> [City 1, City 2]
//City 1 -> [City 0, City 2, City 3]
//City 2 -> [City 0, City 1, City 3]
//City 3 -> [City 1, City 2]

func FindTheCity(n int, edges [][]int, distanceThreshold int) int {
	var graph [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		graph[i] = make([]int,n)
	}
	for i := 0;i < n;i++ {
		for j := 0; j < n; j++ {
			if i == j {
				graph[i][j] = 0
				continue
			}
			graph[i][j] = 2147483647
		}
	}
	var l int = len(edges)
	for i := 0;i < l;i++{
		graph[edges[i][0]][edges[i][1]] = edges[i][2]
		graph[edges[i][1]][edges[i][0]] = edges[i][2]
	}
	for k := 0;k < n;k++{
		for j := 0;j < n;j++{
			for i := 0;i < n;i++{
				if graph[i][j] > graph[i][k] + graph[k][j]{
					graph[i][j] = graph[i][k] + graph[k][j]
				}
			}
		}
	}
	var min_cnt int = 2147483647
	var res int = 0
	for i := 0;i < n;i++{
		var cnt int = 0
		for j := 0;j < n;j++{
			if i == j{
				continue
			}
			if graph[i][j] <= distanceThreshold{
				cnt++
			}
		}
		if cnt <= min_cnt{
			min_cnt = cnt
			res = i
		}
	}
	return res
}