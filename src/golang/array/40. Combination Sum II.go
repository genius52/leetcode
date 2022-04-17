package array

import "sort"

func dfs_combinationSum2(candidates []int,l int,idx int,pre_sum int,target int,cur_nums []int,res *[][]int){
	if pre_sum == target{
		var b = make([]int, len(cur_nums))
		copy(b, cur_nums)
		*res = append(*res,b)
		return
	}
	if pre_sum > target || idx >= l || pre_sum + candidates[idx] > target{
		return
	}
	for i := idx;i < l;i++{
		if i > idx && candidates[i] == candidates[i - 1]{
			continue
		}
		pre_sum += candidates[i]
		cur_nums = append(cur_nums,candidates[i])
		dfs_combinationSum2(candidates,l,i + 1,pre_sum,target,cur_nums,res)
		pre_sum -= candidates[i]
		cur_nums = cur_nums[:len(cur_nums) - 1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var cur_nums []int
	dfs_combinationSum2(candidates,len(candidates),0,0,target,cur_nums,&res)
	return res
}