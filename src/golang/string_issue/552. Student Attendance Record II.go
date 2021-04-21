package string_issue

//Any student is eligible for an attendance award if they meet both of the following criteria:
//
//The student was absent ('A') for strictly fewer than 2 days total.
//The student was never late ('L') for 3 or more consecutive days.

//Input: n = 2
//Output: 8
//Explanation: There are 8 records with length 2 that are eligible for an award:
//"PP", "AP", "PA", "LP", "PL", "AL", "LA", "LL"
//Only "AA" is not eligible because there are 2 absences (there need to be fewer than 2).
//func CheckRecord2{
//
//}

func CheckRecord2(n int) int {
	//dp[i][0]  第i + 1天状态为"L"，符合条件的数量
	//dp[i][1]  第i + 1天状态为"P"，符合条件的数量
	if n == 1{
		return 3
	}
	var dp[][2]int = make([][2]int,n)
	dp[0][0] = 1
	dp[0][1] = 1
	dp[1][0] = dp[0][1] + dp[0][1]
	dp[1][1] = dp[0][0] + dp[0][1]

	for i := 2;i < n;i++{
		dp[i][0] = (dp[i - 1][1] + dp[i - 2][1]) % 1000000007
		dp[i][1] = (dp[i - 1][0] + dp[i - 1][1]) % 1000000007
	}
	var total int = (dp[n - 1][0] + dp[n - 1][1]) % 1000000007
	//insert 'A' at 第i天
	for i := 0;i < n;i++{
		var cnt int = 0
		if i == 0 || i == (n - 1){
			cnt = (dp[n - 2][0] + dp[n - 2][1]) % 1000000007
		}else{
			left := dp[n - i - 2][0] + dp[n - i - 2][1]
			right := dp[i - 1][0] + dp[i - 1][1]
			cnt = left * right % 1000000007
		}
		total = (total + cnt) % 1000000007
	}
	return total
}