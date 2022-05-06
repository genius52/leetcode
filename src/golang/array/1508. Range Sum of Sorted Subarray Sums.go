package array

import "sort"

func rangeSum(nums []int, n int, left int, right int) int {
	var sum []int = make([]int,n * (n + 1)/2)
	var prefix []int = make([]int,n + 1)
	for i := 0;i < n;i++{
		prefix[i + 1] = prefix[i] + nums[i]
	}
	var idx int = 0
	for i := 1;i <= n;i++{
		for j := i - 1;j >= 0;j--{
			sum[idx] = prefix[i] - prefix[j]
			idx++
		}
	}
	var res int = 0
	var MOD int = 1e9 + 7
	sort.Ints(sum)
	for i := left - 1;i <= right - 1;i++{
		res += sum[i]
		res %= MOD
	}
	return res
}