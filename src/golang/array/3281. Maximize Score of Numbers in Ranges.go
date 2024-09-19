package array

import "sort"

func check_maxPossibleScore(start []int, d int, target int) bool {
	var pre int = -1 << 31
	for i := 0; i < len(start); i++ {
		cur := max(pre+target, start[i])
		if cur > start[i]+d {
			return false
		}
		pre = cur
	}
	return true
}

func MaxPossibleScore(start []int, d int) int {
	var l int = len(start)
	sort.Ints(start)
	//var res int = 1<<31 - 1
	//for i := 1; i < l; i++ {
	//	cur := start[i]
	//}
	var low int = 0
	var high int = start[l-1] - start[0] + d
	var res int = low
	for low <= high {
		mid := (low + high) / 2
		ret := check_maxPossibleScore(start, d, mid)
		if ret {
			res = mid
			low = mid + 1
		} else {
			//res = mid - 1
			high = mid - 1
		}
	}
	return res
}
