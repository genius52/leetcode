package number

func dp_maxValueOfCoins(prefix_sum [][]int,l int, idx int,k int,memo [][]int) int{
	if idx >= l || k == 0{
		return 0
	}
	if memo[idx][k] != 0{
		return memo[idx][k]
	}
	var res int = dp_maxValueOfCoins(prefix_sum,l,idx + 1,k,memo)
	for i := 0;i < len(prefix_sum[idx]) && i < k;i++{
		cur := prefix_sum[idx][i] + dp_maxValueOfCoins(prefix_sum,l,idx + 1,k - i - 1,memo)
		if cur > res{
			res = cur
		}
	}
	memo[idx][k] = res
	return res
}

func MaxValueOfCoins(piles [][]int, k int) int {
	var l int = len(piles)
	var prefix_sum [][]int = make([][]int,l)
	for i := 0;i < l;i++{
		prefix_sum[i] = make([]int,len(piles[i]))
		for j := 0;j < len(piles[i]);j++{
			if j == 0{
				prefix_sum[i][j] = piles[i][j]
			}else{
				prefix_sum[i][j] = piles[i][j] + prefix_sum[i][j - 1]
			}
		}
	}
	var memo [][]int = make([][]int,l + 1)
	for i := 0;i <= l;i++{
		memo[i] = make([]int,k + 1)
	}
	return dp_maxValueOfCoins(prefix_sum,l,0,k,memo)
}