package diagram

import "container/list"

func check_minMutation(s1 string,s2 string)bool{
	var l1 int = len(s1)
	var l2 int = len(s2)
	if l1 != l2{
		return false
	}
	var diff_cnt int = 0
	for i := 0;i < l1;i++{
		if s1[i] != s2[i]{
			diff_cnt++
		}
		if diff_cnt > 1{
			return false
		}
	}
	return true
}

func minMutation(start string, end string, bank []string) int {
	var l int = len(bank)
	var graph map[string][]string = make(map[string][]string)
	for i := 0;i < l;i++{
		for j := 0;j < l;j++{
			if i == j{
				continue
			}
			if check_minMutation(bank[i],bank[j]){
				graph[bank[i]] = append(graph[bank[i]],bank[j])
				graph[bank[j]] = append(graph[bank[j]],bank[i])
			}
		}
	}
	for i := 0;i < l;i++{
		if check_minMutation(start,bank[i]){
			graph[start] = append(graph[start],bank[i])
		}
	}
	var steps int = 1
	var visited map[string]bool = make(map[string]bool)
	var q list.List
	q.PushBack(start)
	visited[start] = true
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			cur := q.Front().Value.(string)
			q.Remove(q.Front())
			for _,next := range graph[cur]{
				if next == end {
					return steps
				}
				if _,ok := visited[next];!ok{
					visited[next] = true
					q.PushBack(next)
				}
			}
		}
		steps++
	}
	return -1
}