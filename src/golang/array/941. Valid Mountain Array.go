package array

func ValidMountainArray(arr []int) bool {
	var l int = len(arr)
	if l <= 2{
		return false
	}
	var left int = 0
	for left < (l - 1) && arr[left] < arr[left + 1]{
		left++
	}
	if left == 0 || left == (l - 1){
		return false
	}
	var right int = l - 1
	for right >= 0 && arr[right] < arr[right - 1]{
		right--
	}
	return left == right
}