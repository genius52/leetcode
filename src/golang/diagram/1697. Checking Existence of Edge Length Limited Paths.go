package diagram

import "sort"

func DistanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	var q_len int = len(queries)
	var record []int = make([]int,q_len)
	for i := 0;i < q_len;i++{
		record[i] = i
	}
	type funcType func (groups []int,node int)int
	var get_parent funcType
	get_parent = func (groups []int,node int)int{
		if groups[node] != node{
			groups[node] = get_parent(groups,groups[node])
		}
		return groups[node]
	}
	sort.Slice(record, func(i, j int) bool {
		return queries[record[i]][2] < queries[record[j]][2]
	})
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})
	var groups []int = make([]int,n)
	for i := 0;i < n;i++{
		groups[i] = i
	}
	var res []bool = make([]bool,q_len)
	var edge_idx int = 0
	var edge_len int = len(edgeList)
	for i := 0;i < q_len;i++{
		for edge_idx < edge_len && edgeList[edge_idx][2] < queries[record[i]][2]{
			node1 := edgeList[edge_idx][0]
			node2 := edgeList[edge_idx][1]
			group1 := get_parent(groups,node1)
			group2 := get_parent(groups,node2)
			if group1 != group2{
				if group1 < group2{//组号小的优先
					groups[group2] = group1
				}else{
					groups[group1] = group2
				}
			}
			edge_idx++
		}
		n1 := queries[record[i]][0]
		n2 := queries[record[i]][1]
		if get_parent(groups,n1) == get_parent(groups,n2){
			res[record[i]] = true
		}else{
			res[record[i]] = false
		}
	}
	return res
}