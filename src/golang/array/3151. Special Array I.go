package array

func isArraySpecial(nums []int) bool {
	var l int = len(nums)
	for i := 1; i < l; i++ {
		if (nums[i-1]^nums[i])&1 != 1 {
			return false
		}
	}
	return true
}
