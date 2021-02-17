package array

import "sort"

type Position [][]int

func (pos Position) Len()int{
	return len(pos)
}

func (pos Position) Less(i,j int)bool{
	if pos[i][0] == pos[j][0]{
		return pos[i][1] >= pos[j][1]
	}
	return pos[i][0] <= pos[j][0]
}

func (pos Position) Swap(i,j int){
	pos[i],pos[j] = pos[j],pos[i]
}

//56
//Input: [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]
//Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
func merge(intervals [][]int) [][]int {
	var l int = len(intervals)
	if l <= 1{
		return intervals
	}
	sort.Sort(Position(intervals))
	var res [][]int
	res = append(res,intervals[0])
	for i := 1;i < len(intervals);i++{
		if res[len(res)-1][1] >= intervals[i][0]{
			if res[len(res)-1][1] <= intervals[i][1]{
				res[len(res)-1][1] = intervals[i][1]
			}
		} else{
			res = append(res,intervals[i])
		}
	}
	return res
}