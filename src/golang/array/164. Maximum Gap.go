package array

func MaximumGap(nums []int) int {
	var l int = len(nums)
	if l < 2 {
		return 0
	}
	var offset int = 1
	var max_val int = 0
	for i := 0; i < l; i++ {
		if nums[i] > max_val {
			max_val = nums[i]
		}
	}
	for offset <= max_val {
		var buf []int = make([]int, l)
		var prefix_cnt [10]int
		for i := 0; i < l; i++ {
			m := (nums[i] / offset) % 10
			prefix_cnt[m]++
		}
		for i := 1; i <= 9; i++ {
			prefix_cnt[i] += prefix_cnt[i-1]
		}
		//sort by mod
		for i := l - 1; i >= 0; i-- {
			m := (nums[i] / offset) % 10
			idx := prefix_cnt[m] - 1
			buf[idx] = nums[i]
			prefix_cnt[m]--
		}
		copy(nums, buf)
		offset *= 10
	}
	var res int = 0
	for i := 1; i < l; i++ {
		if (nums[i] - nums[i-1]) > res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}
