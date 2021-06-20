package diagram

//Input: edges = [[1,2],[2,3],[3,4],[4,1],[1,5]]
//Output: [4,1]
func find_group(group []int,i int)int {
	if group[i] != i{
		group[i] = find_group(group,group[i]);
	}
	return group[i];
}

func FindRedundantDirectedConnection(edges [][]int) []int {
	var l int = len(edges)
	var group []int = make([]int,l + 1)
	for i := 0;i <= l;i++{
		group[i] = i
	}
	var indegree map[int]int = make(map[int]int)
	var indegree_two int = -1
	var circle [][]int
	var circle_num int = -1
	for _,edge := range edges{
		if _,ok := indegree[edge[1]];!ok{
			indegree[edge[1]] = 1
		}else{
			indegree[edge[1]]++
			indegree_two = edge[1]
		}
		group1 := find_group(group,edge[0])
		group2 := find_group(group,edge[1])
		if group1 == group2 {
			//circle = append(circle,edge)
			circle_num = group1
		}else{
			group[group1] = group2
		}
	}
	var indegree_res [][]int
	for _,edge := range edges{
		if indegree_two != -1{
			if edge[1] == indegree_two{
				indegree_res = append(indegree_res,edge)
			}
		}
		if circle_num != -1{
			if edge[0] == circle_num && edge[1] == circle_num{
				circle = append(circle,edge)
			}
		}
	}

	if len(circle) == 0{
		if len(indegree_res) > 0{
			return indegree_res[1]
		}
	}
	for _,in := range indegree_res{
		for _,c := range circle{
			if in[0] == c[0] && in[1] == c[1]{
				return in
			}
		}
	}
	return circle[len(circle) - 1]
}