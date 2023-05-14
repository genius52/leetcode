package array

func distinctDifferenceArray(nums []int) []int {
	var l int = len(nums)
	var prefix_cnt []int = make([]int, l)
	var prefix_num map[int]bool = make(map[int]bool)
	var cur_cnt int = 0
	for i := 0; i < l; i++ {
		if _, ok := prefix_num[nums[i]]; !ok {
			cur_cnt++
		}
		prefix_num[nums[i]] = true
		prefix_cnt[i] = cur_cnt
	}
	var suffix_num map[int]bool = make(map[int]bool)
	var suffix_cnt []int = make([]int, l)
	cur_cnt = 0
	for i := l - 1; i >= 0; i-- {
		if _, ok := suffix_num[nums[i]]; !ok {
			cur_cnt++
		}
		suffix_num[nums[i]] = true
		suffix_cnt[i] = cur_cnt
	}
	var res []int = make([]int, l)
	res[l-1] = prefix_cnt[l-1]
	for i := 0; i < l-1; i++ {
		res[i] = prefix_cnt[i] - suffix_cnt[i+1]
	}
	return res
}
