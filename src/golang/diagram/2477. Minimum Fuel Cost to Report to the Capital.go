package diagram

func dfs_minimumFuelCost(graph [][]int,node int,seats int,visited []bool,res *int64)int{
	if visited[node]{
		return 0
	}
	visited[node] = true
	var node_cnt int = 1
	for _,next := range graph[node]{
		node_cnt += dfs_minimumFuelCost(graph,next,seats,visited,res)
	}
	if node != 0{
		//所有汽车前进一步耗费的汽油
		*res += int64((node_cnt - 1 + seats)/ seats)
		//*res += int64(math.Ceil(float64(node_cnt)/ float64(seats)))
	}
	return node_cnt
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	var l int = len(roads) + 1
	var graph [][]int = make([][]int,l)
	for _,road := range roads{
		graph[road[0]] = append(graph[road[0]],road[1])
		graph[road[1]] = append(graph[road[1]],road[0])
	}
	var visited []bool = make([]bool,l)
	var res int64 = 0
	dfs_minimumFuelCost(graph,0,seats,visited,&res)
	return res
}