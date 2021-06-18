package array

func PeakIndexInMountainArray(arr []int) int {
	var l int = len(arr)
	var low int = 0
	var high int = l - 1
	for low < high{
		mid := (low + high)/2
		if arr[mid - 1] < arr[mid] && arr[mid] > arr[mid + 1]{
			return mid
		}else if arr[mid] > arr[mid - 1]{
			low = mid + 1
		}else if arr[mid] > arr[mid + 1]{
			high = mid
		}
	}
	return low
}