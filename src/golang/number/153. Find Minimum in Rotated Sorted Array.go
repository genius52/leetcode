package number

//Input: nums = [4,5,6,7,0,1,2]
//Output: 0
func findMin(nums []int) int {
	var l int = len(nums)
	if l == 1{
		return nums[0]
	}
	var low int = 0
	var high int = l - 1
	var min_num int = 1<<31
	for low < high{
		mid := low + (high - low)/2
		if nums[mid] > nums[high]{
			low = mid + 1
			min_num = nums[low]
		}else{
			high = mid
			min_num = nums[high]
		}
	}
	return min_num
}