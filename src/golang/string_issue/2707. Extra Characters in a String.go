package string_issue

func dfs_minExtraChar(s string, l int, idx int, words map[string]bool, memo []int) int {
	if idx >= l {
		return 0
	}
	if memo[idx] != -1 {
		return memo[idx]
	}
	var res int = l - idx
	for i := idx + 1; i <= l; i++ {
		if _, ok := words[s[idx:i]]; ok {
			cur := dfs_minExtraChar(s, l, i, words, memo)
			if cur < res {
				res = cur
			}
		} else {
			cur := i - idx + dfs_minExtraChar(s, l, i, words, memo)
			if cur < res {
				res = cur
			}
		}
	}
	memo[idx] = res
	return res
}

func MinExtraChar(s string, dictionary []string) int {
	var l int = len(s)
	var words map[string]bool = make(map[string]bool)
	for _, d := range dictionary {
		words[d] = true
	}
	var memo []int = make([]int, l)
	for i := 0; i < l; i++ {
		memo[i] = -1
	}
	return dfs_minExtraChar(s, l, 0, words, memo)
}
