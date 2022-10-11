package string_issue

//func deleteString(s string) int {
//	var l int = len(s)
//	var dp []int = make([]int,l)//dp[i]: max steps from 0 to i
//	dp[0] = 1
//	for i := 0;i < l;i++{
//
//	}
//	return dp[l - 1]
//}

func dfs_deleteString(s string,l int,idx int,memo []int)int{
	if idx >= l{
		return 0
	}
	if idx + 1 == l{
		return 1
	}
	if memo[idx] != -1{
		return memo[idx]
	}
	var res int = 1
	for i := (l - idx)/2;i >= 1;i--{
		if s[idx : idx + i] == s[idx + i: idx + i * 2]{
			cur := 1 + dfs_deleteString(s,l,idx + i,memo)
			if cur > res{
				res = cur
			}
		}
	}
	memo[idx] = res
	return res
}

func DeleteString(s string) int {
	var l int = len(s)
	var memo []int = make([]int,l)
	for i := 0;i < l;i++{
		memo[i] = -1
	}
	return dfs_deleteString(s,l,0,memo)
}