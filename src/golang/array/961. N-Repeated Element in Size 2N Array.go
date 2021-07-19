package array

//Input: nums[2,1,2,5,3,2]
//Output: 2
func repeatedNTimes(nums []int) int {
	var l int = len (nums)
	for i := 2;i < l;i++{
		if nums[i] == nums[i - 1] || nums[i] == nums[i - 2]{
			return nums[i]
		}
	}
	return nums[0]
}