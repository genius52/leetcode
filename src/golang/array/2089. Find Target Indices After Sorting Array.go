package array

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	var l int = len(nums)
	var res []int
	for i := 0;i < l;i++{
		if nums[i] == target{
			res = append(res,i)
		}
	}
	return res
}

func targetIndices2(nums []int, target int) []int {
	var less_cnt int = 0
	var equal_cnt int = 0
	var l int = len(nums)
	for i := 0;i < l;i++{
		if nums[i] < target{
			less_cnt++
		}else if nums[i] == target{
			equal_cnt++
		}
	}
	var res []int = make([]int,equal_cnt)
	var idx int = 0
	for i := 0;i < equal_cnt;i++{
		res[i] = less_cnt + idx
		idx++
	}
	return res
}