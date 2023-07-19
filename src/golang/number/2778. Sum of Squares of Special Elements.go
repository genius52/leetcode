package number

func sumOfSquares(nums []int) int {
	var l int = len(nums)
	var res int = 0
	for i := 0; i < l; i++ {
		if (l % (i + 1)) == 0 {
			res += nums[i] * nums[i]
		}
	}
	return res
}
