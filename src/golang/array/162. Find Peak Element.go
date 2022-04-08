package array

func findPeakElement(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	left := 0
	right := l - 1
	for left < right{
		mid := (left + right)/2
		if nums[mid] < nums[mid + 1]{
			left = mid + 1
		}else{
			right = mid
		}
	}
	return left
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