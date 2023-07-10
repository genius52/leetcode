package array

func alternatingSubarray(nums []int) int {
	var l int = len(nums)
	var res int = -1
	var pre_small_start_len int = 1
	var pre_big_start_len int = 1
	for i := l - 2; i >= 0; i-- {
		if (nums[i] + 1) == nums[i+1] {
			res = max_int(res, 1+pre_big_start_len)
			pre_small_start_len = 1 + pre_big_start_len
			pre_big_start_len = 1
		} else if (nums[i] - 1) == nums[i+1] {
			pre_big_start_len = 1 + pre_small_start_len
			pre_small_start_len = 1
		} else {
			pre_small_start_len = 1
			pre_big_start_len = 1
		}
	}
	return res
}
