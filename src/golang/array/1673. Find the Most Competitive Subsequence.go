package array

import "container/list"

func MostCompetitive(nums []int, k int) []int {
	var l int = len(nums)
	var q list.List
	for i := 0;i < l;i++{
		//rest_len := l - i //剩下的长度 + q的长度必须大于k,等于K也不行
		for q.Len() > 0 && q.Back().Value.(int) > nums[i] && (l - i + q.Len()) > k {
			q.Remove(q.Back())
		}
		if q.Len() < k{
			q.PushBack(nums[i])
		}
	}
	var res []int
	for q.Len() > 0{
		res = append(res,q.Front().Value.(int))
		q.Remove(q.Front())
	}
	return res
}