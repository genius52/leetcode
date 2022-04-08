package array

func PeakIndexInMountainArray(arr []int) int {
	var l int = len(arr)
	var left int = 0
	var right int = l - 1
	for left < right{
		mid := (left + right)/2
		if arr[mid - 1] < arr[mid] && arr[mid] > arr[mid + 1]{
			return mid
		}else if arr[mid] > arr[mid - 1]{
			left = mid + 1
		}else if arr[mid] > arr[mid + 1]{
			right = mid - 1
		}
	}
	return left
}