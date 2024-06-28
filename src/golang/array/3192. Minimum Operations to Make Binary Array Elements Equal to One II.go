package array

func minOperations3192(nums []int) int {
	var l int = len(nums)
	var res int = 0
	var flip bool = false
	for i := 0; i < l; i++ {
		if (nums[i] == 0 && !flip) || (nums[i] == 1 && flip) {
			res++
			flip = !flip
		}
	}
	return res
}
