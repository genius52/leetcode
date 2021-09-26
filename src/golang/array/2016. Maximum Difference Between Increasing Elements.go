package array

func maximumDifference(nums []int) int {
	var l int = len(nums)
	var max_diff int = -1
	var min_num int = 2147483647
	for i := 0;i < l;i++{
		if nums[i] > min_num{
			max_diff = max_int(max_diff,nums[i] - min_num)
		}
		min_num = min_int(min_num,nums[i])
	}
	return max_diff
}