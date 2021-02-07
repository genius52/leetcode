package number

import "math"

//Input: nums = [2,-5,1,-4,3,-2]
//Output: 8
func MaxAbsoluteSum(nums []int) int {
	var l int = len(nums)
	var dp_positive []int = make([]int,l)
	var dp_negtive []int = make([]int,l)
	dp_positive[0] = nums[0]
	dp_negtive[0] = nums[0]
	var res int = int(math.Abs(float64(nums[0])))
	for i := 1;i < l;i++{
		dp_positive[i] = max_int(dp_positive[i - 1] + nums[i],nums[i])
		res = max_int(res,dp_positive[i])
		dp_negtive[i] = min_int(dp_negtive[i - 1] + nums[i],nums[i])
		res = max_int(res,int(math.Abs(float64(dp_negtive[i]))))
	}
	return res
}