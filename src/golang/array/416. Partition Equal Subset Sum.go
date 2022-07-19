package array

func dp_canPartition(nums []int,cur_sum int,cur_pos int,target int,memo []map[int]bool)bool{
	if cur_pos >= len(nums) || cur_sum > target{
		return false
	}
	if _,ok := memo[cur_pos][cur_sum];ok{
		return false
	}
	if cur_sum == target{
		return true
	}
	res := dp_canPartition(nums,cur_sum + nums[cur_pos],cur_pos + 1,target,memo) || dp_canPartition(nums,cur_sum,cur_pos + 1,target,memo)
	memo[cur_pos][cur_sum] = res
	return res
}

func CanPartition2(nums []int) bool {
	var sum int = 0
	for _,v := range nums{
		sum += v
	}
	if sum % 2 == 1{
		return false
	}
	var l int = len(nums)
	var memo []map[int]bool = make([]map[int]bool,l)
	for i := 0;i < l;i++{
		memo[i] = make(map[int]bool)
	}
	return dp_canPartition(nums,0,0,sum / 2,memo)
}

//func CanPartition2(nums []int) bool{
//	var l int = len(nums)
//	var sum int = 0
//	for i := 0;i < l;i++{
//		sum += nums[i]
//	}
//	if sum % 2 == 1{
//		return false
//	}
//	var dp [][]bool = make([][]bool,l + 1)//dp[i][j] = true,sum to j is available from [nums[0] ..nums[i])
//	for i := 0;i <= l;i++{
//		dp[i] = make([]bool,sum + 1)
//	}
//	var target int = sum/2
//	for i := 0;i <= l;i++{
//		dp[i][0] = true
//	}
//	for i := 1;i < l;i++{
//		for j := 1;j <= sum;j++{
//			dp[i][j] = dp[i - 1][j]
//			if j >= nums[i - 1]{
//				dp[i][j] = dp[i][j] || dp[i - 1][j - nums[i - 1]]
//			}
//		}
//	}
//	for i := 0;i < l;i++{
//		if dp[i][target]{
//			return true
//		}
//	}
//	return false
//}