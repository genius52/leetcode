package array

//Input: nums = [2,4,1,1,6,5]
//Output: 3
func CountHillValley(nums []int) int {
	var l int = len(nums)
	var res int = 0
	var idx int = 1
	for idx < l - 1{
		if nums[idx] == nums[idx - 1]{
			idx++
			continue
		}
		//var left int = nums[idx - 1]
		var right int = idx + 1
		for right < l && nums[right] == nums[idx]{
			right++
		}
		if right >= l{
			break
		}
		if (nums[idx] > nums[idx - 1] && nums[idx] > nums[right]) || (nums[idx] < nums[idx - 1] && nums[idx] < nums[right]){
			res++
		}
		idx = right
	}
	return res
}