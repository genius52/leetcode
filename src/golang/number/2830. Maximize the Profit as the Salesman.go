package number

import (
	"sort"
)

func dp_maximizeTheProfit(n int, cur int, offers [][]int, l int, idx int, memo []int) int {
	if idx == l {
		return 0
	}
	if cur >= n {
		return 0
	}
	//key := strconv.Itoa(cur) + "," + strconv.Itoa(idx)
	if memo[cur] != -1 {
		return memo[cur]
	}
	start := offers[idx][0]
	end := offers[idx][1]
	price := offers[idx][2]
	//skip current
	var res int = dp_maximizeTheProfit(n, cur, offers, l, idx+1, memo)
	if start >= cur {
		//choose current
		res = max_int(res, price+dp_maximizeTheProfit(n, end+1, offers, l, idx+1, memo))
	}
	memo[cur] = res
	return res
}

func MaximizeTheProfit(n int, offers [][]int) int {
	sort.Slice(offers, func(i, j int) bool {
		return offers[i][0] < offers[j][0]
	})
	var l int = len(offers)
	var memo []int = make([]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = -1
	}
	return dp_maximizeTheProfit(n, 0, offers, l, 0, memo)
}

func maximizeTheProfit(n int, offers [][]int) int {
	var dp []int = make([]int, n+1) //dp[i]:前i - 1间房子最多能卖的价钱
	sort.Slice(offers, func(i, j int) bool {
		if offers[i][1] == offers[j][1] {
			if offers[i][0] == offers[j][0] {
				return offers[i][2] > offers[j][2]
			} else {
				return offers[i][0] < offers[j][0]
			}
		} else {
			return offers[i][1] < offers[j][1]
		}
	})
	var res int = 0
	var l2 int = len(offers)
	var j int = 0
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i]
		for j < l2 && offers[j][1] == i {
			dp[i+1] = max_int(dp[i+1], offers[j][2]+dp[offers[j][0]])
			res = max_int(res, dp[i+1])
			j++
		}
	}
	return dp[n]
}
