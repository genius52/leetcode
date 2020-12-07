package array

//Input: A = [3,8,1,3,2,1,8,9,0], L = 3, M = 2
//Output: 29
//Explanation: One choice of subarrays is [3,8,1] with length 3, and [8,9] with length 2.
func MaxSumTwoNoOverlap(A []int, L int, M int) int {
	var l int = len(A)
	var dp []int = make([]int,l + 1)
	dp[0] = 0
	for i := 1;i <= l;i++{
		dp[i] = A[i - 1] + dp[i - 1]
	}
	var max_sum int = 0
	for i := 0;i + L <= l;i++{
		sum_l := dp[i + L] - dp[i]
		for j := 0;j + M < i;j++{
			sum_m := dp[j + M] - dp[j]
			val := sum_l + sum_m
			if val > max_sum{
				max_sum = val
			}
		}
		for j := i + L;j + M <= l;j++{
			sum_m := dp[j + M] - dp[j]
			val := sum_l + sum_m
			if val > max_sum{
				max_sum = val
			}
		}
	}
	return max_sum
}