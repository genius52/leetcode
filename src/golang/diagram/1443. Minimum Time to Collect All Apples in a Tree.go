package diagram

func dfs_minTime(graph *[][]int,hasApple *[]bool,cur int,visited *[]bool)int{
	if (*visited)[cur]{
		return 0
	}
	(*visited)[cur] = true
	var children int = 0
	for i := 0;i < len((*graph)[cur]);i++{
		if (*visited)[(*graph)[cur][i]]{
			continue
		}
		children += dfs_minTime(graph,hasApple,(*graph)[cur][i],visited)
	}
	if !(*hasApple)[cur] && children == 0{
		return 0
	}else{
		return children + 2
	}
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	var graph [][]int = make([][]int,n)
	for _,edge := range edges{
		graph[edge[0]] = append(graph[edge[0]],edge[1])
		graph[edge[1]] = append(graph[edge[1]],edge[0])
	}
	var visited []bool = make([]bool,n)
	res :=  dfs_minTime(&graph,&hasApple,0,&visited) - 2
	if res < 0{
		return 0
	}
	return res
}