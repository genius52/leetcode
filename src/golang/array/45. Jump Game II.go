package array


//45
//Input: [2,3,1,1,4]
//Output: 2
//Explanation: The minimum number of jumps to reach the last index is 2.
//    Jump 1 step from index 0 to 1, then 3 steps to the last index.
func jump(nums []int) int {
	var dp []int = make([]int,len(nums))
	dp[len(nums) - 1] = 0
	for i := len(nums) - 2;i >=0;i--{
		min_step := len(nums)
		for l := 1;l <= nums[i] && (i + l) <= len(nums) - 1;l++{
			min_step = min_int_number(1 + dp[i + l],min_step)
			//min_step = min_int_number(1 + dp[len(nums) - 1 - i - j],min_step)
		}
		dp[i] = min_step
	}
	return dp[0]
}

func jump3(nums []int) int {
	var l int = len(nums)
	var steps int = 0
	var start int = 0
	var end int = 0
	var max_pos int = 0
	for max_pos < l - 1{
		for i := start;i <= end;i++{
			max_pos = max_int(max_pos,i + nums[i])
		}
		start = end + 1
		end = max_pos
		steps++
	}
	return steps
}