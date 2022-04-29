package diagram

//DFS
func dfs_makeConnected(pre_node int,cur_node int,graph map[int]map[int]bool,visited []bool)int{
	if visited[cur_node]{
		return 0
	}
	visited[cur_node] = true
	if _,ok := graph[cur_node];!ok{
		return 0
	}
	var dup_lines int = 0
	for neigh,_:= range graph[cur_node]{
		if visited[neigh]{
			if pre_node != neigh && pre_node != -1{
				dup_lines++
			}
		}
		dup_lines += dfs_makeConnected(cur_node,neigh,graph,visited)
	}
	return dup_lines
}

func MakeConnected(n int, connections [][]int) int {
	var visited []bool = make([]bool,n)
	var graph map[int]map[int]bool = make(map[int]map[int]bool)
	var l int = len(connections)
	for i := 0;i < l;i++{
		if _,ok := graph[connections[i][0]];!ok{
			graph[connections[i][0]] = make(map[int]bool)
		}
		graph[connections[i][0]][connections[i][1]] = true
		if _,ok := graph[connections[i][1]];!ok{
			graph[connections[i][1]] = make(map[int]bool)
		}
		graph[connections[i][1]][connections[i][0]] = true
	}
	var groups int = 0
	var dup_lines int = 0
	for i := 0;i < n;i++{
		if visited[i]{
			continue
		}
		dup_lines += dfs_makeConnected(-1,i,graph,visited)
		groups++
	}
	if dup_lines < groups - 1 {
		return -1
	}
	return groups - 1
}

//Union find
func MakeConnected2(n int, connections [][]int) int{
	var graph [][]int = make([][]int,n)
	var total_lines int = 0
	for _,con := range connections{
		graph[con[0]] = append(graph[con[0]],con[1])
		graph[con[1]] = append(graph[con[1]],con[0])
		total_lines++
	}
	var groups []int = make([]int,n)
	for i := 0;i < n;i++{
		groups[i] = i
	}
	for i := 0;i < n;i++{
		group1 := get_parent(groups,i)
		for j := 0;j < len(graph[i]);j++{
			group2 := get_parent(groups,graph[i][j])
			if group1 != group2{
				if group1 < group2{
					groups[group2] = group1
				}else{
					groups[group1] = group2
				}
			}
		}
	}
	var record map[int]bool = make(map[int]bool)
	for i := 0;i < n;i++{
		record[get_parent(groups,i)] = true
	}
	group_cnt := len(record)
	//if group_cnt == 1{
	//	return 0
	//}
	need := group_cnt - 1 //需要额外的线
	duplines := total_lines - (n - 1) + (group_cnt - 1) //4个点，当在一个组的时候，多余的线为 total_lines - (4 - 1) + (group_cnt - 1)
	//标准为一个组，groups - 1条线。每多1个组，就多1条多余的线
	if duplines < need{
		return -1
	}
	return need
}