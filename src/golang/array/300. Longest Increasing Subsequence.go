package array

func LengthOfLIS(nums []int) int{
	var l int = len(nums)
	if l == 0{
		return 0
	}
	var tail []int//tail[i] 表示：长度为 i + 1 的 所有 上升子序列的结尾的最小值
	for i := 0;i < l;i++{
		if len(tail) == 0 || nums[i] > tail[len(tail) - 1]{
			tail = append(tail,nums[i])
		}else{
			var left int = 0
			var right int = len(tail) - 1
			var res int = 0 //找到第一个大于等于nums[i]的
			for left < right{
				mid := (left + right)/2
				if nums[i] == tail[mid]{
					res = mid
					break
				}else if nums[i] > tail[mid]{
					left = mid + 1
					res = left
				}else if nums[i] < tail[mid]{
					right = mid
					res = right
				}
			}
			tail[res] = nums[i]
		}
	}
	return len(tail)
}

//func LengthOfLIS(nums []int) int {
//	if len(nums) == 0{
//		return 0
//	}
//	var dp []int = make([]int,len(nums)+1)
//	dp[0] = 1
//	var res int = 1
//	for i := 1;i < len(nums);i++{
//		max := 1
//		for j := 0;j < i;j++{
//			if nums[i] > nums[j]{
//				if (dp[j] + 1) > max{
//					max = dp[j] + 1
//				}
//			}
//		}
//		if max > res{
//			res = max
//		}
//		dp[i] = max
//	}
//	return res
//}