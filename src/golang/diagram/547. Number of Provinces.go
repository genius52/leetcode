package diagram

func dfs_findCircleNum(graph map[int]map[int]bool,node int,visited []bool){
	if visited[node]{
		return
	}
	visited[node] = true
	if _,ok := graph[node];!ok{ //no neighbour
		return
	}
	for neighbour,_ := range graph[node]{
		dfs_findCircleNum(graph,neighbour,visited)
	}
}

func findCircleNum(M [][]int) int {
	var n int = len(M)
	var graph map[int]map[int]bool = make(map[int]map[int]bool)
	for i := 0;i < n;i++{
		for j := 0;j < n;j++{
			if M[i][j] == 1{
				if _,ok := graph[i];!ok{
					graph[i] = make(map[int]bool)
				}
				if _,ok := graph[j];!ok{
					graph[j] = make(map[int]bool)
				}
				graph[i][j] = true
				graph[j][i] = true
			}
		}
	}
	var visited []bool = make([]bool,n)
	var groups int = 0
	for i := 0;i < n;i++{
		if visited[i]{
			continue
		}
		dfs_findCircleNum(graph,i,visited)
		groups++
	}
	return groups
}