package array

import "sort"

//Input: candidates = [2,3,6,7], target = 7,
//A solution set is:
//[
//  [7],
//  [2,2,3]
//]
func dfs_combinationSum(candidates []int,l int,idx int,pre_sum int,target int,cur_nums []int,res *[][]int){
	if pre_sum == target{
		var b = make([]int, len(cur_nums))
		copy(b, cur_nums)
		*res = append(*res,b)
		return
	}
	if pre_sum > target || idx >= l || pre_sum + candidates[idx] > target{
		return
	}
	for i := 0;pre_sum <= target;i++{
		dfs_combinationSum(candidates,l,idx + 1,pre_sum,target,cur_nums,res)
		pre_sum += candidates[idx]
		cur_nums = append(cur_nums,candidates[idx])
	}
}

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var cur_nums []int
	dfs_combinationSum(candidates,len(candidates),0,0,target,cur_nums,&res)
	return res
}