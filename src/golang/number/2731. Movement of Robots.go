package number

import "sort"

func SumDistance(nums []int, s string, d int) int {
	var l int = len(nums)
	var pos []int = make([]int, l)
	for i := 0; i < l; i++ {
		if s[i] == 'L' {
			pos[i] = nums[i] - d
		} else if s[i] == 'R' {
			pos[i] = nums[i] + d
		}
	}
	sort.Ints(pos)
	var MOD int = 1e9 + 7
	var res int = 0
	var prefix_sum int = 0
	for i := 1; i < l; i++ {
		res = (res + prefix_sum + (pos[i]-pos[i-1])*i) % MOD
		prefix_sum += (pos[i] - pos[i-1]) * i
		prefix_sum %= MOD
	}
	return res
}
