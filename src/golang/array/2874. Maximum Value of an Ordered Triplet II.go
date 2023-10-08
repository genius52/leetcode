package array

func maximumTripletValue2(nums []int) int64 {
	var l int = len(nums)
	var res int64 = 0
	var left_max []int = make([]int, l)
	var right_max []int = make([]int, l)
	left_max[0] = nums[0]
	for i := 1; i < l; i++ {
		left_max[i] = max_int(left_max[i-1], nums[i])
	}
	right_max[l-1] = nums[l-1]
	for i := l - 2; i >= 0; i-- {
		right_max[i] = max_int(right_max[i+1], nums[i])
	}
	for i := 1; i < l-1; i++ {
		val := int64(left_max[i-1]-nums[i]) * int64(right_max[i+1])
		if val > res {
			res = val
		}
	}
	return res
}
