package array

import "container/list"

func MinJumps(arr []int) int {
	var l int = len(arr)
	var visited []bool = make([]bool,l)
	var same_val map[int]map[int]bool = make(map[int]map[int]bool)
	for i := 1;i < l;i++{
		if _,ok := same_val[arr[i]];!ok{
			same_val[arr[i]] = make(map[int]bool)
		}
		same_val[arr[i]][i] = true
	}
	var q list.List
	q.PushBack(0)
	visited[0] = true
	var steps int = 0
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			var cur int = q.Front().Value.(int)
			if cur == l - 1{
				return steps
			}
			q.Remove(q.Front())
			if _,ok := same_val[arr[cur]];ok{
				for idx,_ := range same_val[arr[cur]]{
					if !visited[idx]{
						q.PushBack(idx)
						visited[idx] = true
					}
				}
			}
			delete(same_val,arr[cur])
			if (cur - 1) >= 0 && !visited[cur - 1]{
				q.PushBack(cur - 1)
				visited[cur - 1] = true
			}
			if (cur + 1) < l && !visited[cur + 1]{
				q.PushBack(cur + 1)
				visited[cur + 1] = true
			}
		}
		steps++
	}
	return l - 1
}