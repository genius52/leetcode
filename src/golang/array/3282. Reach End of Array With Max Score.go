package array

func dfs_findMaximumScore(nums []int, l int, idx int) int64 {
	if idx >= l {
		return 0
	}
	var res int64 = int64(l-1-idx) * int64(nums[idx])
	var next int = idx + 1
	for next < l {
		if nums[next] > nums[idx] {
			score := int64(next-idx)*int64(nums[idx]) + dfs_findMaximumScore(nums, l, next)
			res = max(res, score)
			break
		}
		next++
	}
	return res
}

func findMaximumScore(nums []int) int64 {
	var l int = len(nums)
	return dfs_findMaximumScore(nums, l, 0)
}
