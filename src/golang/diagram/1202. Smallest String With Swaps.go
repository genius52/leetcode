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

func SmallestStringWithSwaps2(s string, pairs [][]int) string{
	var l int = len(s)
	var groups []int = make([]int,l)
	for i := 0;i < l;i++{
		groups[i] = i
	}
	for _,pair := range pairs{
		i := pair[0]
		j := pair[1]
		group1 := get_parent(groups,i)
		group2 := get_parent(groups,j)
		if group1 != group2{
			if group1 < group2{
				groups[group2] = group1
			}else{
				groups[group1] = group2
			}
		}
	}
	var record [][]int = make([][]int,l) //将同一组的索引放在一起
	var ss [][]string = make([][]string,l)
	for i := 0;i < l;i++{
		group := get_parent(groups,i)
		record[group] = append(record[group],i)
		ss[group] = append(ss[group],string(s[i]))
	}
	var res []string = make([]string,l)
	for idx,r := range record{
		if len(r) == 0{
			continue
		}
		sort.Strings(ss[idx])
		for i := 0;i < len(r);i++{
			res[r[i]] = ss[idx][i]
		}
	}
	return strings.Join(res,"")
}