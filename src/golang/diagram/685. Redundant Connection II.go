package diagram

//Input: edges = [[1,2],[2,3],[3,4],[4,1],[1,5]]
//Output: [4,1]
func union_find(groups []int,cur int)int{
	if groups[cur] != cur{
		groups[cur] = union_find(groups,groups[cur])
	}
	return groups[cur]
}

func check_loop(edges [][]int,delete_edge []int)[]int{
	//find circle
	var  l int = len(edges)
	var groups []int = make([]int,l + 1)
	for i := 1; i <= l; i++{
		groups[i] = i
	}
	for _,edge := range edges{
		if len(delete_edge) > 0{
			if edge[0] == delete_edge[0] && edge[1] == delete_edge[1]{
				continue
			}
		}
		parent := edge[0]
		child := edge[1]
		parent_group := union_find(groups,parent)
		child_group := union_find(groups,child)
		if parent_group == child_group{
			return edge
		}else{
			groups[child] = parent_group
		}
	}
	return []int{}
}

func FindRedundantDirectedConnection(edges [][]int) []int {
	var l int = len(edges)
	var graph [][]int = make([][]int,l + 1)//graph[child][parent]:
	var dup_edge [][]int
	//find indegree equal 2
	for _,edge := range edges{
		graph[edge[1]] = append(graph[edge[1]],edge[0])
		if len(graph[edge[1]]) == 2{
			dup_edge = append(dup_edge,[]int{graph[edge[1]][0],edge[1]})
			dup_edge = append(dup_edge,edge)
		}
	}

	if len(dup_edge) == 0{
		return check_loop(edges,[]int{})
	}else{
		res := check_loop(edges,dup_edge[1])
		if len(res) == 0{
			return dup_edge[1]
		}else{
			return dup_edge[0]
		}
	}
}