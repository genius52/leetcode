package diagram

import "container/list"

//210
//Input: 4, [[1,0],[2,0],[3,1],[3,2]]
//Output: [0,1,2,3] or [0,2,1,3]
func dfs_findOrder(relation [][]int,cur_course int,visited []int,order *[]int) bool{
	if visited[cur_course] == -1{
		return false
	}
	visited[cur_course] = -1
	for _,c := range relation[cur_course]{
		//-1 means loop existed
		if visited[c] == 1{
			continue
		}
		if visited[c] == -1{
			return false
		}
		if !dfs_findOrder(relation,c,visited,order){
			return false
		}
	}
	visited[cur_course] = 1
	*order = append(*order, cur_course)
	return true
}

func FindOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0{
		return []int{}
	}
	var relation [][]int = make([][]int,numCourses)
	for _,pair := range prerequisites{
		relation[pair[0]] = append(relation[pair[0]],pair[1])
	}
	var visited []int = make([]int,numCourses)  //0: not visited  1:visited and pass   -1:start current search,prevent loop
	var order []int
	for i := 0;i < numCourses;i++{
		if visited[i] != 0{
			continue
		}
		if !dfs_findOrder(relation,i,visited,&order){
			return []int{}
		}
	}
	return order
}

func bfs_FindOrder(numCourses int, prerequisites [][]int) []int{
	var graph map[int]map[int]bool = make(map[int]map[int]bool)
	var indegree []int = make([]int,numCourses)
	for _,pre := range prerequisites{
		if _,ok := graph[pre[1]];!ok{
			graph[pre[1]] = make(map[int]bool)
		}
		graph[pre[1]][pre[0]] = true
		indegree[pre[0]]++
	}
	var res []int
	var q list.List
	for i,in_cnt := range indegree{
		if in_cnt == 0{
			q.PushBack(i)
			res = append(res,i)
		}
	}
	if q.Len() == 0{
		return res
	}
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			var node int = q.Back().Value.(int)
			q.Remove(q.Back())
			for next,_ := range graph[node]{
				indegree[next]--
				if indegree[next] == 0{
					q.PushBack(next)
					res = append(res,next)
				}
			}
		}
	}
	if len(res) < numCourses{
		return []int{}
	}
	return res
}