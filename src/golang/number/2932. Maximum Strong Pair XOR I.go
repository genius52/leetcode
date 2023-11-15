package number

func maximumStrongPairXor(nums []int) int {
	var l int = len(nums)
	var max_val int = 0
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if abs_int(nums[i]-nums[j]) <= min_int(nums[i], nums[j]) {
				cur := nums[i] ^ nums[j]
				if cur > max_val {
					max_val = cur
				}
			}
		}
	}
	return max_val
}
