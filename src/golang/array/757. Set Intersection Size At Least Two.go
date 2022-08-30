package array

import "sort"

//Example 1:
//
//Input: intervals = [[1,3],[1,4],[2,5],[3,5]]
//Output: 3
//Explanation: Consider the set S = {2, 3, 4}.  For each interval, there are at least 2 elements from S in the interval.
//Also, there isn't a smaller size set that fulfills the above condition.
//Thus, we output the size of this set, which is 3.

//Input: intervals = [[1,2],[2,3],[2,4],[4,5]]
//Output: 5
//Explanation: An example of a minimum sized set is {1, 2, 3, 4, 5}.
func IntersectionSizeTwo(intervals [][]int) int{
	var l int = len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1]{
			return intervals[i][0] > intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	var right int = -2147483648
	var left int = -2147483648
	var res int = 2
	for i := 0;i < l;i++{
		cur_left := intervals[i][0]
		cur_right := intervals[i][1]
		if cur_left > right{
			res += 2
			left = cur_right - 1
			right = cur_right
		}else if cur_left > left{
			res += 1
			left = right
			right = cur_right
		}
	}
	return res
}



//type End_begin [][]int
//
//func (s End_begin) Less(i, j int) bool {
//	if s[i][1] < s[j][1]{
//		return true
//	}else if s[i][1] == s[j][1]{
//		return s[i][0] < s[j][0]
//	}
//	return false
//}
//
//func (s End_begin) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//
//func (s End_begin) Len() int {
//	return len(s)
//}
//
//func IntersectionSizeTwo(intervals [][]int) int {
//	var l int = len(intervals)
//	sort.Sort(End_begin(intervals))
//	var pre_right int = intervals[0][len(intervals[0]) - 1]
//	var pre_left int = pre_right - 1
//	for i := 1;i < l;i++{
//		cur_left := intervals[i][0]
//		cur_right := intervals[i][1]
//		if cur_right == pre_right{
//			continue
//		}
//		if cur_left <= (pre_right - 1){
//			continue
//		}
//		pre_right = cur_left + 1
//	}
//	return pre_right - pre_left + 1
//}