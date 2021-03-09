package array

//Input: nums = [3,0,1]
//Output: 2
//0,1,2,3 ... n
func missingNumber(nums []int) int {
	var l int = len(nums)
	var cur_sum int = 0
	for i := 0;i < l;i++{
		cur_sum += nums[i]
	}
	var expect_sum int = ((l + 1) * l)/2
	return expect_sum - cur_sum
}