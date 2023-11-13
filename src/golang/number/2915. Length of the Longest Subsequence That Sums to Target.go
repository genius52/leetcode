package number

import "sort"

func dp_lengthOfLongestSubsequence(nums []int, idx int, l int, target int, memo [][]int) int {
	if target == 0 {
		return 0
	}
	if idx == l {
		return -1
	}
	if memo[idx][target] != -2 {
		return memo[idx][target]
	}
	//skip current
	ret1 := dp_lengthOfLongestSubsequence(nums, idx+1, l, target, memo)
	//choose current
	var ret2 int = -1
	if nums[idx] <= target {
		ret2 = dp_lengthOfLongestSubsequence(nums, idx+1, l, target-nums[idx], memo)
		if ret2 != -1 {
			ret2++
		}
	}
	if ret1 == -1 {
		memo[idx][target] = ret2
	} else if ret2 == -1 {
		memo[idx][target] = ret1
	} else {
		memo[idx][target] = max_int(ret1, ret2)
	}
	return memo[idx][target]
}

func lengthOfLongestSubsequence(nums []int, target int) int {
	sort.Ints(nums)
	var l int = len(nums)
	var memo [][]int = make([][]int, l)
	for i := 0; i < l; i++ {
		memo[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			memo[i][j] = -2
		}
	}
	return dp_lengthOfLongestSubsequence(nums, 0, l, target, memo)
}
