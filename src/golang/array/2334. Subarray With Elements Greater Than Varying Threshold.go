package array

import "container/list"

//1,3,4,3,1

func ValidSubarraySize(nums []int, threshold int) int {
	var l int = len(nums)
	var dp_left []int = make([]int,l)//dp_left[i]: 左边第一个小于当前数字的长度
	var q1 list.List //单调递减
	for i := 0;i < l;i++{
		var find_idx int = -1
		for q1.Len() > 0 && nums[i] <= nums[q1.Back().Value.(int)]{
			q1.Remove(q1.Back())
		}
		if q1.Len() > 0{
			find_idx = q1.Back().Value.(int)
		}
		if find_idx == -1{
			dp_left[i] = i + 1
		}else{
			dp_left[i] = i - find_idx
		}
		q1.PushBack(i)
	}
	var dp_right []int = make([]int,l)//dp_right[i]: 左边第一个小于当前数字的长度
	var q2 list.List
	for i := l - 1;i >= 0;i--{
		var find_idx int = l
		for q2.Len() > 0 && nums[i] <= nums[q2.Back().Value.(int)]{
			q2.Remove(q2.Back())
		}
		if q2.Len() > 0{
			find_idx = q2.Back().Value.(int)
		}
		if find_idx == l{
			dp_right[i] = l - i
		}else{
			dp_right[i] = find_idx - i
		}
		q2.PushBack(i)
	}
	var res int = -1
	for i := 0;i < l;i++{
		cur_len := dp_left[i] + dp_right[i] - 1
		if nums[i] > threshold / cur_len{
			//if cur_len > res{
			//	res = cur_len
			//}
			return cur_len //返回任意子数组的长度
		}
	}
	return res
}