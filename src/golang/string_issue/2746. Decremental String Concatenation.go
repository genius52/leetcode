package string_issue

func dfs_minimizeConcatenatedLength(words []string, l int, idx int, first uint8, end uint8, memo [][26][26]int) int {
	if idx == l {
		return 0
	}
	if memo[idx][first-'a'][end-'a'] != -1 {
		return memo[idx][first-'a'][end-'a']
	}
	var l1 int = len(words[idx])
	var l2 int = len(words[idx])
	if words[idx][len(words[idx])-1] == first {
		l1--
	}
	l1 += dfs_minimizeConcatenatedLength(words, l, idx+1, words[idx][0], end, memo)
	if words[idx][0] == end {
		l2--
	}
	l2 += dfs_minimizeConcatenatedLength(words, l, idx+1, first, words[idx][len(words[idx])-1], memo)
	memo[idx][first-'a'][end-'a'] = min_int(l1, l2)
	return memo[idx][first-'a'][end-'a']
}

func MinimizeConcatenatedLength(words []string) int {
	var l int = len(words)
	var memo [][26][26]int = make([][26][26]int, l)
	for i := 0; i < l; i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 26; k++ {
				memo[i][j][k] = -1
			}
		}
	}
	return len(words[0]) + dfs_minimizeConcatenatedLength(words, l, 1, words[0][0], words[0][len(words[0])-1], memo)
}
