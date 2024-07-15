package array

func countSubarrays(nums []int, k int) int64 {
	var l int = len(nums)
	var res int64 = 0
	//var dp []int = make([]int, l) //dp[i]: 以nums[i]结尾的子数组数量
	for i := 0; i < l; i++ {
		var sum int = nums[i]
		for j := i - 1; j >= 0; j-- {
			sum &= nums[j]
			if sum == k {
				res++
			}
		}
	}
	return res
}
