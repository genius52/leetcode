package array

import "sort"

//Input: intervals = [[1,4],[3,6],[2,8]]
//Output: 2
//Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.


type Interval [][]int

func (s Interval) Less(i, j int) bool {
	if s[i][0] < s[j][0]{
		return true
	}else if s[i][0] == s[j][0]{
		return s[i][1] > s[j][1]
	}
	return false
}

func (s Interval) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Interval) Len() int {
	return len(s)
}

func is_contain(start1 int,end1 int,start2 int,end2 int)bool{
	if start1 >= start2 && end1 <= end2{
		return true
	}
	return false
}

func RemoveCoveredIntervals(intervals [][]int) int {
	sort.Sort(Interval(intervals))
	l := len(intervals)
	var cnt int = l
	var cur_start int = intervals[0][0]
	var cur_end int = intervals[0][1]
	for i := 1;i < l;i++{
		if is_contain(intervals[i][0],intervals[i][1],cur_start,cur_end){
			cnt--
		}else{
			if intervals[i][0] > cur_end{
				cur_start = intervals[i][0]
				cur_end = intervals[i][1]
			}else{
				cur_end = intervals[i][1]
			}
		}
	}
	return cnt
}