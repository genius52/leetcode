package array

//Input: [3,1,2,4]
//Output: 17
//Explanation: Subarrays are [3], [1], [2], [4], [3,1], [1,2], [2,4], [3,1,2], [1,2,4], [3,1,2,4].
//Minimums are 3, 1, 2, 4, 1, 1, 2, 1, 1, 1.  Sum is 17.
func SumSubarrayMins(A []int) int {
	var l int = len(A)
	var dp []int = make([]int,l)
	dp[0] = A[0]
	var res int = A[0]
	for i := 1;i < l;i++{
		j := i - 1
		for j >= 0 && A[i] < A[j]{
			j--
		}
		if j == -1{
			dp[i] = (i + 1) * A[i]
		}else{
			dp[i] = (i - j) * A[i] + dp[j]
		}
		res = (res + dp[i])% 1000000007
	}
	return res
}