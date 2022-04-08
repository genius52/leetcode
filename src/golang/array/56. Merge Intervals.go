package array

import "sort"

//56
//Input: [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]
//Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
func merge(intervals [][]int) [][]int {
	var l int = len(intervals)
	if l <= 1{
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0]{
			return intervals[i][1] < intervals[j][1]
		}else{
			return intervals[i][0] < intervals[j][0]
		}
	})
	var res [][]int
	res = append(res,intervals[0])
	for i := 1;i < len(intervals);i++{
		if res[len(res)-1][1] < intervals[i][0]{
			res = append(res,intervals[i])
		} else{
			if res[len(res)-1][1] <= intervals[i][1]{
				res[len(res)-1][1] = intervals[i][1]
			}
		}
	}
	return res
}