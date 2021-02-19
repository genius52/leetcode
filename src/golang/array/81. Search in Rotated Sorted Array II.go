package array

//Input: nums = [2,5,6,0,0,1,2], target = 0
//Output: true
func search(nums []int, target int) bool {
	var l int = len(nums)
	var low int = 0
	var high int = l - 1
	for low <= high{
		mid := low + (high - low)/2
		if nums[mid] == target{
			return true
		}
		if nums[low] == nums[mid] && nums[high] == nums[mid]{
			for low <= high{
				if nums[low] == target{
					return true
				}
				low += 1
			}
			return false
		}

		if nums[low] <= nums[mid] {
			if nums[low] <= target && target <= nums[mid]{
				high = mid - 1
			}else{
				low = mid + 1
			}
		}else{
			if nums[mid] <= target && target <= nums[high]{
				low = mid + 1
			}else{
				high = mid - 1
			}
		}
	}
	return false
}