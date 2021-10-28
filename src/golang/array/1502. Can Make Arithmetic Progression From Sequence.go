package array

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	var l int = len(arr)
	if l <= 2{
		return true
	}
	sort.Ints(arr)
	for i := 1;i < l - 1;i++{
		if arr[i] - arr[i - 1] != arr[i + 1] - arr[i]{
			return false
		}
	}
	return true
}