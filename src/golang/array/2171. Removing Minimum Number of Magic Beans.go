package array

import "sort"

func MinimumRemoval(beans []int) int64 {
	var l int = len(beans)
	sort.Ints(beans)
	var prefix []int64 = make([]int64,l + 1)
	//prefix[0] = beans[0]
	for i := 0;i < l;i++{
		prefix[i + 1] = prefix[i] + int64(beans[i])
	}
	var res int64 = 1 << 62
	for i := 0;i < l;i++{
		//把比自己小的清空
		less_than_cur := prefix[i]
		more_than_cur := prefix[l] - prefix[i] - int64(beans[i]) * int64(l - i)
		if less_than_cur + more_than_cur < res{
			res = less_than_cur + more_than_cur
		}
	}
	return res
}