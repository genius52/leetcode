package tree

// left->right->current
func dfs_countPalindromePaths(graph [][]int, s string, cur_node int, cur_status int, paths map[int]int) int64 {
	var res int64 = 0
	for i := 0; i < 26; i++ {
		if cnt, ok := paths[cur_status^(1<<i)]; ok { //寻找已经存在,只差一个字符的
			res += int64(cnt)
		}
	}
	if cnt, ok := paths[cur_status]; ok {
		res += int64(cnt)
	}
	paths[cur_status]++
	for i := 0; i < len(graph[cur_node]); i++ {
		child := graph[cur_node][i]
		res += dfs_countPalindromePaths(graph, s, child, cur_status^(1<<int(s[child]-'a')), paths)
	}
	return res
}

func CountPalindromePaths(parent []int, s string) int64 {
	var l int = len(parent)
	var graph [][]int = make([][]int, l)
	for i := 1; i < l; i++ {
		graph[parent[i]] = append(graph[parent[i]], i)
	}
	var paths map[int]int = make(map[int]int)
	return dfs_countPalindromePaths(graph, s, 0, 0, paths)
}
