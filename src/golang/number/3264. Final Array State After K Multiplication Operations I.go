package number

func getFinalState(nums []int, k int, multiplier int) []int {
	var l int = len(nums)
	for k > 0 {
		var min_idx int = -1
		var min_num int = 1<<31 - 1
		for i := 0; i < l; i++ {
			if nums[i] < min_num {
				min_idx = i
				min_num = nums[i]
			}
		}
		nums[min_idx] = nums[min_idx] * multiplier
		k--
	}
	return nums
}
