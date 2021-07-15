package array

import (
	"container/list"
	"sort"
)

//Input: [17,13,11,2,3,5,7]
//Output: [2,13,3,11,5,17,7]


//Simulate reverse queue
func DeckRevealedIncreasing(deck []int) []int {
	var l int = len(deck)
	var res []int = make([]int,l)
	sort.Ints(deck)
	var q list.List
	for i := l - 1;i >= 0;i--{
		if q.Len() != 0{
			//把最后的放到最前面来
			pre := q.Back().Value.(int)
			q.Remove(q.Back())
			q.PushFront(pre)
		}
		q.PushFront(deck[i])
	}
	for i := 0;i < l;i++{
		res[i] = q.Front().Value.(int)
		q.Remove(q.Front())
	}
	return res
}