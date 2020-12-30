package diagram

import (
	"sort"
	"strings"
)

//Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
//Output: "abcd"
func dfs_smallestStringWithSwaps(cur_index int,l int,group *[]int,graph map[int]map[int]bool,visited []bool){
	if visited[cur_index]{
		return
	}
	visited[cur_index] = true
	*group = append(*group,cur_index)
	for k,_:= range graph[cur_index]{
		if visited[k]{
			continue
		}
		dfs_smallestStringWithSwaps(k,l,group,graph,visited)
	}
}

func SmallestStringWithSwaps(s string, pairs [][]int) string {
	var l int = len(s)
	var visited []bool = make([]bool,l)
	var graph map[int]map[int]bool = make(map[int]map[int]bool)
	for i := 0;i < l;i++{
		graph[i] = make(map[int]bool)
	}
	for _,p := range pairs{
		graph[p[0]][p[1]] = true
		graph[p[1]][p[0]] = true
	}
	var groups [][]int
	for i := 0;i < l;i++{
		if visited[i]{
			continue
		}
		var group []int
		dfs_smallestStringWithSwaps(i,l,&group,graph,visited)
		groups = append(groups,group)
	}
	var res []string = make([]string,l)
	var group_len int = len(groups)
	for i := 0;i < group_len;i++{
		var sub_len int = len(groups[i])
		var sub []string = make([]string,sub_len)
		for j := 0;j < sub_len;j++{
			sub[j] = string(s[groups[i][j]])
		}
		sort.Strings(sub)
		sort.Ints(groups[i])
		for k := 0;k < sub_len;k++{
			res[groups[i][k]] = sub[k]
		}
	}
	return strings.Join(res,"")
}