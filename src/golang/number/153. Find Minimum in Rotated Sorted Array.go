package number

//Input: nums = [4,5,6,7,0,1,2]
//Output: 0
func findMin(nums []int) int {
	var l int = len(nums)
	if l == 1{
		return nums[0]
	}
	var left int = 0
	var right int = l - 1
	var min_num int = 1<<31
	for left < right{
		mid := left + (right - left)/2
		if nums[mid] > nums[right]{
			left = mid + 1
			min_num = nums[left]
		}else{
			right = mid
			min_num = nums[right]
		}
	}
	return min_num
}