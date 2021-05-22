package list_queue

import "container/list"

func DailyTemperatures(T []int) []int{
	var l int = len(T)
	var res []int = make([]int,l)
	var q list.List
	for i := 0;i < l;i++{
		for q.Len() > 0 && T[q.Back().Value.(int)] < T[i]{
			var index int = q.Back().Value.(int)
			res[index] = i - index
			q.Remove(q.Back())
		}
		q.PushBack(i)
	}
	return res
}

//func dailyTemperatures(T []int) []int {
//	l := len(T)
//	var s []int
//	var res []int = make([]int,l)
//	for i := 0;i < l;i++{
//		if len(s) == 0{
//			s = append(s,i)
//			continue
//		}
//		for len(s) > 0 && T[i] > T[s[len(s) - 1]]{
//			res[s[len(s) - 1]] = i - s[len(s) - 1]
//			s = s[0:len(s) - 1]
//		}
//		s = append(s,i)
//	}
//	return res
//}