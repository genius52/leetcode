package diagram

//Input: n = 7, headID = 6, manager = [1,2,3,4,5,6,-1], informTime = [0,6,5,4,3,2,1]
func dfs_numOfMinutes(cur_id int,graph map[int][]int,informTime []int)int{
	if _,ok := graph[cur_id];!ok{
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

func NumOfMinutes(n int, headID int, manager []int, informTime []int) int{
	var graph map[int][]int = make(map[int][]int)
	var l int = len(manager)
	for i := 0;i < l;i++{
		if manager[i] != -1{
			graph[manager[i]] = append(graph[manager[i]],i)
		}
	}
	return dfs_numOfMinutes(headID,graph,informTime)
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