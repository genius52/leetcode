package array

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	var res [][]int
	sort.Ints(arr)
	var min_diff int = 2147483647
	var l int = len(arr)
	for i := 1;i < l;i++{
		cur := arr[i] - arr[i - 1]
		if cur < min_diff{
			min_diff = cur
		}
	}
	for i := 1;i < l;i++ {
		if (arr[i] - arr[i - 1]) == min_diff{
			res = append(res,[]int{arr[i - 1],arr[i]})
		}
	}
	return res
}