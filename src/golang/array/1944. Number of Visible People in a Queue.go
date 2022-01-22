package array

import "container/list"

func CanSeePersonsCount(heights []int) []int {
	var l int = len(heights)
	var q list.List
	var res []int = make([]int,l)
	res[l - 1] = 0
	var greater []int = make([]int,l)
	greater[l - 1] = 0
	for i := l - 1;i >= 0;i--{
		if q.Len() == 0{
			q.PushBack(i)
		}else{
			var cnt int = 0
			for q.Len() > 0 && heights[i] > heights[q.Back().Value.(int)]{
				cnt++
				q.Remove(q.Back())
			}
			res[i] = cnt
			if q.Len() > 0{
				res[i]++
			}
			q.PushBack(i)
		}
	}
	return res
}