package array

func SumCounts(nums []int) int {
	var MOD int = 1e9 + 7
	var l int = len(nums)
	var res int = l
	for sub_len := 2; sub_len <= l; sub_len++ {
		for left := 0; left+sub_len <= l; left++ {
			var num_cnt map[int]bool = make(map[int]bool)
			for right := left; right < left+sub_len; right++ {
				num_cnt[nums[right]] = true
			}
			cur := len(num_cnt)
			res += cur * cur
			res %= MOD
		}
	}
	return res
}
