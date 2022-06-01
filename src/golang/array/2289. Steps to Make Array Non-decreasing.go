package array

import "container/list"

//每个数会将其右侧所有比自己小的数消除，直到碰到第一个不小于自己的数为止。找右侧首个不小于自己的数是单调栈的典型应用场景

func TotalSteps(nums []int) int {
	var l int = len(nums)
	var q list.List //单调递减的栈
	var dp []int = make([]int,l)//dp[i]: 在第几轮被删除
	var res int = 0
	for i := 0;i < l;i++{
		if q.Len() == 0{
			q.PushBack(i)
		}else{
			var max_round int = 0
			for q.Len() > 0 && nums[i] >= nums[q.Back().Value.(int)]{
				max_round = max_int(max_round,dp[q.Back().Value.(int)])
				q.Remove(q.Back())
			}
			if q.Len() != 0{
				res = max_int(res,max_round + 1)
				dp[i] = max_round + 1
			}
			q.PushBack(i)
		}
	}
	return res
}