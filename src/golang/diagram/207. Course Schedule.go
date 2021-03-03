package diagram

import "container/list"

//207
//check cycle
func dfs_canFinish(relation [][]int,cur_course int,total_course_num int,depth int)bool{
	if depth >= total_course_num{
		return false
	}
	if len(relation[cur_course]) == 0{
		return true
	}
	for _,c := range relation[cur_course]{
		if !dfs_canFinish(relation,c,total_course_num,depth + 1){
			return false
		}
	}
	return true
}

func CanFinish(numCourses int, prerequisites [][]int) bool {
	var relation [][]int = make([][]int,numCourses)
	for _,pair := range prerequisites{
		relation[pair[0]] = append(relation[pair[0]],pair[1])
	}

	for i := 0;i < numCourses;i++{
		if len(relation[i]) == 0{
			continue
		}
		for _,c := range relation[i]{
			if !dfs_canFinish(relation,c,numCourses,0){
				return false
			}
		}
	}
	return true
}

//bfs check visited
//Input: numCourses = 2, prerequisites = [[1,0]]
//Output: true
func CanFinish2(numCourses int, prerequisites [][]int) bool{
	var graph map[int]map[int]bool = make(map[int]map[int]bool)
	var indegree []int = make([]int,numCourses)
	for _,pre := range prerequisites{
		if _,ok := graph[pre[0]];!ok{
			graph[pre[0]] = make(map[int]bool)
		}
		graph[pre[0]][pre[1]] = true
		indegree[pre[1]]++
	}
	var q list.List
	for i,in_cnt := range indegree{
		if in_cnt == 0{
			q.PushBack(i)
		}
	}
	if q.Len() == 0{
		return false
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
				}
			}
		}
	}
	for _,in := range indegree{
		if in != 0{
			return false
		}
	}
	return true
}