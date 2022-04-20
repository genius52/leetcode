package tree

import "sort"

//返回最大的单向路径
func dp_longestPath(graph [][]int,s string,idx int,pre_val uint8,res *int,visited []bool)int{
	var l int = len(graph[idx])
	if s[idx] == pre_val{
		return 0
	}
	if l == 0{
		return 1
	}
	if visited[idx]{
		return 0
	}
	visited[idx] = true
	var max_len int = 1
	var record []int
	for i := 0;i < l;i++{
		child := graph[idx][i]
		sub_len := dp_longestPath(graph,s,child,s[idx],res,visited)
		if sub_len + 1 > max_len{
			max_len = sub_len + 1
		}
		if sub_len >= 1{
			record = append(record,sub_len)
		}
		dp_longestPath(graph,s,child,0,res,visited)
	}
	if len(record) > 1{
		sort.Ints(record)
		*res = max_int(*res,record[len(record) - 1] + record[len(record) - 2] + 1)
	}
	*res = max_int(*res,max_len)
	return max_len
}

func LongestPath(parent []int, s string) int {
	var l int = len(parent)
	var graph [][]int = make([][]int,l)
	for i := 1;i < l;i++{
		graph[parent[i]] = append(graph[parent[i]],i)
	}
	var visited []bool = make([]bool,l)
	var res int = 1
	ret := dp_longestPath(graph,s,0,0,&res,visited)
	return max_int(res,ret)
}