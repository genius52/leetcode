package array

func search704(nums []int, target int) int {
	var l int = len(nums)
	var left int = 0
	var right int = l - 1
	for left < right{
		mid := left + (right - left)/2
		if nums[mid] == target{
			return mid
		}else if nums[mid] < target{
			left = mid + 1
		}else if nums[mid] > target{
			right = mid - 1
		}
	}
	return -1
}