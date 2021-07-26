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
	var last int = days[l - 1]
	var dp []int = make([]int,last + 1)//dp[i] = min cost on 'i'th day
	for i := 1;i <= last;i++{

	}
	return dp[last]
}