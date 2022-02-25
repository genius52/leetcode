package diagram

func dfs_maximalPathQuality(values []int,l int,graph [][][2]int,visited []int,cur int,time int,points int)int{
	var max_point int = 0
	if cur == 0{
		max_point = points
	}
	for _,item := range graph[cur]{
		next_node := item[0]
		next_cost := item[1]
		if time >= next_cost{
			if visited[next_node] == 0{
				visited[next_node]++
				ret := dfs_maximalPathQuality(values,l,graph,visited,next_node,time - next_cost,points + values[next_node])
				if ret > max_point{
					max_point = ret
				}
				visited[next_node]--
			}else{
				visited[next_node]++
				ret := dfs_maximalPathQuality(values,l,graph,visited,next_node,time - next_cost,points)
				if ret > max_point{
					max_point = ret
				}
				visited[next_node]--
			}
		}
	}
	return max_point
}

func MaximalPathQuality(values []int, edges [][]int, maxTime int) int {
	var l int = len(values)
	var visited []int = make([]int,l)
	var graph [][][2]int = make([][][2]int,l)
	for _,edge := range edges{
		graph[edge[0]] = append(graph[edge[0]],[2]int{edge[1],edge[2]})
		graph[edge[1]] = append(graph[edge[1]],[2]int{edge[0],edge[2]})
	}
	visited[0]++
	return dfs_maximalPathQuality(values,l,graph,visited,0,maxTime,values[0])
}