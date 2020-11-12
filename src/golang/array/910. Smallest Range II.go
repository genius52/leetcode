package array

import (
	"sort"
)

//func min_int(a,b int)int{
//	if a < b {
//		return a
//	}else{
//		return b
//	}
//}

//func abs_int(n int)int{
//	if n < 0{
//		return -n
//	}
//	return n
//}
//Input: A = [1,3,6], K = 3
//Output: 3
//Explanation: B = [4,6,3]
func SmallestRangeII(A []int, K int) int {
	sort.Ints(A)
	var l int = len(A)
	max_num := A[l - 1]
	min_num := A[0]
	res := A[l - 1] - A[0]
	for i := 0;i < l - 1;i++{
		if A[i] + K * 2 > max_num {
			max_num = A[i] + K*2
		}
		min_num = min_int(A[0] + K * 2,A[i + 1])
		res = min_int(res,max_num - min_num)
	}
	return res
}