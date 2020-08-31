package array

//Input: [23, 2, 6, 4, 7],  k=6
//Output: True
//Explanation: Because [23, 2, 6, 4, 7] is an continuous subarray of size 5 and sums up to 42.
func CheckSubarraySum(nums []int, k int) bool {
	var l int = len(nums)
	if l <= 1{
		return false
	}
	if k == 0{
		for _,n := range nums{
			if n != 0{
				return false
			}
		}
		return true
	}
	var sum int = 0
	var left_dp []int = make([]int,l + 1) //dp[i] = sum from 0 to i
	var right_dp []int = make([]int,l + 1)
	left_dp[0] = 0
	right_dp[0] = 0
	for i := 1;i <= l;i++{
		sum += nums[i - 1]
		left_dp[i] = left_dp[i - 1] + nums[i - 1]
		if left_dp[i] != 0 && left_dp[i] % k == 0{
			return true
		}
		right_dp[i] = right_dp[i - 1] + nums[l - i]
		if right_dp[i] != 0 && right_dp[i] % k == 0{
			return true
		}
	}
	for left := 0;left <= l;left++{
		for right := left + 1;right <= l;right++{
			val := sum - left_dp[left] - right_dp[right]
			if val != 0 && (val % k) == 0{
				return true
			}
		}
	}
	return false
}