package array

func leftRigthDifference(nums []int) []int {
	var l int = len(nums)
	var res []int = make([]int, l)
	res[0] = nums[0]
	for i := 1; i < l; i++ {
		res[i] = res[i-1] + nums[i]
	}
	res[l-1] -= nums[l-1]
	var right_sum int = nums[l-1]
	for i := l - 2; i >= 0; i-- {
		right_sum += nums[i]
		res[i] -= right_sum
		if res[i] < 0 {
			res[i] = -res[i]
		}
	}
	return res
}
