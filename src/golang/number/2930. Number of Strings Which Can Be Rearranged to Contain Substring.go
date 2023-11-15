package number

func dfs_stringCount(n int, l_cnt int, e_cnt int, t_cnt int, memo map[int]map[int]int) int {
	if n == 0 {
		if l_cnt <= 0 && e_cnt <= 0 && t_cnt <= 0 {
			return 1
		} else {
			return 0
		}
	}
	var res int = 0
	var MOD int = 1e9 + 7
	key := l_cnt*100 + e_cnt*10 + t_cnt
	if _, ok1 := memo[n]; ok1 {
		if _, ok2 := memo[n][key]; ok2 {
			return memo[n][key]
		}
	} else {
		memo[n] = make(map[int]int)
	}
	var meet int = 0
	if l_cnt > 0 {
		res += dfs_stringCount(n-1, l_cnt-1, e_cnt, t_cnt, memo)
		res %= MOD
		meet++
	}
	if e_cnt > 0 {
		res += dfs_stringCount(n-1, l_cnt, e_cnt-1, t_cnt, memo)
		res %= MOD
		meet++
	}
	if t_cnt > 0 {
		res += dfs_stringCount(n-1, l_cnt, e_cnt, t_cnt-1, memo)
		res %= MOD
		meet++
	}
	res += (26 - meet) * dfs_stringCount(n-1, l_cnt, e_cnt, t_cnt, memo)
	res %= MOD
	memo[n][key] = res
	return res
}

func StringCount(n int) int {
	var memo map[int]map[int]int = make(map[int]map[int]int)
	return dfs_stringCount(n, 1, 2, 1, memo)
}
