package array

import "sort"

//Input: nums = [1,2,2]
//Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
func dfs_subsetsWithDup(nums []int,l int,target_len int,pos int,data *[]int,res *[][]int){
	if target_len == 0{
		var l2 int = len(*data)
		if l2 == 0{
			*res = append(*res,[]int{})
			return
		}
		var d []int = make([]int,l2)
		copy(d,*data)
		*res = append(*res,d)
		return
	}
	if pos >= l{
		return
	}
	//skip current number
	//dfs_subsetsWithDup(nums,l,target_len,pos + 1,data,res)
	//add current number
	for i := pos;i < l;i++{
		if i > pos{
			if nums[i] == nums[i - 1]{
				continue
			}
		}
		*data = append(*data,nums[i])
		dfs_subsetsWithDup(nums,l,target_len - 1,i + 1,data,res)
		var cur_len int = len(*data)
		*data = (*data)[:cur_len - 1]
	}
}

func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var l int = len(nums)
	for i := 0;i <= l;i++{
		var data []int
		dfs_subsetsWithDup(nums,l,i,0,&data,&res)
	}
	return res
}