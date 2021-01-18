package array

//Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.
func MaxSubArray(nums []int) int{
	var l int = len(nums)
	var dp []int = make([]int,l)
	dp[0] = nums[0]
	var max_val int = nums[0]
	for i := 1;i < l;i++{
		dp[i] = max_int(nums[i],dp[i - 1] + nums[i])
		max_val = max_int(max_val,dp[i])
	}
	return max_val
}