package array

//Input: nums = [2,3,5]
//Output: [4,3,5]
//Explanation: Assuming the arrays are 0-indexed, then
//result[0] = |2-2| + |2-3| + |2-5| = 0 + 1 + 3 = 4,
//result[1] = |3-2| + |3-3| + |3-5| = 1 + 0 + 2 = 3,
//result[2] = |5-2| + |5-3| + |5-5| = 3 + 2 + 0 = 5.
func GetSumAbsoluteDifferences(nums []int) []int {
	var l int = len(nums)
	var res []int = make([]int,l)
	var left_sum []int = make([]int,l)
	left_sum[0] = nums[0]
	var right_sum []int = make([]int,l)
	right_sum[l - 1] = nums[l - 1]
	for i := 1;i < l;i++{
		left_sum[i] = left_sum[i - 1] + nums[i]
		right_sum[l - 1 - i] = right_sum[l - i] + nums[l - 1 - i]
	}
	for i := 0;i < l;i++{
		var left_val int = 0
		var right_val int = 0
		if i > 0{
			left_val = nums[i] * i - left_sum[i - 1]
		}
		if i < l - 1{
			right_val = right_sum[i + 1] - nums[i] * (l - 1 - i)
		}
		res[i] = left_val + right_val
	}
	return res
}