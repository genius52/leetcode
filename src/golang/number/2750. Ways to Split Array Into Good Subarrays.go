package number

func numberOfGoodSubarraySplits(nums []int) int {
	var l int = len(nums)
	var idx []int
	for i := 0; i < l; i++ {
		if nums[i] == 1 {
			idx = append(idx, i)
		}
	}
	if len(idx) <= 1 {
		return len(idx)
	}
	var res int = 1
	var MOD int = 1e9 + 7
	for i := 1; i < len(idx); i++ {
		res *= idx[i] - idx[i-1]
		res %= MOD
	}
	return res
}
