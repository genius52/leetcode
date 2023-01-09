package number

//((nums[i] | nums[j]) & nums[k]) ^
func xorBeauty(nums []int) int {
	var l int = len(nums)
	var res int = 0
	for i := 0;i < l;i++{
		res ^= nums[i]
	}
	return res
}