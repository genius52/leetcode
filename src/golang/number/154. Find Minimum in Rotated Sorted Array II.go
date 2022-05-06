package number

func FindMin(nums []int) int {
	var l int = len(nums)
	if l == 1{
		return nums[0]
	}
	var left int = 0
	var right int = l - 1
	var min_num int = nums[0]
	for left < right{
		mid := left + (right - left)/2
		if nums[mid] > nums[right]{
			left = mid + 1
			min_num = nums[left]
		}else if nums[mid] < nums[right]{
			right = mid
			min_num = nums[right]
		}else{
			right--
		}
	}
	return min_num
}