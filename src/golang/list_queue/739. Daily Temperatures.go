package list_queue

import "container/list"

func DailyTemperatures(temperatures []int) []int {
	var l int = len(temperatures)
	var res []int = make([]int,l)
	var q list.List
	for i := 0;i < l;i++{
		for q.Len() > 0 && temperatures[q.Back().Value.(int)] < temperatures[i]{
			var index int = q.Back().Value.(int)
			res[index] = i - index
			q.Remove(q.Back())
		}
		q.PushBack(i)
	}
	return res
}