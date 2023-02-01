package array

import "sort"

func putMarbles(weights []int, k int) int64 {
	var l int = len(weights)
	var diff []int
	for i := 1;i < l;i++{
		diff = append(diff,weights[i] + weights[i - 1])
	}
	sort.Ints(diff)
	var res int64 = 0
	for i := 0;i < k - 1;i++{
		res += int64(diff[l - 2 - i])
		res -= int64(diff[i])
	}
	return res
}