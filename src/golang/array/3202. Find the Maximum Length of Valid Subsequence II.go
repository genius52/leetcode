package array

func MaximumLength3202(nums []int, k int) int {
	var l int = len(nums)
	var pre_num_mod_cnt [][]int = make([][]int, k) //pre_mod_cnt[i][j],数字i(num % k)结尾，余数为j的最大长度
	for i := 0; i < k; i++ {
		pre_num_mod_cnt[i] = make([]int, k)
	}
	var res int = 0
	for i := 0; i < l; i++ {
		cur_mod := nums[i] % k
		for j := 0; j < k; j++ {
			find_pre_mod := (j - cur_mod + k) % k
			pre_num_mod_cnt[find_pre_mod][cur_mod] = pre_num_mod_cnt[cur_mod][find_pre_mod] + 1
			res = max(res, pre_num_mod_cnt[find_pre_mod][cur_mod])
		}
	}
	return res
}
