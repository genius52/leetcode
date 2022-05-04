package array

import "sort"

func findRightInterval(intervals [][]int) []int {
	var l int = len(intervals)
	var start_idx [][2]int = make([][2]int,l)
	for i := 0;i < l;i++{
		start_idx[i] = [2]int{intervals[i][0],i}
	}
	sort.Slice(start_idx, func(i, j int) bool {
		return start_idx[i][0] < start_idx[j][0]
	})
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		find_idx := sort.Search(l, func(j int) bool {
			return start_idx[j][0] >= intervals[i][1]
		})
		if find_idx == l{
			res[i] = -1
		}else{
			res[i] = start_idx[find_idx][1]
		}
	}
	return res
}