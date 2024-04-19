package array

func maximumPrimeDifference(nums []int) int {
	var l int = len(nums)
	var left int = 0
	for left < l {
		if is_prime(nums[left]) {
			break
		}
		left++
	}
	var right int = l - 1
	for right > left {
		if is_prime(nums[right]) {
			break
		}
		right--
	}
	return right - left
}
