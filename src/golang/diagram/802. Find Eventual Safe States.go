package diagram

import (
	"container/list"
	"sort"
)

//Input: graph = [[1,2],[2,3],[5],[0],[5],[],[]]
//Output: [2,4,5,6]
const (
	cycle int = 1
	nocycle int  = 2
)

func hascycle_eventualSafeNodes(record [][]int,cur int,visited []int)bool{
	if visited[cur] == cycle{
		return true
	}
	if visited[cur] == nocycle{
		return false
	}
	visited[cur] = cycle
	var ret bool = false
	for i := 0;i < len(record[cur]);i++{
		next := record[cur][i]
		ret = hascycle_eventualSafeNodes(record,next,visited)
		if ret{
			break
		}
	}
	if !ret{
		visited[cur] = nocycle
	}
	return ret
}


func EventualSafeNodes(graph [][]int) []int {
	var l int = len(graph)
	var res []int
	var visited []int = make([]int,l)
	for i := 0;i < l;i++{
		if visited[i] == cycle{
			continue
		}else if visited[i] == nocycle{
			res = append(res,i)
			continue
		}
		hascycle := hascycle_eventualSafeNodes(graph,i,visited)
		if !hascycle{
			res = append(res,i)
			//visited[i] = nocycle
		}
	}
	return res
}

//dfs
//1: safe node
//2: unsafe node
func dfs_eventualSafeNodes(end_start [][]int,cur int,is_safe []bool,visited []bool){
	if len(end_start[cur]) == 0{
		return
	}
	if visited[cur]{
		return
	}
	visited[cur] = true
	for _,next := range end_start[cur]{
		is_safe[next] = false
		dfs_eventualSafeNodes(end_start,next,is_safe,visited)
	}
}

func EventualSafeNodes2(graph [][]int) []int {
	var l int = len(graph)
	var end_start []map[int]bool = make([]map[int]bool,l)
	for i := 0;i < l;i++{
		end_start[i] = make(map[int]bool)
	}
	var outdegree []int = make([]int,l)
	var res []int
	var q list.List //安全节点
	for i := 0;i < l;i++{
		if len(graph[i]) == 0{
			q.PushBack(i)
			res = append(res,i)
		}else{
			outdegree[i] = len(graph[i])
			for _,end := range graph[i]{
				end_start[end][i] = true
			}
		}
	}
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			end := q.Front().Value.(int)
			q.Remove(q.Front())
			for start,_ := range end_start[end]{
				outdegree[start]--
				if outdegree[start] == 0{
					q.PushBack(start)
					res = append(res,start)
				}
			}
		}
	}
	sort.Ints(res)
	return res
}