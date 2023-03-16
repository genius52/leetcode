package diagram

// return: with leaf && no leaf
func dfs_maxOutput(cur_node int, pre_node int, graph [][]int, n int, prices []int, res *int64) (int64, int64) {
	var cur_with_leaf int64 = int64(prices[cur_node])
	var cur_no_leaf int64 = 0
	for _, next := range graph[cur_node] {
		if next != pre_node {
			next_with_leaf, next_no_leaf := dfs_maxOutput(next, cur_node, graph, n, prices, res)
			*res = max_int64(*res, cur_no_leaf+next_with_leaf)
			*res = max_int64(*res, cur_with_leaf+next_no_leaf)
			cur_with_leaf = max_int64(cur_with_leaf, next_with_leaf+int64(prices[cur_node]))
			cur_no_leaf = max_int64(cur_no_leaf, next_no_leaf+int64(prices[cur_node]))
		}
	}
	return cur_with_leaf, cur_no_leaf
}

func MaxOutput(n int, edges [][]int, price []int) int64 {
	var graph [][]int = make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	var res int64 = 0
	dfs_maxOutput(0, -1, graph, n, price, &res)
	return res
}

//func dfs_maxOutput(cur_node int, pre_node int, n int, price []int, graph [][]int, memo []map[int]int64) int64 {
//	var res int64 = 0
//	for _, next := range graph[cur_node] {
//		if next != pre_node {
//			if val, ok := memo[cur_node][next]; ok {
//				if val > res {
//					res = val
//				}
//			} else {
//				cur := dfs_maxOutput(next, cur_node, n, price, graph, memo)
//				if cur > res {
//					res = cur
//				}
//				memo[cur_node][next] = cur
//			}
//		}
//	}
//	return res + int64(price[cur_node])
//}
//
//func MaxOutput(n int, edges [][]int, price []int) int64 {
//	var memo []map[int]int64 = make([]map[int]int64, n) //memo[i][j]: 从i往j的最大值
//	for i := 0; i < n; i++ {
//		memo[i] = make(map[int]int64)
//	}
//	var graph [][]int = make([][]int, n)
//	for _, edge := range edges {
//		graph[edge[0]] = append(graph[edge[0]], edge[1])
//		graph[edge[1]] = append(graph[edge[1]], edge[0])
//	}
//	var leaves []int
//	for i := 0; i < n; i++ {
//		if len(graph[i]) == 1 {
//			leaves = append(leaves, i)
//		}
//	}
//	var res int64 = 0
//	for _, leaf := range leaves {
//		cur := dfs_maxOutput(leaf, -1, n, price, graph, memo) - int64(price[leaf])
//		if cur > res {
//			res = cur
//		}
//	}
//	return res
//}
