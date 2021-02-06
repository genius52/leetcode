package array

import (
	"sort"
	"strconv"
)

//47
//[1,1,2,3,3,4,5,6]
func dfs_permuteUnique(cur_pos int,nums []int,res *[][]int,record map[string]bool){
	if cur_pos == len(nums){
		var s string
		for _,n := range nums{
			s += strconv.Itoa(n)
		}
		if _,ok := record[s];ok{
			return
		}else{
			record[s] = true
		}
		var c []int = make([]int,len(nums))
		copy(c,nums)
		*res = append(*res,c)
		return
	}
	for i := cur_pos;i < len(nums);i++{
		if i > cur_pos{
			if nums[i] == nums[i-1] || nums[i] == nums[cur_pos]{
				continue
			}
		}
		nums[cur_pos],nums[i] = nums[i],nums[cur_pos]
		dfs_permuteUnique(cur_pos + 1,nums,res,record)
		nums[cur_pos],nums[i] = nums[i],nums[cur_pos]
	}
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	var record map[string]bool = make(map[string]bool)
	dfs_permuteUnique(0,nums,&res,record)
	return res
}
