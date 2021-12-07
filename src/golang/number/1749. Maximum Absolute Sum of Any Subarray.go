package number

//Input: nums = [2,-5,1,-4,3,-2]
//Output: 8
func MaxAbsoluteSum(nums []int) int {
	var pre_max int = nums[0] //最大
	var pre_min int = nums[0] //最小
	var res int = abs_int(nums[0])
	var l int = len(nums)
	for i := 1; i < l; i++ {
		cur_max := max_int(nums[i], nums[i]+pre_max)
		cur_min := min_int(nums[i], nums[i]+pre_min)
		pre_max = cur_max
		pre_min = cur_min
		res = max_int(res, max_int(abs_int(cur_max), abs_int(cur_min)))
	}
	return res
}

//func MaxAbsoluteSum(nums []int) int {
//	var l int = len(nums)
//	var dp_positive []int = make([]int, l)
//	var dp_negtive []int = make([]int, l)
//	dp_positive[0] = nums[0]
//	dp_negtive[0] = nums[0]
//	var res int = int(math.Abs(float64(nums[0])))
//	for i := 1; i < l; i++ {
//		dp_positive[i] = max_int(dp_positive[i-1]+nums[i], nums[i])
//		res = max_int(res, dp_positive[i])
//		dp_negtive[i] = min_int(dp_negtive[i-1]+nums[i], nums[i])
//		res = max_int(res, int(math.Abs(float64(dp_negtive[i]))))
//	}
//	return res
//}
