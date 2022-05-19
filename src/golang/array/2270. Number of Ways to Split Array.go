package array

func waysToSplitArray(nums []int) int {
	var sum int = 0
	var l int = len(nums)
	for i := 0;i < l;i++{
		sum += nums[i]
	}
	var res int = 0
	var left_sum int = 0
	for i := 0;i < l - 1;i++{
		left_sum += nums[i]
		right_sum := sum - left_sum
		if left_sum >= right_sum{
			res++
		}
	}
	return res
}