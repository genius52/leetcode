package diagram

import "container/list"

//DFS
//Input: n = 7, headID = 6, manager = [1,2,3,4,5,6,-1], informTime = [0,6,5,4,3,2,1]
func dfs_numOfMinutes(cur_id int,graph [][]int,informTime []int)int{
	if len(graph[cur_id]) == 0{
		return 0
	}
	var max_duration int = 0
	for _,sub_id := range graph[cur_id]{
		cur := informTime[cur_id] + dfs_numOfMinutes(sub_id,graph,informTime)
		if cur > max_duration{
			max_duration = cur
		}
	}
	return max_duration
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int{
	var graph [][]int = make([][]int,n)
	for cur,parent := range manager{
		if parent != -1{
			graph[parent] = append(graph[parent],cur)
		}
	}
	return dfs_numOfMinutes(headID,graph,informTime)
}

//BFS
type Durantion_id struct {
	duration int
	id int
}

func NumOfMinutes(n int, headID int, manager []int, informTime []int) int{
	var manager_employee [][]int = make([][]int,n)
	var l int = len(manager)
	for i := 0;i < l;i++{
		if manager[i] != -1{
			manager_employee[manager[i]] = append(manager_employee[manager[i]],i)
		}
	}

	var q list.List
	var first Durantion_id
	first.duration = 0
	first.id = headID
	q.PushBack(first)
	var visited []bool = make([]bool,n)
	visited[headID] = true
	var res int = 0
	for q.Len() > 0{
		var l int = q.Len()
		for i := 0;i < l;i++{
			cur := q.Front().Value.(Durantion_id)
			res = max_int(res,cur.duration)
			q.Remove(q.Front())
			for j := 0;j < len(manager_employee[cur.id]);j++{
				if visited[manager_employee[cur.id][j]]{
					continue
				}
				var next Durantion_id
				next.duration = cur.duration + informTime[cur.id]
				next.id = manager_employee[cur.id][j]
				visited[manager_employee[cur.id][j]] = true
				q.PushBack(next)
			}
		}
	}
	return res
}

//func NumOfMinutes(n int, headID int, manager []int, informTime []int) int {
//	var graph map[int][]int = make(map[int][]int)
//	var l int = len(manager)
//	for i := 0;i < l;i++{
//		if manager[i] != -1{
//			graph[manager[i]] = append(graph[manager[i]],i)
//		}
//	}
//	var res int = 0
//	var q list.List
//	q.PushBack(headID)
//	for q.Len() > 0{
//		var cur_len int = q.Len()
//		var max_duration int = 0
//		for i := 0;i < cur_len;i++{
//			var subid int = q.Front().Value.(int)
//			if informTime[subid] > max_duration{
//				max_duration = informTime[subid]
//			}
//			if sub_slice,ok := graph[subid];ok{
//				for _,sub := range sub_slice{
//					q.PushBack(sub)
//				}
//			}
//			q.Remove(q.Front())
//		}
//		res += max_duration
//	}
//	return res
//}