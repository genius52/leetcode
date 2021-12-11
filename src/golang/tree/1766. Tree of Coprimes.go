package tree

func dfs_getCoprimes(nums []int, cur_idx int, pre_idx int, depth int, graph [][]int, prime_memo [51][]int, num_idx [51]int, idx_depth []int, res []int) {
	var prime_idx int = -1
	var most_depth int = -1
	idx_depth[cur_idx] = depth
	for _, prime_num := range prime_memo[nums[cur_idx]] { //当前节点数字的互质数，找到深度最大的那个
		if num_idx[prime_num] != -1 {
			if idx_depth[num_idx[prime_num]] > most_depth {
				most_depth = idx_depth[num_idx[prime_num]]
				prime_idx = num_idx[prime_num]
			}
		}
	}
	res[cur_idx] = prime_idx
	old := num_idx[nums[cur_idx]]
	num_idx[nums[cur_idx]] = cur_idx
	//num_depth[nums[cur_idx]] = Depth_idx{depth, cur_idx}
	for _, neighbour_idx := range graph[cur_idx] {
		if neighbour_idx != pre_idx {
			dfs_getCoprimes(nums, neighbour_idx, cur_idx, depth+1, graph, prime_memo, num_idx, idx_depth, res)
		}
	}
	num_idx[nums[cur_idx]] = old
}

func GetCoprimes(nums []int, edges [][]int) []int {
	var prime_memo [51][]int
	for i := 1; i <= 50; i++ {
		for j := 1; j <= 50; j++ {
			if gcd(j, i) == 1 {
				prime_memo[i] = append(prime_memo[i], j)
				prime_memo[j] = append(prime_memo[j], i)
			}
		}
	}
	var l int = len(nums)
	var graph [][]int = make([][]int, l)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	// 祖先出现的最近深度和索引，不同索引数字可能相同，这个数字可能出现在同一深度
	var idx_depth []int = make([]int, l)
	var num_idx [51]int //num:idx
	for i := 1; i <= 50; i++ {
		num_idx[i] = -1
	}
	var res []int = make([]int, l)
	dfs_getCoprimes(nums, 0, -1, 0, graph, prime_memo, num_idx, idx_depth, res)
	return res
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
