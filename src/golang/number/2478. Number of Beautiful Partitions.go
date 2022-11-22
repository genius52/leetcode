package number


func dfs_beautifulPartitions(s string,l int,idx int,k int,minLength int,memo [][]int)int{
	if idx == l{
		if k == 0{
			return 1
		}
		return 0
	}
	if k == 0{
		return 0
	}
	if memo[idx][k] != -1{
		return memo[idx][k]
	}
	if s[idx] != '2' && s[idx] != '3' && s[idx] != '5'&& s[idx] != '7'{
		return 0
	}
	var res int = 0
	var MOD int = 1e9 + 7
	for i := idx + minLength - 1;i < l;i++{
		if s[i] == '2' || s[i] == '3' || s[i] == '5' || s[i] == '7'{
			continue
		}
		res += dfs_beautifulPartitions(s,l,i + 1,k - 1,minLength,memo)
		res %= MOD
	}
	memo[idx][k] = res
	return res
}

func BeautifulPartitions(s string, k int, minLength int) int {
	var l int = len(s)
	var memo [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		memo[i] = make([]int,k + 1)
		for j := 0;j <= k;j++{
			memo[i][j] = -1
		}
	}
	return dfs_beautifulPartitions(s,l,0,k,minLength,memo)
}