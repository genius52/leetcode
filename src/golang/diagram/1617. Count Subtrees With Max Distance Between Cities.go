package diagram

//floyd
func CountSubgraphsForEachDiameter(n int, edges [][]int) []int{
	var graph [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		graph[i] = make([]int,n)
		for j := 0;j < n;j++{
			if i == j{
				graph[i][j] = 0
			}else{
				graph[i][j] = 1 << 31 - 1
			}
		}
	}
	for _,edge := range edges{
		graph[edge[0] - 1][edge[1] - 1] = 1
		graph[edge[1] - 1][edge[0] - 1] = 1
	}
	for mid := 0;mid < n;mid++{
		for i := 0;i < n;i++{
			for j := 0;j < n;j++{
				graph[i][j] = min_int(graph[i][j],graph[i][mid] + graph[mid][j])
			}
		}
	}
	var res []int = make([]int,n - 1)
	for i := 1;i <= (1 << (n + 1) - 1);i++{
		dis := max_distance(i,graph,n)
		if dis >= 1{
			res[dis - 1]++
		}
	}
	return res
}

func max_distance(state int,graph [][]int,n int)int{
	var max_val int =0
	var edge_cnt int = 0
	var a int = state
	var k int = 0
	for a > 0{
		if a | 1 == a{
			k++
		}
		a = a >> 1
	}
	for i := 0;i < n;i++{
		if state | (1 << i) != state{ //i is not in state
			continue
		}
		for j := i + 1;j < n;j++{
			if state | (1 << j) != state{ //i is not in state
				continue
			}
			if graph[i][j] == 1{
				edge_cnt++
			}
			max_val = max_int(max_val,graph[i][j])
		}
	}
	if edge_cnt == k - 1{
		return max_val
	}else{
		return 0
	}
}

//bfs + bitmask
//func CountSubgraphsForEachDiameter(n int, edges [][]int) []int {
//	var graph [][]bool = make([][]bool,n + 1)
//	for i := 0;i <= n;i++{
//		graph[i] = make([]bool,n + 1)
//	}
//	for i := 0;i < len(edges);i++{
//		graph[edges[i][0]][edges[i][1]] = true
//		graph[edges[i][1]][edges[i][0]] = true
//	}
//	var res []int = make([]int,n)
//	var visited map[int]bool = make(map[int]bool)
//
//	var q list.List
//	for i := 1;i <= n;i++{
//		var obj node_collection
//		obj.node = i
//		obj.mask = 1 << i
//		q.PushBack(obj)
//		visited[obj.mask] = true
//	}
//	dis := 0
//	for q.Len() > 0{
//		cur_len := q.Len()
//		if dis > 0{
//			res[dis - 1] = cur_len
//		}
//		for i := 0;i < cur_len;i++{
//			cur := q.Front().Value.(node_collection)
//			q.Remove(q.Front())
//			var near []int
//			for j := 0;j < n;j++{
//				if graph[cur.node][j]{
//					next := cur
//					next.mask = cur.mask | 1 << j
//					next.node = j
//					if _,ok := visited[next.mask];ok{
//						continue
//					}
//					visited[next.mask] = true
//					q.PushBack(next)
//				}
//			}
//			for _,j := range near{//可以任选1 到 n个
//
//			}
//		}
//		dis++
//	}
//	return res
//}