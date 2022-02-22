package number

func dfs_countMaxOrSubsets(nums []int,l int,idx int,cur_val int,max_val int)int{
	if cur_val == max_val{
		return 1 << (l - idx)
	}
	if idx >= l{
		if cur_val == max_val{
			return 1
		}
		return 0
	}
	return dfs_countMaxOrSubsets(nums,l,idx + 1,cur_val,max_val) +
		dfs_countMaxOrSubsets(nums,l,idx + 1,cur_val | nums[idx],max_val)
}

func countMaxOrSubsets(nums []int) int {
	var l int = len(nums)
	var max_val int = 0
	for i := 0;i < l;i++{
		max_val |= nums[i]
	}
	return dfs_countMaxOrSubsets(nums,l,0,0,max_val)
}