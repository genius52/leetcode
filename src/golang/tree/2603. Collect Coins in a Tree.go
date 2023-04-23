package tree

import "container/list"

func CollectTheCoins(coins []int, edges [][]int) int {
	var l1 int = len(coins)
	var graph []map[int]bool = make([]map[int]bool, l1)
	for i := 0; i < l1; i++ {
		graph[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		graph[edge[0]][edge[1]] = true
		graph[edge[1]][edge[0]] = true
	}
	var q1 list.List
	var delete_node map[int]bool = make(map[int]bool)
	for i := 0; i < l1; i++ {
		if len(graph[i]) == 1 && coins[i] == 0 {
			q1.PushBack(i)
		}
	}
	for q1.Len() > 0 {
		var cur_len int = q1.Len()
		for i := 0; i < cur_len; i++ {
			var cur int = q1.Front().Value.(int)
			delete_node[cur] = true
			q1.Remove(q1.Front())
			for next, _ := range graph[cur] {
				delete(graph[next], cur)
				if _, ok := delete_node[next]; ok {
					continue
				}
				if coins[next] == 1 {
					continue
				}
				if len(graph[next]) == 1 {
					q1.PushBack(next)
				}
			}
			graph[cur] = make(map[int]bool)
		}
	}
	var q2 list.List
	for i := 0; i < l1; i++ {
		if len(graph[i]) == 1 {
			q2.PushBack(i)
			//delete_node[i] = true
		}
	}
	var steps int = 2
	for steps > 0 {
		var cur_len int = q2.Len()
		for i := 0; i < cur_len; i++ {
			var cur int = q2.Front().Value.(int)
			q2.Remove(q2.Front())
			delete_node[cur] = true
			for next, _ := range graph[cur] {
				delete(graph[next], cur)
				if _, ok := delete_node[next]; ok {
					continue
				}
				if len(graph[next]) == 1 {
					q2.PushBack(next)
				}
			}
			//graph[cur] = make(map[int]bool)
		}
		steps--
	}

	var res int = 0
	for _, edge := range edges {
		node1 := edge[0]
		node2 := edge[1]
		if _, ok1 := delete_node[node1]; !ok1 {
			if _, ok2 := delete_node[node2]; !ok2 {
				res += 2
			}
		}
	}
	return res
}
