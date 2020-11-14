package array

import "math"

func dp_minSwap(A []int, B []int,prev_a int,prev_b int,pos int,memo []int)int{
	if pos >= len(A){
		return 0
	}
	if memo[pos] != 0{
		return memo[pos]
	}
	var res int = math.MaxInt32
	if A[pos] > prev_b && B[pos] > prev_a{
		change := 1 + dp_minSwap(A,B,B[pos],A[pos],pos + 1,memo)
		res = int(math.Min(float64(res),float64(change)))
	}
	if A[pos] > prev_a && B[pos] > prev_b{
		no_change := dp_minSwap(A,B,A[pos],B[pos],pos + 1,memo)
		res = int(math.Min(float64(res),float64(no_change)))
	}
	memo[pos] = res
	return res
}

func MinSwap(A []int, B []int) int {
	var l int = len(A)
	if l <= 1{
		return 0
	}
	var memo []int = make([]int,l)
	return dp_minSwap(A,B,-1,-1,0,memo)
}