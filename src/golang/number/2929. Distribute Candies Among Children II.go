package number

func DistributeCandies2929(n int, limit int) int64 {
	//var memo map[int]int
	//return dfs_distributeCandies2929(n, limit)
	var res int64 = 0
	for first := 0; first <= limit && first <= n; first++ {
		rest := n - first
		if rest > limit*2 {
			continue
		}
		if rest <= limit {
			res += int64(rest + 1)
		} else {
			begin := min_int(limit, rest-limit)
			end := max_int(limit, rest-limit)
			res += int64(end - begin + 1)
		}
	}
	return res
}
