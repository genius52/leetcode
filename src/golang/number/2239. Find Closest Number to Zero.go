package number

import "math"

func findClosestNumber(nums []int) int {
	var l int = len(nums)
	var min_diff int = 2147483647
	var min_diff_num int = 0
	for i := 0;i < l;i++{
		if nums[i] == 0{
			return 0
		}
		diff := int(math.Abs(float64(nums[i])))
		if diff < min_diff{
			min_diff = diff
			min_diff_num = nums[i]
		}else if diff == min_diff{
			if min_diff_num < 0{
				min_diff_num = nums[i]
			}
		}
	}
	return min_diff_num
}