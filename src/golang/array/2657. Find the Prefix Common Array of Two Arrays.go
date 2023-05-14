package array

func FindThePrefixCommonArray(A []int, B []int) []int {
	var l int = len(A)
	var res []int = make([]int, l)
	var visited_A []bool = make([]bool, l+1)
	var visited_B []bool = make([]bool, l+1)
	var same_cnt int = 0
	for i := 0; i < l; i++ {
		visited_A[A[i]] = true
		visited_B[B[i]] = true
		if A[i] == B[i] {
			same_cnt++
		} else {
			if visited_A[B[i]] {
				same_cnt++
			}
			if visited_B[A[i]] {
				same_cnt++
			}
		}
		res[i] = same_cnt
	}
	return res
}
