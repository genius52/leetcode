package diagram

//Input: n = 4, k = 2
//Output: 5
//Explanation:
//The two line segments are shown in red and blue.
//The image above shows the 5 different ways {(0,2),(2,3)}, {(0,1),(1,3)}, {(0,1),(2,3)}, {(1,2),(2,3)}, {(0,1),(1,2)}.
func dfs_numberOfSets(n,k int,memo [][]int)int{
	if n <= 1 || k == 0{
		return 0
	}
	if k == 1{
		res := n * (n - 1)/2
		return res
	}
	if memo[n][k] != 0{
		return memo[n][k]
	}
	var res int = 0
	for i := 1;i <= n - k;i++{//i = used points
		res += i * dfs_numberOfSets(n - i,k - 1,memo)//start from current point with length = i
		res = res % 1000000007
	}
	memo[n][k] = res
	return res
}

func NumberOfSets(n int, k int) int {
	var memo [][]int = make([][]int,1001)
	for i := 0;i < 1001;i++{
		memo[i] = make([]int,1001)
	}
	return dfs_numberOfSets(n,k,memo)
}

//func numberOfSets(n int, k int) int{
//	var dp [][]int = make([][]int,n + 1)//dp[i][j]: 使用i个点，有j个线段的方案数
//	for i := 0;i <= n;i++{
//		dp[i] = make([]int,k + 1)
//	}
//	for i := 2;i <= n;i++{ //i个点，一个线段的方案数
//		dp[i][1] = i * (i - 1)/2
//	}
//	var MOD int = 1e9 + 7
//	for i := 3;i <= n;i++{
//		for j := 2;j < i && j < k;j++{
//			dp[i][j] =
//		}
//	}
//	return dp[n][k]
//}