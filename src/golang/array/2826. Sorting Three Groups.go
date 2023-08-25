package array

//func minimumOperations2826(nums []int) int {
//	var l int = len(nums)
//	var dp [][4]int = make([][4]int, l)
//	if nums[0] == 1 {
//		dp[0][2] = 1
//		dp[0][3] = 1
//	} else if nums[0] == 2 {
//		dp[0][1] = 1
//		dp[0][3] = 1
//	} else if nums[0] == 3 {
//		dp[0][1] = 1
//		dp[0][2] = 1
//	}
//	for i := 1; i < l; i++ {
//		if nums[i] == 1 {
//			dp[i][1] = dp[i][1]
//			dp[i][2] = dp[i-1][2] + 1
//			dp[i][3] = dp[i-1][3] + 1
//		} else if nums[i] == 2 {
//
//		} else if nums[i] == 3 {
//
//		}
//	}
//	return min_int_number(dp[l-1][1], dp[l-1][2], dp[l-1][3])
//}

func dp_minimumOperations2826(nums []int, l int, idx int, last int, memo [][4]int) int {
	if idx == l {
		return 0
	}
	if memo[idx][last] != -1 {
		return memo[idx][last]
	}
	var res int = 1<<31 - 1
	for i := 1; i <= 3; i++ {
		if i < last {
			continue
		}
		var cur int = 0
		if nums[idx] == i {
			cur = dp_minimumOperations2826(nums, l, idx+1, nums[idx], memo)
		} else {
			cur = 1 + dp_minimumOperations2826(nums, l, idx+1, i, memo)
		}
		if cur < res {
			res = cur
		}
	}
	memo[idx][last] = res
	return res
}

func MinimumOperations2826(nums []int) int {
	var l int = len(nums)
	var memo [][4]int = make([][4]int, l)
	for i := 0; i < l; i++ {
		for j := 1; j <= 3; j++ {
			memo[i][j] = -1
		}
	}
	return dp_minimumOperations2826(nums, l, 0, 1, memo)
}
