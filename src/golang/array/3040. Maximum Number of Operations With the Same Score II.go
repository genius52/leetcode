package array

func dp_maxOperations3040(nums []int, left int, right int, score int, memo [][]int) int {
	if left >= right {
		return 0
	}
	if memo[left][right] != -1 {
		return memo[left][right]
	}
	var res = 0
	if nums[left]+nums[left+1] == score {
		res = max_int(res, 1+dp_maxOperations3040(nums, left+2, right, score, memo))
	}
	if nums[right]+nums[right-1] == score {
		res = max_int(res, 1+dp_maxOperations3040(nums, left, right-2, score, memo))
	}
	if nums[left]+nums[right] == score {
		res = max_int(res, 1+dp_maxOperations3040(nums, left+1, right-1, score, memo))
	}
	memo[left][right] = res
	return res
}

func MaxOperations3040(nums []int) int {
	var l int = len(nums)
	var memo [][]int = make([][]int, l)
	for i := 0; i < l; i++ {
		memo[i] = make([]int, l)
		for j := 0; j < l; j++ {
			memo[i][j] = -1
		}
	}
	return 1 + max_int_number(dp_maxOperations3040(nums, 2, l-1, nums[0]+nums[1], memo),
		dp_maxOperations3040(nums, 0, l-3, nums[l-1]+nums[l-2], memo),
		dp_maxOperations3040(nums, 1, l-2, nums[0]+nums[l-1], memo))
}
