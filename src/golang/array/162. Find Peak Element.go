package array


//162
func findPeakElement(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	if nums[0] >= nums[1] {
		return 0
	}
	if nums[l - 1] >= nums[l - 2]{
		return l - 1
	}
	low := 0
	high := l - 1
	for low < high{
		mid := (low + high)/2
		if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1]{
			return mid
		}
		if nums[mid-1] > nums[mid]{
			high = mid
		}else{
			low = mid + 1
		}
	}
	return 0
}

func findPeakElement2(nums []int) int {
	var l int = len(nums)
	if l == 1{
		return 0
	}
	if nums[0] > nums[1]{
		return 0
	}
	if nums[l - 1] > nums[l - 2]{
		return l - 1
	}
	for i := 1;i < l - 1;i++{
		if nums[i] > nums[i - 1] && nums[i] > nums[i + 1]{
			return i
		}
	}
	return 0
}