package string_issue

import "strconv"

func dp_minimumPartition(s string,l int,idx int,k int,k_len int,memo []int)int  {
	if idx == l{
		return 0
	}
	if memo[idx] != -2{
		return memo[idx]
	}
	var min_cnt int = -1
	for i := 1;idx + i <= l && i <= k_len;i++{
		n,_ := strconv.Atoi(s[idx:idx + i])
		if n <= k{
			cur := 1 + dp_minimumPartition(s,l,idx + i,k,k_len,memo)
			if cur != 0{
				if min_cnt == -1 || cur < min_cnt{
					min_cnt = cur
				}
			}
		}
	}
	memo[idx] = min_cnt
	return memo[idx]
}

func MinimumPartition(s string, k int) int {
	var l int = len(s)
	var memo []int = make([]int,l)
	for i := 0;i < l;i++{
		memo[i] = -2
	}
	var k_len int = 0
	var n int = k
	for n > 0{
		n /= 10
		k_len++
	}
	return dp_minimumPartition(s,l,0,k,k_len,memo)
}