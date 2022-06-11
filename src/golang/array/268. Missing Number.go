package array

//Input: nums = [3,0,1]
//Output: 2
//0,1,2,3 ... n

//Sum
func missingNumber(nums []int) int {
	var l int = len(nums)
	var cur_sum int = 0
	for i := 0;i < l;i++{
		cur_sum += nums[i]
	}
	var expect_sum int = ((l + 1) * l)/2
	return expect_sum - cur_sum
}

//XOR  i * nums[i]
func missingNumber2(nums []int) int {
	var res int = 0
	var l int = len(nums)
	for i := 0;i < l;i++{
		res = res ^ i ^ nums[i]
	}
	return res ^ l
}