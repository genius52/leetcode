package diagram

import "container/list"

func NetworkBecomesIdle(edges [][]int, patience []int) int {
	var l int = len(patience)
	var graph [][]int = make([][]int,l)
	for _,edge := range edges{
		graph[edge[0]] = append(graph[edge[0]],edge[1])
		graph[edge[1]] = append(graph[edge[1]],edge[0])
	}
	var dis []int = make([]int,l)
	for i := 0;i < l;i++{
		dis[i] = 2147483647
	}
	var q list.List
	q.PushBack(0)
	dis[0] = 0
	var steps int = 1
	for q.Len() > 0{
		var cur_len int = q.Len()
		for i := 0;i < cur_len;i++{
			cur := q.Front().Value.(int)
			q.Remove(q.Front())
			for j := 0;j < len(graph[cur]);j++{
				if dis[graph[cur][j]] == 2147483647{
					q.PushBack(graph[cur][j])
					dis[graph[cur][j]] = steps
				}
			}
		}
		steps++
	}
	var max_second int = 0
	for i := 1;i < l;i++{
		round_duration := dis[i] * 2
		if patience[i] < round_duration{
			send_times := round_duration / patience[i]
			if round_duration % patience[i] != 0{
				send_times++
			}
			cur_duration := (send_times - 1) * patience[i] + round_duration //last send go duraion + back duration
			round_duration = cur_duration
		}
		if round_duration > max_second{
			max_second = round_duration
		}
	}
	return max_second + 1
}