package array

func Search33(nums []int, target int) int {
	var l int = len(nums)
	var left int = 0
	var right int = l - 1
	for left < right{
		mid := (left + right)/2
		if nums[mid] == target{
			return mid
		}else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}else{
		//}else if nums[mid] < nums[right]{
			if nums[mid] < target && target <= nums[right]{
				left = mid + 1
			}else{
				right = mid - 1
			}
		}
	}
	if nums[left] == target{
		return left
	}
	return -1
}