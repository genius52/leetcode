package array

func findSpecialInteger(arr []int) int {
	var l int = len(arr)
	var step int = l/4
	for i := 0;i + step < l;i++{
		if arr[i] == arr[i + step]{
			return arr[i]
		}
	}
	return -1
}