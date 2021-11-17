package array

import "sort"

func bestTeamScore(scores []int, ages []int) int {
	var l int = len(scores)
	var record [][2]int = make([][2]int,l)
	for i := 0;i < l;i++{
		record[i] = [2]int{ages[i],scores[i]}
	}
	sort.Slice(record, func(i, j int) bool {
		if record[i][0] == record[j][0]{
			return record[i][1] < record[j][1]
		}else{
			return record[i][0] < record[j][0]
		}
	})
	var res int = 0
	var dp []int = make([]int,l)
	for i := 0;i < l;i++{
		dp[i] = record[i][1]
		for j := i - 1;j >= 0;j--{
			if record[j][1] <= record[i][1]{ //年龄小的，得分也较少
				dp[i] = max_int(dp[i],dp[j] + record[i][1])
			}
		}
		res = max_int(res,dp[i])
	}
	return res
}