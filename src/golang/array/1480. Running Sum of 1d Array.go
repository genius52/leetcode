package array

func runningSum(nums []int) []int {
	var l int = len(nums)
	for i := 1;i < l;i++{
		nums[i] = nums[i - 1] + nums[i]
	}
	return nums
}