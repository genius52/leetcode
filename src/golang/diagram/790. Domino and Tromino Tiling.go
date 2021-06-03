package diagram

//Input: 3
//Output: 5
//Explanation:
//The five different ways are listed below, different letters indicates different tiles:
//XYZ XXZ XYY XXY XYY
//XYZ YYZ XZZ XYY XXY
func NumTilings(N int) int {
	if N == 1{
		return 1
	}else if N == 2{
		return 2
	}else if N == 3{
		return 5
	}
	var dp []int = make([]int,N + 1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5
	for i := 4;i <= N;i++{
		dp[i] = 2 * dp[i - 1] + dp[i - 3]
		dp[i] = dp[i] % 1000000007
	}
	return dp[N]
}