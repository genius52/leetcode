package number

func FindMin(nums []int) int {
	var l int = len(nums)
	if l == 1{
		return nums[0]
	}
	var low int = 0
	var high int = l - 1
	var min_num int = nums[0]
	for low < high{
		mid := low + (high - low)/2
		if nums[mid] > nums[high]{
			low = mid + 1
			min_num = nums[low]
		}else if nums[mid] < nums[high]{
			high = mid
			min_num = nums[high]
		}else{
			high--;
			//min_num = nums[high]
		}
	}
	return min_num
}