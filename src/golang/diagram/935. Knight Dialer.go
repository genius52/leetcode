package diagram

//Input: n = 2
//Output: 20
//Explanation: All the valid number we can dial are
//[04, 06, 16, 18, 27, 29, 34, 38, 40, 43, 49, 60, 61, 67, 72, 76, 81, 83, 92, 94]
func KnightDialer(n int) int {
	if n == 1{
		return 10
	}
	if n == 2{
		return 20
	}
	var graph map[int][]int = make(map[int][]int)
	graph[0] = []int{4,6}
	graph[1] = []int{6,8}
	graph[2] = []int{7,9}
	graph[3] = []int{4,8}
	graph[4] = []int{0,3,9}
	graph[6] = []int{0,1,7}
	graph[7] = []int{2,6}
	graph[8] = []int{1,3}
	graph[9] = []int{2,4}

	//dp[i][0] = dp[i - 1][4] + dp[i - 1][6]
	//dp[2][0] = dp[1][4] + dp[1][6]
	//04,06
	//040,043,049, 060,061,067
	var dp [][]int = make([][]int,n)//dp[i][j]: when length = i + 1,the number j's counts
	for i := 0;i < n;i++{
		dp[i] = make([]int,10)
	}
	for i := 0;i < 10;i++{
		dp[0][i] = 1
	}
	dp[1][0] = 2
	dp[1][1] = 2
	dp[1][2] = 2
	dp[1][3] = 2
	dp[1][4] = 3
	dp[1][5] = 0
	dp[1][6] = 3
	dp[1][7] = 2
	dp[1][8] = 2
	dp[1][9] = 2
	for i := 2;i < n;i++{
		for j := 0;j < 10;j++{
			if nums,ok := graph[j];ok{
				for _,num := range nums{
					dp[i][j] += dp[i - 1][num]
				}
			}
			dp[i][j] %= 1000000007
		}
	}
	var res int = 0
	for i := 0;i < 10;i++{
		res += dp[n - 1][i]
		res %= 1000000007
	}
	return res
}