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
	memo[pos] = res
	return res
}

func MinSwap(A []int, B []int) int {
	var l int = len(A)
	if l <= 1{
		return 0
	}
	var memo []int = make([]int,l)
	res := dp_minSwap(A,B,-1,-1,0,memo)
	return res
}