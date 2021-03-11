package array

func LengthOfLIS(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	var dp []int = make([]int,len(nums)+1)
	dp[0] = 1
	for i := 1;i < len(nums);i++{
		max := 1
		for j := 0;j < i;j++{
			if nums[i] > nums[j]{
				if (dp[j] + 1) > max{
					max = dp[j] + 1
				}
			}
		}
		dp[i] = max
	}
	var res int = 1
	for _,val := range dp{
		if val > res{
			res = val
		}
	}
	return res
}