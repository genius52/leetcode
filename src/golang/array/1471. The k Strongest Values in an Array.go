package array

import (
	"sort"
)

func getStrongest(arr []int, k int) []int {
	var l int = len(arr)
	if k >= l{
		return arr
	}
	sort.Ints(arr)
	median := arr[(l - 1)/2]
	var res []int
	var left int = 0
	var right int = l - 1
	for left < right && k > 0{
		l := abs_int(arr[left] - median)
		r := abs_int(arr[right] - median)
		if l > r{
			res = append(res,arr[left])
			left++
		}else {
			res = append(res,arr[right])
			right--
		}
		k--
	}
	return res
}