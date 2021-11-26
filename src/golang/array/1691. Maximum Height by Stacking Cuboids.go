package array

import "sort"

func MaxHeight(cuboids [][]int) int {
	var l int = len(cuboids)
	for i := 0;i < l;i++{
		sort.Ints(cuboids[i])
	}
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] <= cuboids[j][0] && cuboids[i][1] <= cuboids[j][1] && cuboids[i][2] <= cuboids[j][2]{
			return true
		}
		if cuboids[i][0] >= cuboids[j][0] && cuboids[i][1] >= cuboids[j][1] && cuboids[i][2] >= cuboids[j][2]{
			return false
		}
		if cuboids[i][2] != cuboids[j][2]{
			return cuboids[i][2] < cuboids[j][2]
		}
		if cuboids[i][1] != cuboids[j][1]{
			return cuboids[i][1] < cuboids[j][1]
		}
		if cuboids[i][0] != cuboids[j][0]{
			return cuboids[i][0] < cuboids[j][0]
		}
		return true
	})
	var dp []int = make([]int,l)
	for i := 0;i < l;i++{
		dp[i] = cuboids[i][2]
	}
	var res int = cuboids[0][2]
	for i := 1;i < l;i++{
		for j := i - 1;j >= 0;j--{
			if cuboids[i][0] >= cuboids[j][0] && cuboids[i][1] >= cuboids[j][1] && cuboids[i][2] >= cuboids[j][2]{
				dp[i] = max_int(dp[i],dp[j] + cuboids[i][2])
			}
		}
		res = max_int(res,dp[i])
	}
	return res
}