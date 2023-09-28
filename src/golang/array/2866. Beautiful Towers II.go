package array

import "container/list"

func MaximumSumOfHeights2866(maxHeights []int) int64 {
	var l int = len(maxHeights)
	var left []int64 = make([]int64, l) //left[i]: maxHeights[i]为最大值时，sum(0,i)
	var right []int64 = make([]int64, l)
	var q1 list.List
	q1.PushBack(-1)
	var left_sum int64 = 0
	for i := 0; i < l; i++ {
		for q1.Len() > 1 && maxHeights[q1.Back().Value.(int)] > maxHeights[i] {
			pre_idx1 := q1.Back().Value.(int)
			q1.Remove(q1.Back())
			pre_idx2 := q1.Back().Value.(int)
			left_sum -= int64(maxHeights[pre_idx1] * (pre_idx1 - pre_idx2))
		}
		left_sum += int64(maxHeights[i] * (i - q1.Back().Value.(int)))
		left[i] = left_sum
		q1.PushBack(i)
	}
	var q2 list.List
	q2.PushBack(l)
	var right_sum int64 = 0
	for i := l - 1; i >= 0; i-- {
		for q2.Len() > 1 && maxHeights[q2.Back().Value.(int)] > maxHeights[i] {
			pre_idx1 := q2.Back().Value.(int)
			q2.Remove(q2.Back())
			pre_idx2 := q2.Back().Value.(int)
			right_sum -= int64(maxHeights[pre_idx1] * (pre_idx2 - pre_idx1))
		}
		right_sum += int64(maxHeights[i] * (q2.Back().Value.(int) - i))
		right[i] = right_sum
		q2.PushBack(i)
	}
	var res int64 = 0
	for i := 0; i < l; i++ {
		if i == l-1 {
			if left[i] > res {
				res = left[i]
			}
		} else {
			cur := left[i] + right[i+1]
			if cur > res {
				res = cur
			}
		}
	}
	return res
}
