package diagram

//tarjan
func dfs_criticalConnections(graph [][]int,cur int,parent int,id int,groups []int,res *[][]int){
	groups[cur] = id
	for _,neighbour := range graph[cur]{
		if neighbour == parent{
			continue
		}
		if groups[neighbour] == -1{ //之前没有走过,走一遍
			dfs_criticalConnections(graph,neighbour,cur,id + 1,groups,res)
		}
		if id < groups[neighbour]{//如果当前组号比下一个组号小,说明没有形成环，而是一条直线，是必不可少的通路。
			*res = append(*res,[]int{cur,neighbour})
		}
		groups[cur] = min_int(groups[cur],groups[neighbour])
	}
}

func CriticalConnections(n int, connections [][]int) [][]int {
	var graph [][]int = make([][]int,n)
	for _,con := range connections{
		graph[con[0]] = append(graph[con[0]],con[1])
		graph[con[1]] = append(graph[con[1]],con[0])
	}
	var groups []int = make([]int,n)
	for i := 0;i < n;i++{
		groups[i] = -1
	}
	var res [][]int
	dfs_criticalConnections(graph,0,-1,1,groups,&res)
	return res
}