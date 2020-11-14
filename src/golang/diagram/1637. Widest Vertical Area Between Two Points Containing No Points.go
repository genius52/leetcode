package diagram

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	var record map[int]bool = make(map[int]bool)
	for _,p := range points{
		record[p[0]] = true
	}
	var nums []int
	for pos,_ := range record{
		nums = append(nums,pos)
	}
	sort.Ints(nums)
	var res int = 0
	var l int = len(nums)
	for i := 1;i < l;i++{
		interval := nums[i] - nums[i - 1]
		if interval > res{
			res = interval
		}
	}
	return res
}