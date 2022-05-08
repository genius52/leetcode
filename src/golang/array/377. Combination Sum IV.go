package array

func dfs_combinationSum4(nums []int,l int,target int,memo []int)int{
	if target == 0{
		return 1
	}
	if target < 0{
		return 0
	}
	if memo[target] != -1{
		return memo[target]
	}
	var res int = 0
	for i := 0;i < l;i++{
		res += dfs_combinationSum4(nums,l,target - nums[i],memo)
	}
	memo[target] = res
	return res
}

func combinationSum4(nums []int, target int) int {
	var l int = len(nums)
	var memo []int = make([]int,target + 1)
	for i := 1;i <= target;i++{
		memo[i] = -1
	}
	return dfs_combinationSum4(nums,l,target,memo)
}