package array

func maxAscendingSum(nums []int) int {
	var l int = len(nums)
	var res int = nums[0]
	var cur_sum int = nums[0]
	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			cur_sum += nums[i]
		} else {
			cur_sum = nums[i]
		}
		if cur_sum > res {
			res = cur_sum
		}
	}
	return res
}
