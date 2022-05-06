package array

import "sort"

func NumSubseq(nums []int, target int) int {
	var l int = len(nums)
	var MOD int = 1e9 + 7
	var pre []int = make([]int,l)
	pre[0] = 1;
	for i := 1;i < l;i++ {
		pre[i] = (pre[i - 1] << 1 ) % MOD
	}
	sort.Ints(nums)
	var res int = 0
	var left int = 0
	var right int = l - 1
	for left < l - 1{
		if nums[left] * 2 > target{
			break
		}
		for right > left && nums[left] + nums[right] > target{
			right--
		}
		res += pre[right - left]
		res %= MOD
		left++
	}
	return res
}