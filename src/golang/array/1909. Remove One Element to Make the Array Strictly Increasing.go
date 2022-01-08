package array

func CanBeIncreasing(nums []int) bool {
	var l int = len(nums)
	var dp []int = make([]int, l)
	var res int = 0
	for i := 0; i < l; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max_int(dp[i], 1+dp[j])
			}
		}
		res = max_int(res, dp[i])
	}
	return res >= l-1
}
