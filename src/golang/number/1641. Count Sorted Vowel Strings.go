package number


//将 n 个小球放到 5 个盒子里，盒子可以为空
func countVowelStrings(n int) int{
	var cnt [5]int = [5]int{1,1,1,1,1}
	for i := 1;i < n;i++{
		var cur_cnt [5]int
		for j := 0;j < 5;j++{
			cur := 0
			for k := 0;k <= j;k++{
				cur += cnt[k]
			}
			cur_cnt[j] = cur
		}
		cnt[0],cnt[1],cnt[2],cnt[3],cnt[4] = cur_cnt[0],cur_cnt[1],cur_cnt[2],cur_cnt[3],cur_cnt[4]
	}
	return cnt[0] + cnt[1] + cnt[2] + cnt[3] + cnt[4]
}

func CountVowelStrings(n int) int {
	var dp [][]int = make([][]int,n)
	for i := 0;i < n;i++{
		dp[i] = make([]int,5)
	}
	dp[0][0] = 1//a
	dp[0][1] = 1//e
	dp[0][2] = 1//i
	dp[0][3] = 1//o
	dp[0][4] = 1//u
	for i := 1;i < n;i++{
		for j := 0;j < 5;j++{
			for k := 0;k <= j;k++{
				dp[i][j] += dp[i - 1][k]
			}
		}
	}
	return dp[n - 1][0] + dp[n - 1][1] + dp[n - 1][2] + dp[n - 1][3] + dp[n - 1][4]
}