package diagram

import "sort"

//kruskal solution
func get_parent(groups []int,node int)int{
	if groups[node] != node{
		groups[node] = get_parent(groups,groups[node])
	}
	return groups[node]
}

func MinCostConnectPoints(points [][]int) int {
	var l int = len(points)
	var edges [][]int
	for i := 0;i < l;i++{
		for j := i + 1;j < l;j++{
			if i == j{
				continue
			}
			dis := abs_int(points[i][0] - points[j][0]) + abs_int(points[i][1] - points[j][1])
			edges = append(edges,[]int{dis,i,j})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})
	var groups []int = make([]int,l)
	for i := 0;i < l;i++{
		groups[i] = i
	}
	var edge_len int = len(edges)
	groups_cnt := l
	var sum int = 0
	for i := 0;i < edge_len;i++{
		distance := edges[i][0]
		src := edges[i][1]
		dst := edges[i][2]
		group1 := get_parent(groups,src)
		group2 := get_parent(groups,dst)
		if group1 != group2{
			groups[group2] = group1
			sum += distance
			groups_cnt--
		}
		if groups_cnt == 1{
			break
		}
	}
	return sum
}