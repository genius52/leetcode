package array

func subArrayRanges(nums []int) int64 {
	var l int = len(nums)
	var res int64 = 0
	for i := 0; i < l; i++ {
		min := nums[i]
		max := nums[i]
		for j := i + 1; j < l; j++ {
			if nums[j] < min {
				min = nums[j]
			}
			if nums[j] > max {
				max = nums[j]
			}
			res += int64(max) - int64(min)
		}
	}
	return res
}
