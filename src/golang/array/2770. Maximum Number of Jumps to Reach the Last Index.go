package array

func dfs_maximumJumps(nums []int, l int, idx int, target int, memo []int) int {
	if idx == l {
		return 0
	}
	if idx == l-1 {
		return 0
	}
	if memo[idx] != -2 {
		return memo[idx]
	}
	var res int = -1
	for i := idx + 1; i < l; i++ {
		if abs_int(nums[idx]-nums[i]) <= target {
			cur := dfs_maximumJumps(nums, l, i, target, memo)
			if cur != -1 {
				res = max_int(res, 1+cur)
			}
		}
	}
	memo[idx] = res
	return res
}

func MaximumJumps(nums []int, target int) int {
	var l int = len(nums)
	var memo []int = make([]int, l)
	for i := 0; i < l; i++ {
		memo[i] = -2
	}
	return dfs_maximumJumps(nums, l, 0, target, memo)
}
