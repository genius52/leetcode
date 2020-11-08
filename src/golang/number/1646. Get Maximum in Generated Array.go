package number

//Input: n = 7
//Output: 3
//Explanation: According to the given rules:
//  nums[0] = 0
//  nums[1] = 1
//  nums[(1 * 2) = 2] = nums[1] = 1
//  nums[(1 * 2) + 1 = 3] = nums[1] + nums[2] = 1 + 1 = 2
//  nums[(2 * 2) = 4] = nums[2] = 1
//  nums[(2 * 2) + 1 = 5] = nums[2] + nums[3] = 1 + 2 = 3
//  nums[(3 * 2) = 6] = nums[3] = 2
//  nums[(3 * 2) + 1 = 7] = nums[3] + nums[4] = 2 + 1 = 3
//Hence, nums = [0,1,1,2,1,3,2,3], and the maximum is 3.
func getMaximumGenerated(n int) int {
	if n <= 1{
		return n
	}
	var dp []int = make([]int,n + 1)
	dp[0] = 0
	dp[1] = 1
	var res int = 0
	for i := 2;i <= n;i++{
		if i | 1 == i{
			sub_index := (i - 1)/2
			dp[i] = dp[sub_index] + dp[sub_index + 1]
		}else{
			dp[i] = dp[i / 2]
		}
		if dp[i] > res{
			res = dp[i]
		}
	}
	return res
}