package tree

func dfs_minimumTotalPrice(n int, cur int, pre int, dst int, graph []map[int]bool, node_visit_cnt []int) bool {
	if cur == dst {
		node_visit_cnt[cur]++
		return true
	}
	var ret bool = false
	for next, _ := range graph[cur] {
		if next == pre {
			continue
		}
		r := dfs_minimumTotalPrice(n, next, cur, dst, graph, node_visit_cnt)
		if r {
			node_visit_cnt[cur]++
			ret = true
		}
	}
	return ret
}

func dp_minimumTotalPrice(n int, cur int, graph []map[int]bool, prices []int, node_visit_cnt []int, parent int, can_half bool, memo [][2]int) int {
	if cur == n {
		return 0
	}
	if can_half {
		if memo[cur][1] != -1 {
			return memo[cur][1]
		}
	} else {
		if memo[cur][0] != -1 {
			return memo[cur][0]
		}
	}
	var res int = 1<<31 - 1
	var half_rest int = 0
	var nothalf_rest int = 0
	for next, _ := range graph[cur] {
		if next == parent {
			continue
		}
		if can_half {
			nothalf_rest += dp_minimumTotalPrice(n, next, graph, prices, node_visit_cnt, cur, true, memo)
			half_rest += dp_minimumTotalPrice(n, next, graph, prices, node_visit_cnt, cur, false, memo)
			//res = min_int(res, prices[cur]*node_visit_cnt[cur]+dp_minimumTotalPrice(n, next, graph, prices, node_visit_cnt, cur, true, memo))
			//res = min_int(res, prices[cur]*node_visit_cnt[cur]/2+dp_minimumTotalPrice(n, next, graph, prices, node_visit_cnt, cur, false, memo))
		} else {
			//res = min_int(res, prices[cur]*node_visit_cnt[cur]+dp_minimumTotalPrice(n, next, graph, prices, node_visit_cnt, cur, true, memo))
			nothalf_rest += dp_minimumTotalPrice(n, next, graph, prices, node_visit_cnt, cur, true, memo)
		}
	}
	if nothalf_rest == 0 { //last node
		if can_half {
			res = prices[cur] * node_visit_cnt[cur] / 2
		} else {
			res = prices[cur] * node_visit_cnt[cur]
		}
	} else {
		if can_half {
			res = min_int(prices[cur]*node_visit_cnt[cur]/2+half_rest, prices[cur]*node_visit_cnt[cur]+nothalf_rest)
		} else {
			res = min_int(res, prices[cur]*node_visit_cnt[cur]+nothalf_rest)
		}
	}
	if can_half {
		memo[cur][1] = res
	} else {
		memo[cur][0] = res
	}
	return res
}

func MinimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	var graph []map[int]bool = make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		graph[edge[0]][edge[1]] = true
		graph[edge[1]][edge[0]] = true
	}
	var node_visit_cnt []int = make([]int, n)
	for _, trip := range trips {
		start := trip[0]
		end := trip[1]
		dfs_minimumTotalPrice(n, start, -1, end, graph, node_visit_cnt)
	}
	var memo [][2]int = make([][2]int, n)
	for i := 0; i < n; i++ {
		memo[i][0] = -1
		memo[i][1] = -1
	}
	res := dp_minimumTotalPrice(n, 0, graph, price, node_visit_cnt, -1, true, memo)
	return res
}
