package number

//Input: days = [1,4,6,7,8,20], costs = [2,7,15]
//Output: 11
//Explanation: For example, here is one way to buy passes that lets you travel your travel plan:
//On day 1, you bought a 1-day pass for costs[0] = $2, which covered day 1.
//On day 3, you bought a 7-day pass for costs[1] = $7, which covered days 3, 4, ..., 9.
//On day 20, you bought a 1-day pass for costs[0] = $2, which covered day 20.
//In total, you spent $11 and covered all the days of your travel.
func mincostTickets(days []int, costs []int) int {
	var l int = len(days)
	var dp []int = make([]int,days[l - 1] + 1)//dp[i] = min cost on 'i'th day
	var record map[int]bool = make(map[int]bool)
	for _,d := range days{
		record[d] = true
	}
	for i := 1;i <= days[l - 1];i++{
		if _,ok := record[i];!ok{
			dp[i] = dp[i - 1]
		}else{
			dp[i] = costs[0] + dp[i - 1]
			if i >= 7{
				dp[i] = min_int(dp[i],costs[1] + dp[i - 7])
			}else{
				dp[i] = min_int(dp[i],costs[1] + dp[0])
			}
			if i >= 30{
				dp[i] = min_int(dp[i],costs[2] + dp[i - 30])
			}else{
				dp[i] = min_int(dp[i],costs[2] + dp[0])
			}
		}
	}
	return dp[days[l - 1]]
}