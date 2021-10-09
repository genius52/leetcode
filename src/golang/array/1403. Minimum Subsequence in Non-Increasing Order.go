package array

import "sort"

func minSubsequence(nums []int) []int {
	var sum int = 0
	var l int = len(nums)
	for i := 0;i < l;i++{
		sum += nums[i]
	}
	sort.Ints(nums)
	var cur int = 0
	var res []int
	for i := l - 1;i >= 0;i--{
		cur += nums[i]
		res = append(res,nums[i])
		if cur > (sum - cur){
			break
		}
	}
	return res
}