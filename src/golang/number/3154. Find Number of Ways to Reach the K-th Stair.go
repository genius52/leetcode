package number

func dfs_waysToReachStair(target int, cur int, jump_cnt int, pre_decrease bool, memo map[int64]int) int {
	if cur > target+1 {
		return 0
	}
	var status int64 = int64(cur)<<32 | int64(jump_cnt)<<1
	if pre_decrease {
		status |= 1
	}
	if _, ok := memo[status]; ok {
		return memo[status]
	}
	var res int = 0
	if cur == target {
		res = 1
	}
	res += dfs_waysToReachStair(target, i+(1<<jump_cnt), jump_cnt+1, false, memo)
	if cur != 0 && !pre_decrease {
		res += dfs_waysToReachStair(target, cur-1, jump_cnt, true, memo)
	}
	memo[status] = res
	return res
}

func waysToReachStair(k int) int {
	var memo map[int64]int = make(map[int64]int)
	return dfs_waysToReachStair(k, 0, 0, false, memo)
}
