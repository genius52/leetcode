package string_issue

func dp_numWays(record [][26]int,word_len int,pos1 int,target string,target_len int,pos2 int,memo [][]int)int{
	if pos2 >= target_len{
		return 1
	}
	if pos1 >= word_len {
		return 0
	}
	if memo[pos1][pos2] != -1{
		return memo[pos1][pos2]
	}
	var MOD int = 1e9 + 7
	var res int = 0
	for i := pos1;i < word_len;i++ {
		cnt := record[i][target[pos2] - 'a']
		if cnt > 0 {
			cnt *= dp_numWays(record,word_len,i + 1,target,target_len,pos2 + 1,memo)
			res += cnt
			res %= MOD
		}
	}
	memo[pos1][pos2] = res
	return res
}

func NumWays1639(words []string, target string) int {
	var word_len int = len(words[0])
	var target_len int = len(target)
	var record [][26]int = make([][26]int,word_len)
	for i := 0;i < word_len;i++{
		for _,word := range words{
			record[i][word[i] - 'a']++
		}
	}
	var memo [][]int = make([][]int,word_len)
	for i := 0;i < word_len;i++{
		memo[i] = make([]int,target_len)
		for j := 0;j < target_len;j++{
			memo[i][j] = -1
		}
	}
	return dp_numWays(record,word_len,0,target,target_len,0,memo)
}