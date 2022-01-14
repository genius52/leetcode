package number

//TLE by basic dp solution
func dp_stoneGameVIII(prefix []int, l int, pos int, memo []int) int {
	if pos > l {
		return 0
	}
	if pos == l {
		return prefix[pos]
	}
	if memo[pos] != -2147483648 {
		return memo[pos]
	}
	var res int = prefix[l]
	for i := pos; i <= l; i++ {
		//alice和bob都想最大化「自己得分-对方得分」
		//我的得分 - 对手的得分
		cur_diff := prefix[i] - dp_stoneGameVIII(prefix, l, i+1, memo)
		res = max_int(res, cur_diff)
	}
	memo[pos] = res
	return res
}

func stoneGameVIII(stones []int) int {
	var l int = len(stones)
	var prefix []int = make([]int, l+1)
	for i := 0; i < l; i++ {
		prefix[i+1] = prefix[i] + stones[i]
	}
	var memo []int = make([]int, l+1)
	for i := 0; i <= l; i++ {
		memo[i] = -2147483648
	}
	return dp_stoneGameVIII(prefix, l, 2, memo)
}
