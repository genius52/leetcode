package diagram
//Input: n = 1
//Output: 12

func NumOfWays(n int) int {
	var diff3 []int = make([]int,n)
	var same2 []int = make([]int,n)
	diff3[0] = 6
	same2[0] = 6
	for i := 1;i < n;i++{
		diff3[i] = 2 * diff3[i-1] + 2 * same2[i-1]
		same2[i] = 2 * diff3[i - 1] + 3 * same2[i - 1]
	}
	return diff3[n - 1] + same2[n - 1]
	//for i := 0;i < n;i++{
	//	dp[i] = make([]int,3)
	//	if i == 0{
	//		dp[i][0] = 3
	//	}else{
	//		dp[i][0] = 2
	//	}
	//}
	//dp[0][1] = 2
	//dp[0][2] = 2
	//var res int = 0

}
