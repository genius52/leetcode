package array

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	var l int = len(arr)
	arr[0] = 1
	for i := 1;i < l;i++{
		if arr[i] - arr[i - 1] > 1{
			arr[i] = arr[i - 1] + 1
		}
	}
	return arr[l - 1]
}