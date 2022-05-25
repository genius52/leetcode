package string_issue

func dp_findMaxForm(strs []string,l int,idx int,m int,n int,memo [][][]int)int{
	if idx >= l{
		return 0
	}
	if m < 0 || n < 0{
		return 0
	}
	if memo[idx][m][n] != -1{
		return memo[idx][m][n]
	}
	var res int = 0
	var zero_cnt int = 0
	var one_cnt int = 0
	for _,c := range strs[idx]{
		if c == '0'{
			zero_cnt++
		}else if c == '1'{
			one_cnt++
		}
	}
	if zero_cnt <= m && one_cnt <= n{
		res = 1 + dp_findMaxForm(strs,l,idx + 1,m - zero_cnt,n - one_cnt,memo)
	}
	skip := dp_findMaxForm(strs,l,idx + 1,m,n,memo)
	if skip > res{
		res = skip
	}
	memo[idx][m][n] = res
	return res
}

func findMaxForm(strs []string, m int, n int) int {
	var l int = len(strs)
	var memo [][][]int = make([][][]int,l)
	for i := 0;i < l;i++{
		memo[i] = make([][]int,m + 1)
		for j := 0;j <= m;j++{
			memo[i][j] = make([]int,n + 1)
			for k := 0;k <= n;k++{
				memo[i][j][k] = -1
			}
		}
	}
	return dp_findMaxForm(strs,l,0,m,n,memo)
}