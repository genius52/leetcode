package array

import "math"

//Input:
//A: [1,2,3,2,1]
//B: [3,2,1,4,7]
//Output: 3
//Explanation:
//The repeated subarray with maximum length is [3, 2, 1].
func FindLength(A []int, B []int) int {
	var len_a int = len(A)
	var len_b int = len(B)
	var dp [][]int = make([][]int,len_a + 1)//dp[i][j] = longest length from A[0] to A[i]    B[0] to B[j]
	for i := 0;i <= len_a;i++{
		dp[i] = make([]int,len_b + 1)
	}
	var res int = 0
	for i := 0;i < len_a;i++{
		for j := 0;j < len_b;j++{
			if A[i] == B[j]{
				dp[i+1][j+1] = dp[i][j] + 1
				res = int(math.Max(float64(dp[i+1][j+1]),float64(res)))
			}
		}
	}
	return res
}