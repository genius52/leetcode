package array

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0]{
			return intervals[i][1] < intervals[j][1]
		}else{
			return intervals[i][0] < intervals[j][0]
		}
	})
	var l int = len(intervals)
	var dup_cnt int = 0
	var pre_end int = -2147483648
	for i := 0;i < l;i++{
		if intervals[i][0] < pre_end{
			dup_cnt++
			pre_end = min_int(pre_end,intervals[i][1])
		}else{
			pre_end = intervals[i][1]
		}
	}
	return dup_cnt
}