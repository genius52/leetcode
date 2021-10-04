package diagram

import "sort"

func max_int64(a,b int64)int64{
	if a > b {
		return a
	}else{
		return b
	}
}

func MaxTaxiEarnings(n int, rides [][]int) int64 {
	var dp []int64 = make([]int64,n + 1)//dp[i] 从1到i的最大收益
	sort.Slice(rides, func(i, j int) bool {
		if rides[i][1] == rides[j][1]{
			return rides[i][0] < rides[j][0]
		}
		return rides[i][1] < rides[j][1]
	})
	var l int = len(rides)
	var res int64 = 0
	for i := 0;i < l;i++{
		cur := rides[i][1] - rides[i][0] + rides[i][2]
		var total int64 = int64(cur)
		for j := i - 1;j >= 0;j--{
			if rides[j][1] <= rides[i][0]{
				total = max_int64(total,int64(cur) + dp[rides[j][1]])
				break
			}else{
				total = max_int64(total,dp[rides[j][1]])
			}
		}
		dp[rides[i][1]] = max_int64(dp[rides[i][1]],total)
		res = max_int64(res,dp[rides[i][1]])
	}
	return res
}