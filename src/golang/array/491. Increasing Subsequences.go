package array

import (
	"strconv"
	"strings"
)

func dfs_findSubsequences(nums []int,cur_pos int,cur_nums *[]int,record map[string]bool){
	l := len(*cur_nums)
	if l >= 2{
		var data string
		for _,n := range *cur_nums{
			if len(data) != 0{
				data += ","
			}
			data += strconv.Itoa(n)
		}
		record[data] = true
	}

	if cur_pos >= len(nums){
		return
	}
	if l > 0 {
		if nums[cur_pos] >= (*cur_nums)[l - 1]{
			var add_cur_nums []int
			add_cur_nums  = make([]int,l)
			copy(add_cur_nums,*cur_nums)
			add_cur_nums = append(add_cur_nums,nums[cur_pos])
			dfs_findSubsequences(nums,cur_pos+1,&add_cur_nums,record)
		}
		dfs_findSubsequences(nums,cur_pos+1,cur_nums,record)
	}else{
		var start_cur_nums []int
		start_cur_nums = append(start_cur_nums,nums[cur_pos])
		dfs_findSubsequences(nums,cur_pos+1,&start_cur_nums,record)
		var ignore_cur_nums []int
		dfs_findSubsequences(nums,cur_pos+1,&ignore_cur_nums,record)
	}
}
func findSubsequences(nums []int) [][]int {
	var res [][]int
	if len(nums) <= 0{
		return res
	}
	var cur_nums []int
	var record map[string]bool = make(map[string]bool)
	dfs_findSubsequences(nums,0,&cur_nums,record)
	for  r,_ := range record{
		var data_list []string = strings.Split(r,",")
		var tmp []int
		for i := 0;i < len(data_list);i++{
			d , err := strconv.ParseInt(data_list[i],10,64)
			if err == nil{
				tmp = append(tmp, int(d))
			}
		}
		res = append(res, tmp)
	}
	return res
}