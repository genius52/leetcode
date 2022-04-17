package array

import "sort"

//Input: nums = [1,2,2]
//Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
//func dfs_subsetsWithDup(nums []int,l int,target_len int,pos int,data *[]int,res *[][]int){
//	if target_len == 0{
//		var l2 int = len(*data)
//		if l2 == 0{
//			*res = append(*res,[]int{})
//			return
//		}
//		var d []int = make([]int,l2)
//		copy(d,*data)
//		*res = append(*res,d)
//		return
//	}
//	if pos >= l{
//		return
//	}
//	//add current number
//	for i := pos;i < l;i++{
//		if i > pos{
//			if nums[i] == nums[i - 1]{
//				continue
//			}
//		}
//		*data = append(*data,nums[i])
//		dfs_subsetsWithDup(nums,l,target_len - 1,i + 1,data,res)
//		var cur_len int = len(*data)
//		*data = (*data)[:cur_len - 1]
//	}
//}
//
//func SubsetsWithDup(nums []int) [][]int {
//	sort.Ints(nums)
//	var res [][]int
//	var l int = len(nums)
//	for i := 0;i <= l;i++{
//		var data []int
//		dfs_subsetsWithDup(nums,l,i,0,&data,&res)
//	}
//	return res
//}

func dfs_subsetsWithDup(nums []int,l int,idx int,pre *[]int,res *[][]int){
	var pre_len int = len(*pre)
	var cur []int = make([]int,pre_len)
	copy(cur,*pre)
	*res = append(*res,cur)
	for i := idx;i < l;i++{
		if i != idx && nums[i] == nums[i - 1]{
			continue
		}
		//if i == idx || nums[i] != nums[i - 1]{
			*pre = append(*pre,nums[i])
			dfs_subsetsWithDup(nums,l,i + 1,pre,res)
			*pre = (*pre)[:len(*pre) - 1]
		//}
	}
}

func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var l int = len(nums)
	var pre []int
	dfs_subsetsWithDup(nums,l,0,&pre,&res)
	return res
}