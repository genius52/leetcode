package array

func SearchInsert(nums []int, target int) int{
	var l int = len(nums)
	if nums[l - 1] < target{
		return l
	}
	var left int = 0
	var right int = l - 1
	for left < right{
		var mid int = left + (right - left)/2
		if nums[mid] == target{
			return mid
		}else if nums[mid] < target{
			left = mid + 1
		}else if nums[mid] > target{
			right = mid
		}
	}
	return left
}

//func SearchInsert(nums []int, target int) int {
//	var l int = len(nums)
//	if nums[0] >= target{
//		return 0
//	}
//	if nums[l - 1] < target{
//		return l
//	}
//	var low int = 0
//	var high int = l - 1
//	for low < high{
//		var mid int = low + (high - low)/2
//		if nums[mid] == target{
//			return mid
//		}else if nums[mid] < target{
//			if nums[mid + 1] >= target{
//				return mid + 1
//			}
//			low = mid + 1
//		}else if nums[mid] > target{
//			if nums[mid - 1] < target{
//				return mid
//			}
//			high = mid - 1
//		}
//	}
//	return -1
//}