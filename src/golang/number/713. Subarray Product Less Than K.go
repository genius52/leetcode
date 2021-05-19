package number

//Input: nums = [10, 5, 2, 6], k = 100
//Output: 8
func NumSubarrayProductLessThanK(nums []int, k int) int {
	var l int = len(nums)
	var left int = 0
	var right int = 0
	var sum int = 1
	var res int = 0
	for left < l && right < l{
		sum *= nums[right]
		for left < l &&  sum >= k{
			sum /= nums[left]
			left++
		}
		res += right - left + 1
		right++
	}
	return res
}