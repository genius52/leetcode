package array


//		var A []int = []int{0,7,8,10,10,11,12,13,19,18}
//		var B []int = []int{4,4,5,7,11,14,15,16,17,20}

//swapRecord means for the ith element in A and B, the minimum swaps if we swap A[i] and B[i]
//fixRecord means for the ith element in A and B, the minimum swaps if we DONOT swap A[i] and B[i]
func MinSwap(A []int, B []int) int {
	var l int = len(A)
	if l <= 1{
		return 0
	}
	var swap []int = make([]int,l)
	swap[0] = 1
	var not_swap []int = make([]int,l)
	not_swap[0] = 0
	for i := 1;i < l;i++{
		//We define areBothSelfIncreasing: A[i - 1] < A[i] && B[i - 1] < B[i], areInterchangeIncreasing: A[i - 1] < B[i] && B[i - 1] < A[i].
		//Since 'the given input always makes it possible', at least one of the two conditions above should be satisfied.
		//无论是否交换都能保持递增
		if (A[i] > A[i - 1] && B[i] > B[i - 1]) && (A[i - 1] < B[i] && B[i - 1] < A[i]){
			not_swap[i] = min_int(not_swap[i - 1],swap[i - 1])
			swap[i] = min_int(not_swap[i - 1],swap[i - 1]) + 1
		} else if A[i] > A[i - 1] && B[i] > B[i - 1]{
			not_swap[i] = not_swap[i - 1]
			swap[i] = swap[i- 1] + 1
		}else if A[i - 1] < B[i] && B[i - 1] < A[i]{
			not_swap[i] = swap[i - 1]
			swap[i] = not_swap[i - 1] + 1
		}
	}
	return min_int(swap[l - 1],not_swap[l - 1])
}