package diagram


//Union find
//func union_find(groups []int,cur int)int{
//	if groups[cur] != cur{
//		groups[cur] = union_find(groups,groups[cur])
//	}
//	return groups[cur]
//}

func FindCircleNum(isConnected [][]int) int {
	var n int = len(isConnected)
	var groups []int = make([]int,n)
	for i := 0;i < n;i++{
		groups[i] = i
	}
	var total int = n
	for i := 0;i < n;i++{
		for j := i + 1;j < n;j++{
			if isConnected[i][j] == 0{
				continue
			}
			group1 := union_find(groups,i)
			group2 := union_find(groups,j)
			if group1 != group2{
				groups[group2] = group1
				total--
				//if group1 < group2{
				//	groups[group2] = group1
				//}else{
				//	groups[group1] = group2
				//}
			}
		}
	}
	return total
}

//DFS solution
//func dfs_findCircleNum(graph map[int]map[int]bool,node int,visited []bool){
//	if visited[node]{
//		return
//	}
//	visited[node] = true
//	if _,ok := graph[node];!ok{ //no neighbour
//		return
//	}
//	for neighbour,_ := range graph[node]{
//		dfs_findCircleNum(graph,neighbour,visited)
//	}
//}
//
//func findCircleNum(M [][]int) int {
//	var n int = len(M)
//	var graph map[int]map[int]bool = make(map[int]map[int]bool)
//	for i := 0;i < n;i++{
//		for j := 0;j < n;j++{
//			if M[i][j] == 1{
//				if _,ok := graph[i];!ok{
//					graph[i] = make(map[int]bool)
//				}
//				if _,ok := graph[j];!ok{
//					graph[j] = make(map[int]bool)
//				}
//				graph[i][j] = true
//				graph[j][i] = true
//			}
//		}
//	}
//	var visited []bool = make([]bool,n)
//	var groups int = 0
//	for i := 0;i < n;i++{
//		if visited[i]{
//			continue
//		}
//		dfs_findCircleNum(graph,i,visited)
//		groups++
//	}
//	return groups
//}