package diagram

type trace_count struct {
	ret   bool
	count []int
}

func dfs_largestPathValue(graph [][]int, l int, cur int, colors string, memo map[int]trace_count, depth int) (bool, []int) {
	if _, ok := memo[cur]; ok {
		return memo[cur].ret, memo[cur].count
	}
	if depth >= l {
		return false, []int{}
	}
	var res bool = false
	if len(graph[cur]) == 0 {
		res = true
	}
	var max_count []int = make([]int, 26)
	for i := 0; i < len(graph[cur]); i++ {
		if graph[cur][i] == cur {
			res = false
			break
		}
		ret, cnt := dfs_largestPathValue(graph, l, graph[cur][i], colors, memo, depth+1)
		if ret {
			for j := 0; j < 26; j++ {
				if cnt[j] > max_count[j] {
					max_count[j] = cnt[j]
				}
			}
			res = true
		} else {
			res = false
			break
		}
	}
	if res {
		max_count[colors[cur]-'a']++
	}
	var obj trace_count
	obj.ret = res
	obj.count = max_count
	memo[cur] = obj
	return res, max_count
}

func LargestPathValue(colors string, edges [][]int) int {
	var l int = len(colors)
	var edge_len int = len(edges)
	if edge_len == 0 {
		return 1
	}
	var graph [][]int = make([][]int, l)
	var exist []bool = make([]bool, l)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		exist[edge[0]] = true
		exist[edge[1]] = true
	}
	var memo map[int]trace_count = make(map[int]trace_count)
	var max_cnt int = -1
	for i := 0; i < l; i++ {
		if !exist[i] {
			continue
		}
		ret, count := dfs_largestPathValue(graph, l, i, colors, memo, 0)
		if ret {
			for j := 0; j < 26; j++ {
				if count[j] > max_cnt {
					max_cnt = count[j]
				}
			}
		} else {
			return -1
		}
	}
	return max_cnt
}

//Toplogic
//type node_count struct {
//	node  int
//	count []int
//}

//func largestPathValue(colors string, edges [][]int) int {
//	var l int = len(colors)
//	var indegree []int = make([]int, l)
//	var graph [][]int = make([][]int, l)
//	for _, edge := range edges {
//		graph[edge[0]] = append(graph[edge[0]], edge[1])
//		indegree[edge[1]]++
//	}
//	var q list.List
//	for i := 0; i < l; i++ {
//		if indegree[i] == 0 {
//			var obj node_count
//			obj.node = i
//			obj.count = make([]int, l)
//			q.PushBack(i)
//		}
//	}
//	var visited []bool = make([]bool, l)
//	for q.Len() > 0 {
//		var q_len int = q.Len()
//		for i := 0; i < q_len; i++ {
//			cur := q.Front().Value.(node_count)
//			q.Remove(q.Front())
//			for j := 0; j < len(graph[cur.node]); j++ {
//
//			}
//		}
//	}
//
//}
