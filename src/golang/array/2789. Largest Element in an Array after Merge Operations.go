package array

func maxArrayValue(nums []int) int64 {
	var l int = len(nums)
	var res int64 = int64(nums[l-1])
	var pre int64 = int64(nums[l-1])
	for i := l - 2; i >= 0; i-- {
		if int64(nums[i]) <= pre {
			pre += int64(nums[i])
		} else {
			pre = int64(nums[i])
		}
		if pre > res {
			res = pre
		}
	}
	return res
}
