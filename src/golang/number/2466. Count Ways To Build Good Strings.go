package number

func dfs_countGoodStrings(cur_len int,low int,high int,zero int,one int,memo map[int]int)int{
	if cur_len > high{
		return 0
	}
	var res int = 0
	if cur_len >= low{
		res++
	}
	if _,ok := memo[cur_len];ok{
		return memo[cur_len]
	}
	var MOD int = 1e9 + 7
	res += dfs_countGoodStrings(cur_len + zero,low,high,zero,one,memo)
	res %= MOD
	res += dfs_countGoodStrings(cur_len + one,low,high,zero,one,memo)
	res %= MOD
	memo[cur_len] = res
	return res
}

func CountGoodStrings(low int, high int, zero int, one int) int {
	var memo map[int]int = make(map[int]int)
	return dfs_countGoodStrings(0,low,high,zero,one,memo)
}