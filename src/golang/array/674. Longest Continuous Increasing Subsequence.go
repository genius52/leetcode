package array

//Input: nums = [1,3,5,4,7]
//Output: 3
//Explanation: The longest continuous increasing subsequence is [1,3,5] with length 3.
//Even though [1,3,5,7] is an increasing subsequence,
//it is not continuous as elements 5 and 7 are separated by element
//4.
func findLengthOfLCIS(nums []int) int {
	var l int = len(nums)
	var max_len int = 1
	var left int = 0
	for left < l{
		var right int = left + 1
		for right < l && nums[right] > nums[right - 1]{
			right++
		}
		cur_len := right - left
		if cur_len > max_len{
			max_len = cur_len
		}
		left = right
	}
	return max_len
}