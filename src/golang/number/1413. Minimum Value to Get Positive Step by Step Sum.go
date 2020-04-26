package number

//Input: nums = [-3,2,-3,4,2]
//Output: 5
func minStartValue(nums []int) int {
	l := len(nums)
	if l == 0{
		return 0
	}
	var sum int = 0
	var min int = 1
	for i := 0;i < l;i++{
		sum += nums[i]
		if sum < 0{
			if (1 - sum) > min{
				min = 1 - sum
			}
		}
	}
	return min
}